package models

type CalibrationProfile struct {
	ID      string
	RobotID string
	Enabled bool
	Version int64

	Commands map[string]CommandCalibration
}

type CommandCalibration struct {
	CommandName string
	Parameters  map[string]ParamCalibration
	GlobalScale float64
}

type ParamCalibration struct {
	Offset float64
	Scale  float64
	Min    *float64
	Max    *float64
}
