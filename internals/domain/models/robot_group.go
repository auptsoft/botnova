package models

import "time"

type GroupMode string

const (
	ModeSinglePrimary  GroupMode = "single_primary" // one active robot
	ModeFanOut         GroupMode = "fan_out"        // send to all
	ModeSimulation     GroupMode = "simulation_only"
	ModePhysicalOnly   GroupMode = "physical_only"
	ModeLeaderFollower GroupMode = "leader_follower" // optional advanced mode
)

type RobotGroup struct {
	Id     string
	UserId string

	Name        string
	Description string

	// Execution behavior
	Mode GroupMode

	// Optional: explicit ordering (useful for leader-follower)
	PrimaryRobotId string

	// Optional metadata
	Options   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RobotGroupMember struct {
	Id      string
	GroupId string
	RobotId string

	Role      string // "primary", "secondary", "observer"
	Priority  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
