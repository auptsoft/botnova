package ports

import "context"

type EventType string

const (
	CmdEvent     EventType = "command"
	StateEvent   EventType = "state"
	DefaultEvent EventType = "default"
)

type EventLocation string

const (
	WebSocket EventLocation = "websocket"
	Zenoh     EventLocation = "zenoh"
	Mqtt      EventLocation = "mqtt"

	OnBus EventLocation = "bus"
)

type Event struct {
	Type       EventType
	RoutingKey string
	UserID     string
	RobotID    string
	Payload    any
	Time       int64

	TraceID string
	Ctx     context.Context

	EventSource      EventLocation
	EventDestination EventLocation
	IsBroadcast      bool
}

type Handler func(Event)
