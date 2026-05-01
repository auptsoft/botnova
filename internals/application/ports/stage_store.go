package ports

import "auptex.com/botnova/internals/domain/models"

type StateStore interface {
	// Robot-level
	GetRobotState(robotID string) (*models.RobotState, bool)
	SetRobotState(state *models.RobotState)

	// Group-level
	GetGroupState(groupID string) (map[string]interface{}, bool)
	SetGroupState(groupID string, state map[string]interface{})

	// Subscription (optional for future)
	SubscribeRobot(robotID string, ch chan *models.RobotState)
}
