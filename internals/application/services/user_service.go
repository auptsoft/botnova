package services

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"auptex.com/botnova/internals/application/ports"
	repositorydefinitions "auptex.com/botnova/internals/application/ports/repository_definitions"
	"auptex.com/botnova/internals/domain/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidTokenConfig = errors.New("invalid token configuration")
)

type AuthResult struct {
	Token string
	User  models.User
}

type UserService struct {
	userRepository repositorydefinitions.UserRepository
	serviceLogger  ports.Logger
	jwtSecret      []byte
	jwtTTL         time.Duration
}

func NewUserService(userRepository repositorydefinitions.UserRepository, serviceLogger ports.Logger) *UserService {
	jwtTTLHours := 24 * 7
	if ttlFromEnv := os.Getenv("JWT_TTL_HOURS"); ttlFromEnv != "" {
		parsedTTL, err := strconv.Atoi(ttlFromEnv)
		if err == nil && parsedTTL > 0 {
			jwtTTLHours = parsedTTL
		}
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-me"
	}

	return &UserService{
		userRepository: userRepository,
		serviceLogger:  serviceLogger,
		jwtSecret:      []byte(jwtSecret),
		jwtTTL:         time.Duration(jwtTTLHours) * time.Hour,
	}
}

func (us *UserService) SignUp(user models.User) (*AuthResult, error) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = normalizeEmail(user.Email)

	existingUser, err := us.userRepository.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := us.userRepository.Create(user); err != nil {
		return nil, err
	}

	createdUser, err := us.userRepository.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	token, err := us.generateToken(createdUser.Id)
	if err != nil {
		return nil, err
	}

	createdUser.Password = ""
	return &AuthResult{Token: token, User: *createdUser}, nil
}

func (us *UserService) Login(email string, password string) (*AuthResult, error) {
	normalizedEmail := normalizeEmail(email)
	user, err := us.userRepository.GetByEmail(normalizedEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); compareErr != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := us.generateToken(user.Id)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return &AuthResult{Token: token, User: *user}, nil
}

func (us *UserService) CreateUser(user models.User) error {
	us.serviceLogger.Info("Creating user...")
	_, err := us.SignUp(user)
	return err
}

func (us *UserService) Delete(userId string) error {
	us.serviceLogger.Info("Deleting user")
	if _, err := us.userRepository.GetById(userId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	return us.userRepository.Delete(userId)
}

func (us *UserService) GetById(userId string) (*models.User, error) {
	us.serviceLogger.Info("Getting user by ID")
	user, err := us.userRepository.GetById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (us *UserService) Update(user models.User) error {
	us.serviceLogger.Info("Updating user")
	_, err := us.UpdateUser(user.Id, user.Name, user.Email, user.Password)
	return err
}

func (us *UserService) UpdateUser(userID string, name string, email string, password string) (*models.User, error) {
	existingUser, err := us.userRepository.GetById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if strings.TrimSpace(name) != "" {
		existingUser.Name = strings.TrimSpace(name)
	}

	if strings.TrimSpace(email) != "" {
		normalizedEmail := normalizeEmail(email)
		if normalizedEmail != existingUser.Email {
			userWithEmail, emailErr := us.userRepository.GetByEmail(normalizedEmail)
			if emailErr == nil && userWithEmail != nil && userWithEmail.Id != existingUser.Id {
				return nil, ErrEmailAlreadyExists
			}
			if emailErr != nil && !errors.Is(emailErr, gorm.ErrRecordNotFound) {
				return nil, emailErr
			}
		}
		existingUser.Email = normalizedEmail
	}

	if strings.TrimSpace(password) != "" {
		existingUser.Password = password
	}

	if err := us.userRepository.Update(*existingUser); err != nil {
		return nil, err
	}

	updatedUser, err := us.userRepository.GetById(userID)
	if err != nil {
		return nil, err
	}

	updatedUser.Password = ""
	return updatedUser, nil
}

func (us *UserService) generateToken(userID string) (string, error) {
	if len(us.jwtSecret) == 0 {
		return "", ErrInvalidTokenConfig
	}

	claims := jwt.RegisteredClaims{
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(us.jwtTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(us.jwtSecret)
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
