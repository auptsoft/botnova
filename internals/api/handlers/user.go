package handlers

import (
	"strconv"

	"auptex.com/botnova/internals/api/dtos"
	"auptex.com/botnova/internals/application/services"
	"auptex.com/botnova/internals/domain/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Summary      Create User
// @Description  Create a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 request body dtos.UserDto true "User Data"
// @Router       /api/user [post]
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var req dtos.UserDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}

	err := uh.userService.CreateUser(models.User{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		c.JSON(500, gin.H{"isSuccessful": false, "message": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "User created successfully"})
}

// @Summary      Update User
// @Description  Update existing user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 request body dtos.UserDto true "User Data"
// @Router       /api/user [put]
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var req dtos.UserDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}

	err := uh.userService.Update(models.User{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		c.JSON(500, gin.H{"isSuccessful": false, "message": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "User updated successfully"})
}

// @Summary      Get User by ID
// @Description  Retrieve a user by its ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Router       /api/user/{id} [get]
func (uh *UserHandler) GetByID(c *gin.Context) {
	idStr, ok := c.Params.Get("id")

	if !ok {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request."})
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request id"})
		return
	}

	data, err := uh.userService.GetById(id)
	if err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "An error occurred while fetching user"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Success", "data": data})
}

// @Summary      Delete User
// @Description  Delete a user by its ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Router       /api/user/{id} [delete]
func (uh *UserHandler) Delete(c *gin.Context) {
	idStr, ok := c.Params.Get("id")

	if !ok {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request."})
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request id"})
		return
	}

	err = uh.userService.Delete(id)
	if err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "An error occurred while deleting"})
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Deleted successfully"})
}
