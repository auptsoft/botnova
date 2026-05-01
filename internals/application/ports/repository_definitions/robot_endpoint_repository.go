package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type RobotEndpointRepository interface {
	Create(robotEndpoint *models.RobotEndpoint) error
	GetById(id string) (*models.RobotEndpoint, error)
	GetByRobotId(robotId string) ([]models.RobotEndpoint, error)
	Update(robotEndpoint *models.RobotEndpoint) error
	Delete(id string) error
}
