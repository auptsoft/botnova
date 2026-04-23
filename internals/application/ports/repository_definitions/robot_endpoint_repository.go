package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type RobotEndpointRepository interface {
	Create(project models.RobotEndpoint) error
	GetById(id string) (*models.RobotEndpoint, error)
	Update(m models.RobotEndpoint) error
	Delete(id string) error
	GetByUserId(id string) ([]models.RobotEndpoint, error)
}
