package models

type Command struct {
	ID      string
	RobotID string
	UserID  string

	Name   string
	Params map[string]interface{}

	Timestamp int64
}

type CommandResult struct {
	CommandID string
	Success   bool
	Message   string
	Data      map[string]interface{}
}
