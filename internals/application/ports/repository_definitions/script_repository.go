package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type ScriptRepository interface {
	Create(script *models.Script) error
	GetById(id string) (*models.Script, error)
	Update(script *models.Script) error
	Delete(id string) error
	GetByUserId(userId string) ([]models.Script, error)
}
