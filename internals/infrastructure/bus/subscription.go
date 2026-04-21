package bus

import (
	"context"
	"sync"

	"auptex.com/botnova/internals/application/ports"
	"auptex.com/botnova/internals/infrastructure/logger"
)

type HandlerRegistry struct {
	mu            sync.RWMutex
	subscriptions map[ports.SubscriptionID]*ports.Subscription
}

func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		subscriptions: make(map[ports.SubscriptionID]*ports.Subscription),
	}
}

func (r *HandlerRegistry) Subscribe(s *ports.Subscription) ports.SubscriptionID {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.subscriptions[s.ID] = s
	return s.ID
}

func (r *HandlerRegistry) Unsubscribe(id ports.SubscriptionID) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.subscriptions, id)
}

func (r *HandlerRegistry) Dispatch(event ports.Event, subLogger ports.Logger) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	log := subLogger.With(
		ports.Field{Key: "event_type", Value: event.Type},
		ports.Field{Key: "user_id", Value: event.UserID},
		ports.Field{Key: "routing_key", Value: event.RoutingKey},
		ports.Field{Key: "trace_id", Value: event.TraceID},
	)

	event.Ctx = logger.WithContext(context.Background(), log)

	for _, sub := range r.subscriptions {
		if sub.Filter == nil || sub.Filter(event) {
			sub.Handler(event)
		}
	}
}
