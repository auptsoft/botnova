package dtos

type UserDto struct {
	Id    string `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type UserSignupDto struct {
	Name     string `json:"Name" binding:"required"`
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"Password" binding:"required,min=8"`
}

type UserLoginDto struct {
	Email    string `json:"Email" binding:"required,email"`
	Password string `json:"Password" binding:"required"`
}

type UserUpdateDto struct {
	Name     string `json:"Name"`
	Email    string `json:"Email" binding:"omitempty,email"`
	Password string `json:"Password" binding:"omitempty,min=8"`
}
