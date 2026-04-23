package entities

import (
	"time"

	"gorm.io/gorm"
)

type RobotModel struct {
	Id        string
	ModelID   string // e.g. "scara_v1"
	ModelName string
	Version   string

	CommandsJson   string
	PropertiesJson string

	CreatedAt time.Time
	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt
}
