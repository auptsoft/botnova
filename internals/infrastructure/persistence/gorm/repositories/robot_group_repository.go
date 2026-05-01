package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type RobotGroupRepository struct {
	db *gorm.DB
}

func NewRobotGroupRepository(db *gorm.DB) *RobotGroupRepository {
	return &RobotGroupRepository{db: db}
}

func (rgr *RobotGroupRepository) Create(group *models.RobotGroup) error {
	// Generate a new UUID for the group
	group.Id = uuid.Must(uuid.NewV7()).String()

	// Map domain model to entity
	entity := entity_mappers.ToRobotGroupEntity(*group)

	// Save to database
	return rgr.db.Create(&entity).Error
}

func (rgr *RobotGroupRepository) GetById(id string) (*models.RobotGroup, error) {
	var entity entities.RobotGroup
	if err := rgr.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Map entity back to domain model
	data := entity_mappers.ToRobotGroupDomain(entity)
	return &data, nil
}

func (rgr *RobotGroupRepository) Update(group *models.RobotGroup) error {
	entity := entity_mappers.ToRobotGroupEntity(*group)
	return rgr.db.Save(&entity).Error
}

func (rgr *RobotGroupRepository) Delete(id string) error {
	return rgr.db.Delete(&entities.RobotGroup{}, "id = ?", id).Error
}

func (rgr *RobotGroupRepository) GetByUserId(userId string) ([]models.RobotGroup, error) {
	var entities []entities.RobotGroup
	if err := rgr.db.Find(&entities, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	var result []models.RobotGroup

	for _, e := range entities {
		result = append(result, entity_mappers.ToRobotGroupDomain(e))
	}
	return result, nil
}

func (rgr *RobotGroupRepository) GetByRobotId(robotId string) (*models.RobotGroup, error) {
	var entity entities.RobotGroup
	if err := rgr.db.First(&entity, "id IN (SELECT group_id FROM robot_group_members WHERE robot_id = ?)", robotId).Error; err != nil {
		return nil, err
	}

	data := entity_mappers.ToRobotGroupDomain(entity)
	return &data, nil
}
