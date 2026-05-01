package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type UserRepository interface {
	Create(user models.User, passwordHash string) (*models.User, error)
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAuthByEmail(email string) (*models.UserAuth, error)
	Update(user models.User) error
	UpdatePassword(userID string, passwordHash string) error
	Delete(id string) error
}
