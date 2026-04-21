package repositorydefinitions

import "auptex.com/botnova/internals/domain/models"

type ProjectRepository interface {
	Create(project models.Project) error
	GetById(id int) (*models.Project, error)
	Update(m models.Project) error
	Delete(id int) error
	GetByUserId(id int) ([]models.Project, error)
}
