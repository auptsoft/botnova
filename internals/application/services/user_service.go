package services

import (
	"auptex.com/botnova/internals/application/ports"
	repositorydefinitions "auptex.com/botnova/internals/application/ports/repository_definitions"
	"auptex.com/botnova/internals/domain/models"
)

type UserService struct {
	userRepository repositorydefinitions.UserRepository
	serviceLogger  ports.Logger
}

func NewUserService(userRepository repositorydefinitions.UserRepository, serviceLogger ports.Logger) *UserService {
	return &UserService{
		userRepository: userRepository,
		serviceLogger:  serviceLogger,
	}
}

func (us *UserService) CreateUser(user models.User) error {
	us.serviceLogger.Info("Creating user...")
	return us.userRepository.Create(user)
}

func (us *UserService) Delete(userId int) error {
	us.serviceLogger.Info("Deleting user")
	return us.userRepository.Delete(userId)
}

func (us *UserService) GetById(userId int) (*models.User, error) {
	us.serviceLogger.Info("Getting user by ID")
	return us.userRepository.GetById(userId)
}

func (us *UserService) Update(user models.User) error {
	us.serviceLogger.Info("Updating user")
	return us.userRepository.Update(user)
}
