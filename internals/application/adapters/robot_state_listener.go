package adapters

import (
	"auptex.com/botnova/internals/application/ports"
	"auptex.com/botnova/internals/application/state"
)

type RobotStateListener struct {
	// This adapter listens to robot state updates and updates the group state accordingly
	groupStateUpdater *state.RobotStateUpdater
	eventBus          ports.EventBus
	logger            ports.Logger
}

func NewRobotStateListener(groupStateUpdater *state.RobotStateUpdater, eventBus ports.EventBus, logger ports.Logger) *RobotStateListener {
	return &RobotStateListener{
		groupStateUpdater: groupStateUpdater,
		eventBus:          eventBus,
		logger:            logger,
	}
}

func (rsl *RobotStateListener) Start() {
	rsl.logger.Info("Starting RobotStateListener")
	subscription := ports.Subscription{
		ID: "robot-state-listener",
		Handler: func(e ports.Event) {
			rsl.logger.With(ports.Field{Key: "robotId", Value: e.RobotID}).Info("Received state update event for robot")
			rsl.groupStateUpdater.Update(e.RobotID, e.Payload.(map[string]interface{}))
		},
		Filter: func(e ports.Event) bool {
			return e.Type == ports.StateEvent && e.RobotID != "" && e.Payload != nil
		},
	}

	rsl.eventBus.Subscribe(ports.StateEvent, &subscription)
	rsl.logger.Info("RobotStateListener started")
}
