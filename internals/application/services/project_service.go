package services

import (
	"auptex.com/botnova/internals/application/ports"
	repositorydefinitions "auptex.com/botnova/internals/application/ports/repository_definitions"
	"auptex.com/botnova/internals/domain/models"
)

type ProjectService struct {
	projectRespository repositorydefinitions.ProjectRepository
	serviceLoger       ports.Logger
}

func NewProjectService(projectRespository repositorydefinitions.ProjectRepository, serviceLogger ports.Logger) *ProjectService {
	return &ProjectService{
		projectRespository: projectRespository,
		serviceLoger:       serviceLogger,
	}
}

func (ps *ProjectService) CreateProject(project models.Project) error {
	ps.serviceLoger.Info("Creating project...")
	return ps.projectRespository.Create(project)
}

func (ps *ProjectService) ListProjects(userID string) ([]models.Project, error) {
	// Implementation for listing projects
	ps.serviceLoger.Info("Listing projects for user: ", ports.Field{Key: "userID", Value: userID})
	return ps.projectRespository.GetByUserId(userID)
}

func (ps *ProjectService) Delete(projectId string) error {
	ps.serviceLoger.Info("Deleting project")
	return ps.projectRespository.Delete(projectId)
}

func (ps *ProjectService) GetById(projectId string) (*models.Project, error) {
	ps.serviceLoger.Info("Getting by ID")
	return ps.projectRespository.GetById(projectId)
}

func (ps *ProjectService) Update(project models.Project) error {
	ps.serviceLoger.Info("Deleting project")
	return ps.projectRespository.Update(project)
}
