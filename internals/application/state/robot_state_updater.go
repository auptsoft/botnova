package state

import (
	"auptex.com/botnova/internals/application/ports"
	repositorydefinitions "auptex.com/botnova/internals/application/ports/repository_definitions"
	"auptex.com/botnova/internals/domain/models"
)

type RobotStateUpdater struct {
	robotGroupRepo repositorydefinitions.RobotGroupRepository
	stateStore     ports.StateStore
	logger         ports.Logger
}

func NewRobotStateUpdater(robotGroupRepo repositorydefinitions.RobotGroupRepository, stateStore ports.StateStore, logger ports.Logger) *RobotStateUpdater {
	return &RobotStateUpdater{
		robotGroupRepo: robotGroupRepo,
		stateStore:     stateStore,
		logger:         logger,
	}
}

func (u *RobotStateUpdater) Update(robotId string, partialData map[string]interface{}) {

	existingState, exists := u.stateStore.GetRobotState(robotId)
	if !exists {
		u.logger.With(ports.Field{Key: "robotId", Value: robotId}).Error("No existing state found for robot, creating new state")

		//Create new state if it doesn't exist
		existingState = &models.RobotState{
			RobotId: robotId,
			Data:    make(map[string]interface{}),
		}
	}

	// Merge partial data into existing state
	for k, v := range partialData {
		existingState.Data[k] = v
	}

	u.updateFullState(existingState)
}

func (u *RobotStateUpdater) updateFullState(state *models.RobotState) {

	group, err := u.robotGroupRepo.GetByRobotId(state.RobotId)
	if group == nil || err != nil {
		u.logger.With(ports.Field{Key: "robotId", Value: state.RobotId}).Error("Failed to find group for robot")
		return
	}

	switch group.Mode {

	case models.ModeSinglePrimary:
		if state.RobotId == group.PrimaryRobotId {
			u.stateStore.SetGroupState(group.Id, state)
		}

	default: // fan-out and others
		u.stateStore.SetGroupState(group.Id, state)
	}

	// Always update robot-level state
	u.stateStore.SetRobotState(state)
}
