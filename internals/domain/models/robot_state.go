package models

type RobotState struct {
	RobotId   string
	GroupId   string
	Data      map[string]interface{}
	TimeStamp int64
}
