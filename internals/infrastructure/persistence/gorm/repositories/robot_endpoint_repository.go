package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type RobotEndpointRepository struct {
	db *gorm.DB
}

func NewRobotEndpointRepository(db *gorm.DB) *RobotEndpointRepository {
	return &RobotEndpointRepository{db: db}
}

// Implement CRUD methods for RobotEndpoint as needed, e.g. Create, GetById, Update, Delete
func (r *RobotEndpointRepository) Create(endpoint *models.RobotEndpoint) error {
	endpoint.Id = uuid.Must(uuid.NewV7()).String()
	entity := entity_mappers.ToRobotEndpointEntity(*endpoint)
	return r.db.Create(entity).Error
}

func (r *RobotEndpointRepository) GetById(id string) (*models.RobotEndpoint, error) {
	var entity entities.RobotEndpoint
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	data := entity_mappers.ToRobotEndpointDomain(entity)
	return &data, nil
}

func (r *RobotEndpointRepository) GetByRobotId(robotId string) ([]models.RobotEndpoint, error) {
	var entities []entities.RobotEndpoint
	if err := r.db.Find(&entities, "robot_id = ?", robotId).Error; err != nil {
		return nil, err
	}

	var result []models.RobotEndpoint

	for _, e := range entities {
		result = append(result, entity_mappers.ToRobotEndpointDomain(e))
	}
	return result, nil
}

func (r *RobotEndpointRepository) Update(endpoint *models.RobotEndpoint) error {
	entity := entity_mappers.ToRobotEndpointEntity(*endpoint)
	return r.db.Save(entity).Error
}

func (r *RobotEndpointRepository) Delete(id string) error {
	return r.db.Delete(&entities.RobotEndpoint{}, "id = ?", id).Error
}
