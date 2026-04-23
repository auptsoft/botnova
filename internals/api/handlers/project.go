package handlers

import (
	"auptex.com/botnova/internals/api/dtos"
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
// @Param 		 request body dtos.ProjectDto true "Project Data"
// @Router       /api/project [post]
func (ph *ProjectHandler) CreateProject(c *gin.Context) {
	var req dtos.ProjectDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}

	err := ph.projectService.CreateProject(models.Project{
		UserId:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
		Options:     req.Options,
	})
	if err != nil {
		c.JSON(500, gin.H{"isSuccessful": false, "message": "Failed to create project"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Project created successfully"})
}

// @Summary      Update Project
// @Description  Update existing project
// @Tags         project
// @Accept       json
// @Produce      json
// @Param 		 request body dtos.ProjectDto true "Project Data"
// @Router       /api/project [put]
func (ph *ProjectHandler) UpdateProject(c *gin.Context) {
	var req dtos.ProjectDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}

	err := ph.projectService.Update(models.Project{
		UserId:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
		Options:     req.Options,
	})
	if err != nil {
		c.JSON(500, gin.H{"isSuccessful": false, "message": "Failed to update project"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Project updated successfully"})
}

// @Summary      Get Project by ID
// @Description  Retrieve a project by its ID
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        id path int true "Project ID"
// @Router       /api/project/{id} [get]
func (ph *ProjectHandler) GetByID(c *gin.Context) {
	idStr, ok := c.Params.Get("id")

	if !ok {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request."})
		return
	}

	data, err := ph.projectService.GetById(idStr)
	if err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "An error occurred while deleting"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Success", "data": data})
}

// @Summary      List All Projects
// @Description  Retrieve all projects
// @Tags         project
// @Accept       json
// @Produce      json
// @Router       /api/project [get]
func (ph *ProjectHandler) ListProjects(c *gin.Context) {
	userId := c.GetString("user_id")
	result, error := ph.projectService.ListProjects(userId)

	if error != nil {
		c.JSON(500, gin.H{"isSuccessful": false, "message": "Failed to list projects"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Success", "data": result})
}

// @Summary      Delete Project
// @Description  Delete a project by its ID
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        id path int true "Project ID"
// @Router       /api/project/{id} [delete]
func (ph *ProjectHandler) Delete(c *gin.Context) {
	idStr, ok := c.Params.Get("id")

	if !ok {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request."})
		return
	}

	err := ph.projectService.Delete(idStr)
	if err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "An error occurred while deleting"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Deleted successfully"})
}
