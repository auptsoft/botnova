package entities

import (
	"time"

	"auptex.com/botnova/internals/domain/models"
	"gorm.io/gorm"
)

type RobotGroup struct {
	Id     string
	UserId string

	Name        string
	Description string

	// Execution behavior
	Mode models.GroupMode

	// Optional: explicit ordering (useful for leader-follower)
	PrimaryRobotId string

	// Optional metadata
	Options string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type RobotGroupMember struct {
	Id      string
	GroupId string
	RobotId string

	Role     string // "primary", "secondary", "observer"
	Priority int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
