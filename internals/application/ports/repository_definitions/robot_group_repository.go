package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type RobotGroupRepository interface {
	Create(group *models.RobotGroup) error
	GetById(id string) (*models.RobotGroup, error)
	Update(m *models.RobotGroup) error
	Delete(id string) error
	GetByUserId(id string) ([]models.RobotGroup, error)
}
