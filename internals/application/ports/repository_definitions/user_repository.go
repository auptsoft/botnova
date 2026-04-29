package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type UserRepository interface {
	Create(user models.User) error
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user models.User) error
	Delete(id string) error
}
