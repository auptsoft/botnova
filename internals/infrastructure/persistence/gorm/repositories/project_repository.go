package repositories

import (
	"auptex.com/botnova/internals/domain/models"
	"auptex.com/botnova/internals/infrastructure/persistence/gorm/entities"
	entity_mappers "auptex.com/botnova/internals/infrastructure/persistence/gorm/mappers"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

// Create(project models.Project) error
func (r *ProjectRepository) Create(project models.Project) error {
	project.Id = uuid.Must(uuid.NewV7()).String()
	entity := entity_mappers.ToEntity(project)
	return r.db.Create(&entity).Error
}

// GetById(id int) (*models.Project, error)
func (r *ProjectRepository) GetById(id string) (*models.Project, error) {
	var entity entities.Project
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	data := entity_mappers.ToDomain(entity)
	return &data, nil
}

// Update(m models.Project) error
func (r *ProjectRepository) Update(project models.Project) error {
	entity := entity_mappers.ToEntity(project)
	return r.db.Save(&entity).Error
}

// Delete(id int) error
func (r *ProjectRepository) Delete(id string) error {
	return r.db.Delete(&entities.Project{}, "id = ?", id).Error
}

// GetByUserId(id int) ([]models.Project, error)
func (r *ProjectRepository) GetByUserId(userId string) ([]models.Project, error) {
	var entities []entities.Project
	if err := r.db.Find(&entities, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	var result []models.Project

	for _, e := range entities {
		result = append(result, entity_mappers.ToDomain(e))
	}

	return result, nil
}
