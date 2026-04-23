package models

type RobotState struct {
	RobotID string
	Data    map[string]interface{}
	Time    int64
}
