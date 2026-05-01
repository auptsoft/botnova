package models

type RobotState struct {
	RobotId   string
	Data      map[string]interface{}
	TimeStamp int64
}
