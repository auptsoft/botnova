package common

import (
	"time"

	"auptex.com/botnova/internals/infrastructure/bus"
)

func GetEventBusConfigs() (bus.UserPoolConfig, bus.UserPoolConfig, bus.UserPoolConfig) {
	cmdConfig := bus.UserPoolConfig{
		PoolConfig: bus.PoolConfig{
			MaxWorkers:           10,
			QueueSize:            100,
			BackPressureStrategy: bus.DropIfFull,
		},
		MaxUsers:        1000,
		IdleTTL:         5 * time.Minute,
		CleanupInterval: 1 * time.Minute,
	}

	stateConfig := bus.UserPoolConfig{
		PoolConfig: bus.PoolConfig{
			MaxWorkers:           5,
			QueueSize:            50,
			BackPressureStrategy: bus.BlockIfFull,
		},
		MaxUsers:        1000,
		IdleTTL:         5 * time.Minute,
		CleanupInterval: 1 * time.Minute,
	}

	defaultConfig := bus.UserPoolConfig{
		PoolConfig: bus.PoolConfig{
			MaxWorkers:           2,
			QueueSize:            20,
			BackPressureStrategy: bus.TimeoutIfFull,
			Timeout:              1 * time.Second,
		},
		MaxUsers:        1000,
		IdleTTL:         5 * time.Minute,
		CleanupInterval: 1 * time.Minute,
	}

	return cmdConfig, stateConfig, defaultConfig
}
