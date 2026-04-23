package models

import "time"

type RobotType string

const (
	RobotTypePhysical  RobotType = "physical"
	RobotTypeSimulated RobotType = "simulated"
)

type Robot struct {
	Id      string
	Name    string
	ModelID string
	UserID  string
	Type    RobotType

	Status    string // "online", "offline"
	Endpoints []RobotEndpoint
	CreatedAt time.Time
	UpdatedAt time.Time
}
