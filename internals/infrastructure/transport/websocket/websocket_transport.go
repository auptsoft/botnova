package websocket

import (
	"auptex.com/botnova/internals/application/ports"
)

type WebsocketTransport struct {
	EventBus ports.EventBus
}

func NewWebsocketTransport(eventBus ports.EventBus) *WebsocketTransport {
	return &WebsocketTransport{
		EventBus: eventBus,
	}
}

func (wt *WebsocketTransport) SendMessage(evt ports.EventMessage) error {
	event := ports.Event{
		Type:             ports.EventType(evt.Type),
		UserID:           evt.UserID,
		RobotID:          evt.RobotID,
		Payload:          evt.Payload,
		RoutingKey:       evt.RoutingKey,
		Time:             evt.Time,
		EventDestination: ports.WebSocket,
	}

	wt.EventBus.Publish(event)

	return nil
}

func (wt *WebsocketTransport) SubscribeToEvents(handler func(event ports.Event)) error {
	subscription := &ports.Subscription{
		Handler: handler,
		Filter: func(e ports.Event) bool {
			return e.EventSource == ports.WebSocket
		},
	}
	wt.EventBus.SubscribeToAllEvents(subscription)
	return nil
}
