package entities

import (
	"time"

	"gorm.io/gorm"
)

type CalibrationEntity struct {
	Id      string `gorm:"primaryKey"`
	RobotId string `gorm:"uniqueIndex"`

	Enabled bool
	Version int64

	// JSON blob for flexibility
	CommandsJSON string `gorm:"type:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
