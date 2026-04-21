package bus

import "time"

type BackPressureStrategy int

const (
	DropIfFull BackPressureStrategy = iota
	BlockIfFull
	TimeoutIfFull
)

type PoolConfig struct {
	MaxWorkers           int
	QueueSize            int
	BackPressureStrategy BackPressureStrategy
	Timeout              time.Duration // Used only for TimeoutIfFull strategy
}

type UserPoolConfig struct {
	PoolConfig
	MaxUsers        int
	IdleTTL         time.Duration
	CleanupInterval time.Duration
}
