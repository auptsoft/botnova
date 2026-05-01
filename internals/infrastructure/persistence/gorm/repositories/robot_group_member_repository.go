package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type RobotGroupMemberRepository struct {
	db *gorm.DB
}

func NewRobotGroupMemberRepository(db *gorm.DB) *RobotGroupMemberRepository {
	return &RobotGroupMemberRepository{db: db}
}

func (r *RobotGroupMemberRepository) Create(member *models.RobotGroupMember) error {
	member.Id = uuid.Must(uuid.NewV7()).String()
	entity := entity_mappers.ToRobotGroupMemberEntity(*member)
	return r.db.Create(&entity).Error
}

func (r *RobotGroupMemberRepository) GetByGroupId(groupId string) ([]models.RobotGroupMember, error) {
	var entities []entities.RobotGroupMember
	if err := r.db.Find(&entities, "group_id = ?", groupId).Error; err != nil {
		return nil, err
	}

	var result []models.RobotGroupMember

	for _, e := range entities {
		result = append(result, entity_mappers.ToRobotGroupMemberDomain(e))
	}
	return result, nil
}

func (r *RobotGroupMemberRepository) GetById(id string) (*models.RobotGroupMember, error) {
	var entity entities.RobotGroupMember
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	data := entity_mappers.ToRobotGroupMemberDomain(entity)
	return &data, nil
}

func (r *RobotGroupMemberRepository) GetByRobotId(robotId string) ([]models.RobotGroupMember, error) {
	var entities []entities.RobotGroupMember
	if err := r.db.Find(&entities, "robot_id = ?", robotId).Error; err != nil {
		return nil, err
	}

	var result []models.RobotGroupMember

	for _, e := range entities {
		result = append(result, entity_mappers.ToRobotGroupMemberDomain(e))
	}
	return result, nil
}

func (r *RobotGroupMemberRepository) GetByGroupIdAndRole(groupId string, role string) ([]models.RobotGroupMember, error) {
	var entities []entities.RobotGroupMember
	if err := r.db.Find(&entities, "group_id = ? AND role = ?", groupId, role).Error; err != nil {
		return nil, err
	}

	var result []models.RobotGroupMember

	for _, e := range entities {
		result = append(result, entity_mappers.ToRobotGroupMemberDomain(e))
	}
	return result, nil
}

func (r *RobotGroupMemberRepository) Update(member *models.RobotGroupMember) error {
	entity := entity_mappers.ToRobotGroupMemberEntity(*member)
	return r.db.Save(&entity).Error
}

func (r *RobotGroupMemberRepository) Delete(id string) error {
	return r.db.Delete(&entities.RobotGroupMember{}, "id = ?", id).Error
}
