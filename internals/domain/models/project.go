package models

import "time"

type Project struct {
	Id          int
	UserId      int
	Name        string
	Description string
	Options     string //Extra details
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
