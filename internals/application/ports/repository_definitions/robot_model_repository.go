package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type RobotModelRepository interface {
	Create(project models.RobotModel) error
	GetById(id string) (*models.RobotModel, error)
	Update(m models.RobotModel) error
	Delete(id string) error
	GetByUserId(id string) ([]models.RobotModel, error)
}
