package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type RobotModelRepository interface {
	Create(robotModel *models.RobotModel) error
	GetById(id string) (*models.RobotModel, error)
	Update(robotModel *models.RobotModel) error
	Delete(id string) error
	GetByUserId(userId string) ([]models.RobotModel, error)
}
