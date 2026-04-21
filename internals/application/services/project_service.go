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

func (ps *ProjectService) ListProjects(userID int) ([]models.Project, error) {
	// Implementation for listing projects
	ps.serviceLoger.Info("Listing projects for user: ", ports.Field{Key: "userID", Value: userID})
	return ps.projectRespository.GetByUserId(userID)
}
