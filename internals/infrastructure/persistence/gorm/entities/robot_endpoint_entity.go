package entities

import (
	"time"

	"auptex.com/botnova/internals/domain/models"
	"gorm.io/gorm"
)

type RobotEndpoint struct {
	Id        string `gorm:"type:text;primaryKey"`
	RobotID   string
	Type      models.TransportType `gorm:"type:text"`
	Address   string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
