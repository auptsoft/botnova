package models

type GroupMode string

const (
	ModeSinglePrimary  GroupMode = "single_primary" // one active robot
	ModeFanOut         GroupMode = "fan_out"        // send to all
	ModeSimulation     GroupMode = "simulation_only"
	ModePhysicalOnly   GroupMode = "physical_only"
	ModeLeaderFollower GroupMode = "leader_follower" // optional advanced mode
)

type RobotGroup struct {
	ID     string
	UserID string

	Name        string
	Description string

	// Execution behavior
	Mode GroupMode

	// Optional: explicit ordering (useful for leader-follower)
	PrimaryRobotID string

	// Optional metadata
	Tags []string
}

type RobotGroupMember struct {
	GroupID string
	RobotID string

	Role     string // "primary", "secondary", "observer"
	Priority int
}
