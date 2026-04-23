package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type UserRepository interface {
	Create(user models.User) error
	GetById(id int) (*models.User, error)
	Update(user models.User) error
	Delete(id int) error
}
