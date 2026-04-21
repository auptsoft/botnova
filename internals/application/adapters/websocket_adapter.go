package adapters

import (
	"time"

	"auptex.com/botnova/internals/application/ports"
	"auptex.com/botnova/internals/infrastructure/transport/websocket"
)

func InitWebsocket(eventBus ports.EventBus, log ports.Logger) *websocket.Server {
	wsServer := websocket.NewServer(log, func(userID string, msg websocket.Message) {
		log.Info("Received data from websocket")

		event := ports.Event{
			Type:             ports.EventType(msg.Type),
			UserID:           userID,
			RobotID:          msg.RobotID,
			Payload:          msg.Payload,
			RoutingKey:       msg.RoutingKey,
			Time:             time.Now().Unix(),
			EventSource:      ports.WebSocket,
			EventDestination: ports.EventLocation(msg.EventDestination),
		}
		eventBus.Publish(event)
	})

	eventBus.SubscribeToAllEvents(&ports.Subscription{
		ID: "ws-event",
		Handler: func(e ports.Event) {
			log.Info("Sending data through websocket.")

			wsServer.SendToUser(e.UserID, websocket.Message{
				Type:             string(e.Type),
				RobotID:          e.RobotID,
				Payload:          e.Payload,
				RoutingKey:       e.RoutingKey,
				Time:             e.Time,
				EventDestination: string(e.EventDestination),
			})
		},
		Filter: func(e ports.Event) bool {
			log.With(ports.Field{Key: "userId", Value: e.UserID}).Debug("Checking if event should be transmitted through websocket")

			if e.UserID == "" {
				log.Error("UserID not added to event")
				return false
			}

			if e.EventDestination != ports.WebSocket {
				log.Info("EventDestination is not WebSocket")
				log.Info(string(e.EventDestination))
				return false
			}

			log.Info("Websocket condition met.")

			return true
		},
	})

	eventBus.SubscribeToAllEvents(&ports.Subscription{
		ID: "ws-event-broadcast",
		Handler: func(e ports.Event) {
			wsServer.Broadcast(websocket.Message{
				Type:       string(e.Type),
				RobotID:    e.RobotID,
				Payload:    e.Payload,
				RoutingKey: e.RoutingKey,
				Time:       e.Time,
			})
		},
		Filter: func(e ports.Event) bool {
			return e.EventDestination == ports.WebSocket && e.IsBroadcast
		},
	})

	return wsServer
}
