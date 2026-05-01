package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RobotRepository struct {
	db *gorm.DB
}

func NewRobotRepository(db *gorm.DB) *RobotRepository {
	return &RobotRepository{db: db}
}

func (r *RobotRepository) Create(robot *models.Robot) error {
	robot.Id = uuid.Must(uuid.NewV7()).String()
	entity := entity_mappers.ToRobotEntity(*robot)
	return r.db.Create(&entity).Error
}

func (r *RobotRepository) GetById(id string) (*models.Robot, error) {
	var entity entities.Robot
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	data := entity_mappers.ToRobotDomain(entity)
	return &data, nil
}

func (r *RobotRepository) GetByUserId(userId string) ([]models.Robot, error) {
	var entities []entities.Robot
	if err := r.db.Find(&entities, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	var result []models.Robot

	for _, e := range entities {
		result = append(result, entity_mappers.ToRobotDomain(e))
	}
	return result, nil
}

func (r *RobotRepository) Update(robot *models.Robot) error {
	entity := entity_mappers.ToRobotEntity(*robot)
	return r.db.Save(&entity).Error
}

func (r *RobotRepository) Delete(id string) error {
	return r.db.Delete(&entities.Robot{}, "id = ?", id).Error
}
