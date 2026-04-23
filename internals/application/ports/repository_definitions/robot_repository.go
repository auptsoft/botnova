package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type RobotRepository interface {
	Create(project models.Robot) error
	GetById(id string) (*models.Robot, error)
	Update(m models.Robot) error
	Delete(id string) error
	GetByUserId(id string) ([]models.Robot, error)
}
