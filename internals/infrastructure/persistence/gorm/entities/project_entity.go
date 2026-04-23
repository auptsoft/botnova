package entities

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	Id          string `gorm:"type:text;primaryKey"`
	UserID      string
	Name        string
	Description string
	Options     string //Extra details
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
