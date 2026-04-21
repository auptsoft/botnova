package websocket

type Message struct {
	Type             string      `json:"type"`
	RobotID          string      `json:"robot_id,omitempty"`
	Payload          interface{} `json:"payload"`
	RoutingKey       string      `json:"routing_key,omitempty"`
	Time             int64       `json:"time,omitempty"`
	EventDestination string      `json:"event_destination"`
}
