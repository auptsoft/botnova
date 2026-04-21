package entities

import "time"

type Project struct {
	Id          int `gorm:"primaryKey"`
	UserID      int
	Name        string
	Description string
	Options     string //Extra details
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
