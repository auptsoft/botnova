package state

import (
	"auptex.com/botnova/internals/application/ports"
	"auptex.com/botnova/internals/domain/models"
)

type RobotGroupRepository interface {
	GetById(id string) (*models.RobotGroup, error)
}

type GroupStateUpdater struct {
	robotGroupRepo RobotGroupRepository
	stateStore     ports.StateStore
}

func NewGroupStateUpdater(robotGroupRepo RobotGroupRepository, stateStore ports.StateStore) *GroupStateUpdater {
	return &GroupStateUpdater{
		robotGroupRepo: robotGroupRepo,
		stateStore:     stateStore,
	}
}

func (u *GroupStateUpdater) Update(state *models.RobotState) {

	if state.GroupId == "" {
		return
	}

	group, err := u.robotGroupRepo.GetById(state.GroupId)
	if group == nil || err != nil {
		return
	}

	switch group.Mode {

	case "single_primary":
		if state.RobotId == group.PrimaryRobotId {
			u.stateStore.SetGroupState(group.Id, state.Data)
		}

	default: // fan-out and others
		u.stateStore.SetGroupState(group.Id, state.Data)
	}
}
