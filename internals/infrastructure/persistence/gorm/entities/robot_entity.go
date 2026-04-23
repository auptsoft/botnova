package entities

import (
	"time"

	"auptex.com/botnova/internals/domain/models"
)

type Robot struct {
	Id      string
	Name    string
	ModelID string
	UserID  string
	Type    models.RobotType

	Status    string // "online", "offline"
	Endpoints []RobotEndpoint
	CreatedAt time.Time
	UpdatedAt time.Time
}
