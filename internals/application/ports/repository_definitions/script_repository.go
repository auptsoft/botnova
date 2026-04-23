package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type ScriptRepository interface {
	Create(project models.Script) error
	GetById(id string) (*models.Script, error)
	Update(m models.Script) error
	Delete(id string) error
	GetByUserId(id string) ([]models.Script, error)
}
