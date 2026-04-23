package models

import "time"

type Project struct {
	Id          string
	UserId      string
	Name        string
	Description string
	Options     string //Extra details
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
