package handlers

import (
	"encoding/json"
	"errors"
	"strings"

	"auptex.com/botnova/internals/api/dtos"
	"auptex.com/botnova/internals/application/services"
	"auptex.com/botnova/internals/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func ToUserDto(user models.User) dtos.UserDto {
	return dtos.UserDto{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

func handleUserServiceError(c *gin.Context, err error, fallbackMessage string) {
	switch {
	case errors.Is(err, services.ErrEmailAlreadyExists):
		c.JSON(409, gin.H{"isSuccessful": false, "message": "Email already exists"})
	case errors.Is(err, services.ErrInvalidCredentials):
		c.JSON(401, gin.H{"isSuccessful": false, "message": "Invalid email or password"})
	case errors.Is(err, services.ErrCurrentPasswordInvalid):
		c.JSON(401, gin.H{"isSuccessful": false, "message": "Current password is incorrect"})
	case errors.Is(err, services.ErrPasswordUnchanged):
		c.JSON(400, gin.H{"isSuccessful": false, "message": "New password must be different from current password"})
	case errors.Is(err, services.ErrUserNotFound):
		c.JSON(404, gin.H{"isSuccessful": false, "message": "User not found"})
	default:
		c.JSON(500, gin.H{"isSuccessful": false, "message": fallbackMessage})
	}
}

// @Summary      Sign Up User
// @Description  Register a new user account
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 request body dtos.UserSignupDto true "Signup Data"
// @Router       /api/user/signup [post]
func (uh *UserHandler) SignUp(c *gin.Context) {
	var req dtos.UserSignupDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}

	result, err := uh.userService.SignUp(models.User{
		Name:  req.Name,
		Email: req.Email,
	}, req.Password)
	if err != nil {
		handleUserServiceError(c, err, "Failed to sign up user")
		return
	}

	c.JSON(201, gin.H{
		"isSuccessful": true,
		"message":      "Signup successful",
		"data": gin.H{
			"token": result.Token,
			"user":  ToUserDto(result.User),
		},
	})
}

// @Summary      Login User
// @Description  Authenticate a user and return a JWT
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 request body dtos.UserLoginDto true "Login Data"
// @Router       /api/user/login [post]
func (uh *UserHandler) Login(c *gin.Context) {
	var req dtos.UserLoginDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}

	result, err := uh.userService.Login(req.Email, req.Password)
	if err != nil {
		handleUserServiceError(c, err, "Failed to login")
		return
	}

	c.JSON(200, gin.H{
		"isSuccessful": true,
		"message":      "Login successful",
		"data": gin.H{
			"token": result.Token,
			"user":  ToUserDto(result.User),
		},
	})
}

// @Summary      Get User by ID
// @Description  Retrieve a user by its ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Router       /api/user/{id} [get]
func (uh *UserHandler) GetByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request."})
		return
	}
	authUserID := c.GetString("user_id")
	if authUserID == "" {
		c.JSON(401, gin.H{"isSuccessful": false, "message": "Unauthorized"})
		return
	}
	if id != authUserID {
		c.JSON(403, gin.H{"isSuccessful": false, "message": "Forbidden"})
		return
	}

	data, err := uh.userService.GetById(id)
	if err != nil {
		handleUserServiceError(c, err, "An error occurred while fetching user")
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Success", "data": ToUserDto(*data)})
}

// @Summary      Get Current User
// @Description  Retrieve the authenticated user's profile
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer <token>"
// @Router       /api/user/me [get]
func (uh *UserHandler) GetCurrentUser(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(401, gin.H{"isSuccessful": false, "message": "Unauthorized"})
		return
	}

	data, err := uh.userService.GetById(userID)
	if err != nil {
		handleUserServiceError(c, err, "An error occurred while fetching user")
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Success", "data": ToUserDto(*data)})
}

// @Summary      Update Current User
// @Description  Update the authenticated user's profile
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer <token>"
// @Param 		 request body dtos.UserUpdateDto true "User Data"
// @Router       /api/user/me [put]
func (uh *UserHandler) UpdateCurrentUser(c *gin.Context) {
	var req dtos.UserUpdateDto
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}

	var raw map[string]json.RawMessage
	if err := c.ShouldBindBodyWith(&raw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}
	for key := range raw {
		normalizedKey := strings.ToLower(strings.TrimSpace(key))
		if normalizedKey == "password" || normalizedKey == "currentpassword" || normalizedKey == "newpassword" {
			c.JSON(400, gin.H{"isSuccessful": false, "message": "Use /api/user/me/change-password to change password"})
			return
		}
	}

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(401, gin.H{"isSuccessful": false, "message": "Unauthorized"})
		return
	}

	updatedUser, err := uh.userService.UpdateUser(userID, req.Name, req.Email)
	if err != nil {
		handleUserServiceError(c, err, "Failed to update user")
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "User updated successfully", "data": ToUserDto(*updatedUser)})
}

// @Summary      Change Current User Password
// @Description  Change the authenticated user's password
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer <token>"
// @Param 		 request body dtos.UserChangePasswordDto true "Password Change Data"
// @Router       /api/user/me/change-password [put]
func (uh *UserHandler) ChangePassword(c *gin.Context) {
	var req dtos.UserChangePasswordDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"isSuccessful": false, "message": "Invalid request"})
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(401, gin.H{"isSuccessful": false, "message": "Unauthorized"})
		return
	}

	err := uh.userService.ChangePassword(userID, req.CurrentPassword, req.NewPassword)
	if err != nil {
		handleUserServiceError(c, err, "Failed to change password")
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Password changed successfully"})
}

// @Summary      Delete Current User
// @Description  Delete the authenticated user's account
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer <token>"
// @Router       /api/user/me [delete]
func (uh *UserHandler) DeleteCurrentUser(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(401, gin.H{"isSuccessful": false, "message": "Unauthorized"})
		return
	}

	err := uh.userService.Delete(userID)
	if err != nil {
		handleUserServiceError(c, err, "An error occurred while deleting")
		return
	}

	c.JSON(200, gin.H{"isSuccessful": true, "message": "Deleted successfully"})
}
