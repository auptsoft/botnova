package ports

type EventMessage struct {
	Type       EventType
	RoutingKey string
	UserID     string
	RobotID    string
	Payload    any
	Time       int64
}

type EventTransport interface {
	SendMessage(eventMsg EventMessage) error
	SubscribeToEvents(handler func(event Event)) error
}
