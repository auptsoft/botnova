package models

type Project struct {
	Id          int
	UserID      int
	Name        string
	Description string
	Options     string //Extra details
}
