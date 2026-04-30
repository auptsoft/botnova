package services

import (
	"errors"
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

type AuthConfig struct {
	JwtSecret      []byte
	JwtTTL         time.Duration
}

type UserService struct {
	userRepository repositorydefinitions.UserRepository
	serviceLogger  ports.Logger
	authConfig AuthConfig
}

func NewUserService(userRepository repositorydefinitions.UserRepository, serviceLogger ports.Logger, authConfig AuthConfig) *UserService {
	return &UserService{
		userRepository: userRepository,
		serviceLogger:  serviceLogger,
		authConfig: authConfig,
	}
}

func (us *UserService) SignUp(user models.User, password string) (*AuthResult, error) {
	normalizedName := strings.TrimSpace(user.Name)
	normalizedEmail := normalizeEmail(user.Email)

	existingUser, err := us.userRepository.GetByEmail(normalizedEmail)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	createdUser, err := us.userRepository.Create(models.User{
		Name:  normalizedName,
		Email: normalizedEmail,
	}, passwordHash)
	if err != nil {
		return nil, err
	}

	token, err := us.generateToken(createdUser.Id)
	if err != nil {
		return nil, err
	}

	return &AuthResult{Token: token, User: *createdUser}, nil
}

func (us *UserService) Login(email string, password string) (*AuthResult, error) {
	normalizedEmail := normalizeEmail(email)
	authUser, err := us.userRepository.GetAuthByEmail(normalizedEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if compareErr := bcrypt.CompareHashAndPassword([]byte(authUser.PasswordHash), []byte(password)); compareErr != nil {
		return nil, ErrInvalidCredentials
	}

	user, err := us.userRepository.GetById(authUser.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	token, err := us.generateToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &AuthResult{Token: token, User: *user}, nil
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

	return user, nil
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

	if err := us.userRepository.Update(*existingUser); err != nil {
		return nil, err
	}

	if strings.TrimSpace(password) != "" {
		passwordHash, hashErr := hashPassword(password)
		if hashErr != nil {
			return nil, hashErr
		}
		if err := us.userRepository.UpdatePassword(userID, passwordHash); err != nil {
			return nil, err
		}
	}

	updatedUser, err := us.userRepository.GetById(userID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (us *UserService) generateToken(userID string) (string, error) {
	if len(us.authConfig.JwtSecret) == 0 {
		return "", ErrInvalidTokenConfig
	}

	claims := jwt.RegisteredClaims{
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(us.authConfig.JwtTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(us.authConfig.JwtSecret)
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
