package ports

type EventBus interface {
	Publish(event Event)
	Subscribe(eventType EventType, s *Subscription) SubscriptionID
	Unsubscribe(eventType EventType, id SubscriptionID)
	SubscribeToAllEvents(s *Subscription) SubscriptionID
	UnsubscribeFromAllEvents(id SubscriptionID)
	SubscribeWithRoutingKey(eventType EventType, routingKey string, s *Subscription) SubscriptionID
	SubscribeToAllWithRoutingKey(routingKey string, s *Subscription) SubscriptionID
}

type SubscriptionID string

type Subscription struct {
	ID      SubscriptionID
	Handler Handler
	Filter  func(Event) bool
}
