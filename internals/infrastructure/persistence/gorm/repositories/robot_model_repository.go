package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type RobotModelRepository struct {
	db *gorm.DB
}

func NewRobotModelRepository(db *gorm.DB) *RobotModelRepository {
	return &RobotModelRepository{db: db}
}

// Implement CRUD methods for RobotModelRepository as needed, e.g. Create, GetById, Update, Delete
func (r *RobotModelRepository) Create(model *models.RobotModel) error {
	// Generate a new UUID for the model
	model.Id = uuid.Must(uuid.NewV7()).String()

	// Map domain model to entity
	entity := entity_mappers.ToRobotModelEntity(*model)

	// Save to database
	return r.db.Create(&entity).Error
}

func (r *RobotModelRepository) GetById(id string) (*models.RobotModel, error) {
	var entity entities.RobotModel
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Map entity back to domain model
	data := entity_mappers.ToRobotModelDomain(entity)
	return &data, nil
}

func (r *RobotModelRepository) GetByUserId(userId string) ([]models.RobotModel, error) {
	var entities []entities.RobotModel
	if err := r.db.Find(&entities, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	var result []models.RobotModel

	for _, e := range entities {
		result = append(result, entity_mappers.ToRobotModelDomain(e))
	}
	return result, nil
}

func (r *RobotModelRepository) Update(model *models.RobotModel) error {
	entity := entity_mappers.ToRobotModelEntity(*model)
	return r.db.Save(&entity).Error
}

func (r *RobotModelRepository) Delete(id string) error {
	return r.db.Delete(&entities.RobotModel{}, "id = ?", id).Error
}
