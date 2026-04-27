package dtos

type UserDto struct {
	Id    string `json:"Id"`
	Name  string `json:"Name" binding:"required"`
	Email string `json:"Email" binding:"required,email"`
}
