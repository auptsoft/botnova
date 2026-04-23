package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type ProjectRepository interface {
	Create(project models.Project) error
	GetById(id string) (*models.Project, error)
	Update(m models.Project) error
	Delete(id string) error
	GetByUserId(id string) ([]models.Project, error)
}
