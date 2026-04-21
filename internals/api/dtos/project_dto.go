package dtos

type ProjectDto struct {
	UserId      int    `json:"UserId" binding:"required"`
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description"`
	Options     string `json:"Options"` //Extra details
}
