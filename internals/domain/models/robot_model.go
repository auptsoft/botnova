package models

import "time"

type RobotModel struct {
	Id        string
	ModelID   string // e.g. "scara_v1"
	ModelName string
	Version   string

	Commands   []CommandDefinition
	Properties []PropertyDefinition

	CreatedAt time.Time
	UpdatedAt time.Time
}
