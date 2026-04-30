package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type ScriptRepository struct {
	db *gorm.DB
}

func NewScriptRepository(db *gorm.DB) *ScriptRepository {
	return &ScriptRepository{db: db}
}

// Implement CRUD methods for ScriptRepository as needed, e.g. Create, GetById, Update, Delete
func (r *ScriptRepository) Create(script *models.Script) error {
	// Generate a new UUID for the script
	script.Id = uuid.Must(uuid.NewV7()).String()

	// Map domain model to entity
	entity := entity_mappers.ToScriptEntity(*script)

	// Save to database
	return r.db.Create(&entity).Error
}

func (r *ScriptRepository) GetById(id string) (*models.Script, error) {
	var entity entities.Script
	if err := r.db.Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, err
	}
	script := entity_mappers.ToScriptDomain(entity)
	return &script, nil
}

func (r *ScriptRepository) Update(script *models.Script) error {
	entity := entity_mappers.ToScriptEntity(*script)
	return r.db.Save(&entity).Error
}

func (r *ScriptRepository) Delete(id string) error {
	return r.db.Delete(&entities.Script{}, "id = ?", id).Error
}

func (r *ScriptRepository) GetByProjectId(projectId string) ([]models.Script, error) {
	var entities []entities.Script
	if err := r.db.Where("project_id = ?", projectId).Find(&entities).Error; err != nil {
		return nil, err
	}

	var scripts []models.Script
	for _, entity := range entities {
		scripts = append(scripts, entity_mappers.ToScriptDomain(entity))
	}
	return scripts, nil
}

func (r *ScriptRepository) GetByUserId(userId string) ([]models.Script, error) {
	var entities []entities.Script
	if err := r.db.Joins("JOIN projects ON scripts.project_id = projects.id").
		Where("projects.user_id = ?", userId).
		Find(&entities).Error; err != nil {
		return nil, err
	}

	var scripts []models.Script
	for _, entity := range entities {
		scripts = append(scripts, entity_mappers.ToScriptDomain(entity))
	}
	return scripts, nil
}
