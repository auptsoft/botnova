package entities

import (
	"time"

	"auptex.com/botnova/internals/domain/models"
	"gorm.io/gorm"
)

type Script struct {
	Id        string
	ProjectID string
	Name      string
	Type      models.ScriptType

	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt
}
