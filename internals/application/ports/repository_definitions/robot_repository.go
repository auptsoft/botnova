package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type RobotRepository interface {
	Create(robot *models.Robot) error
	GetById(id string) (*models.Robot, error)
	Update(robot *models.Robot) error
	Delete(id string) error
	GetByUserId(userId string) ([]models.Robot, error)
}
