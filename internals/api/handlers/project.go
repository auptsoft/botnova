package handlers

import (
	"auptex.com/botnova/internals/application/services"
	"auptex.com/botnova/internals/domain/models"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	// Add any dependencies like services here
	projectService *services.ProjectService
}

func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// @Summary      Create Project
// @Description  Create a new project
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "User ID"
// @Router       /api/project [post]
func (ph *ProjectHandler) CreateProject(c *gin.Context) {
	var req models.Project

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := ph.projectService.CreateProject(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(200, gin.H{"message": "Project created successfully"})
}

func (ph *ProjectHandler) ListProjects(c *gin.Context) {
	result, error := ph.projectService.ListProjects(0)
	if error != nil {
		c.JSON(500, gin.H{"error": "Failed to list projects"})
		return
	}

	c.JSON(200, gin.H{"projects": result})
}
