package models

import "time"

type CalibrationProfile struct {
	Id      string
	RobotId string
	Enabled bool
	Version int64

	Commands  map[string]CommandCalibration
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommandCalibration struct {
	CommandName string
	Parameters  map[string]ParamCalibration
	GlobalScale float64
}

type ParamCalibration struct {
	Offset  *float64
	Scale   *float64
	Min     *float64
	Max     *float64
	Append  *string // For string parameters, e.g. "mm" to append to a distance value
	Prepend *string // For string parameters, e.g. "move " to prepend to a command value
}
