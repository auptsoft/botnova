package ports

import "auptex.com/botnova/internals/domain/models"

type StateStore interface {
	// Robot-level
	GetRobotState(robotID string) (*models.RobotState, bool)
	SetRobotState(state *models.RobotState)

	// Group-level
	GetGroupState(groupID string) (*models.RobotState, bool)
	SetGroupState(groupID string, state *models.RobotState)
}
