package bus

import (
	"auptex.com/botnova/internals/application/ports"
	"github.com/google/uuid"
)

type InMemoryBus struct {
	cmdPools     *UserScopedPool
	statePools   *UserScopedPool
	defaultPools *UserScopedPool

	logger ports.Logger
}

func NewBus(logger ports.Logger,
	cmdConfig, stateConfig, defaultConfig UserPoolConfig,
) *InMemoryBus {
	cmdSubscriptionRegistry := NewHandlerRegistry()
	stateSubscriptionRegistry := NewHandlerRegistry()
	defaultSubscriptionRegistry := NewHandlerRegistry()

	return &InMemoryBus{
		cmdPools:     NewUserScopedPool(logger, cmdConfig, cmdSubscriptionRegistry),
		statePools:   NewUserScopedPool(logger, stateConfig, stateSubscriptionRegistry),
		defaultPools: NewUserScopedPool(logger, defaultConfig, defaultSubscriptionRegistry),
		logger:       logger,
	}
}

func (b *InMemoryBus) Publish(e ports.Event) {
	// fmt.Printf("Publishing event: %+v\n", e)
	e.TraceID = uuid.New().String()
	log := b.logger.With(
		ports.Field{Key: "action", Value: "publish_event"},
		ports.Field{Key: "event_type", Value: e.Type},
		ports.Field{Key: "user_id", Value: e.UserID},
		ports.Field{Key: "routing_key", Value: e.RoutingKey},
		ports.Field{Key: "trace_id", Value: e.TraceID},
	)

	log.Debug("Publishing event")

	switch e.Type {
	case ports.CmdEvent:
		b.cmdPools.Submit(e)

	case ports.StateEvent:
		b.statePools.Submit(e)

	default:
		b.defaultPools.Submit(e)
	}
}

func (b *InMemoryBus) Subscribe(eventType ports.EventType, s *ports.Subscription) ports.SubscriptionID {
	b.logger.With(
		ports.Field{Key: "action", Value: "subscribe_event"},
		ports.Field{Key: "event_type", Value: eventType},
		ports.Field{Key: "subscription_id", Value: s.ID},
	).Debug("Subscribing to event")

	switch eventType {
	case ports.CmdEvent:
		return b.cmdPools.Subscribe(s)
	case ports.StateEvent:
		return b.statePools.Subscribe(s)
	default:
		return b.defaultPools.Subscribe(s)
	}
}

func (b *InMemoryBus) Unsubscribe(eventType ports.EventType, id ports.SubscriptionID) {
	b.logger.With(
		ports.Field{Key: "action", Value: "unsubscribe_event"},
		ports.Field{Key: "event_type", Value: eventType},
		ports.Field{Key: "subscription_id", Value: id},
	).Debug("Unsubscribing from event")

	switch eventType {
	case ports.CmdEvent:
		b.cmdPools.Unsubscribe(id)
	case ports.StateEvent:
		b.statePools.Unsubscribe(id)
	default:
		b.defaultPools.Unsubscribe(id)
	}
}

func (b *InMemoryBus) SubscribeToAllEvents(s *ports.Subscription) ports.SubscriptionID {
	b.logger.With(
		ports.Field{Key: "action", Value: "subscribe_all_events"},
		ports.Field{Key: "subscription_id", Value: s.ID},
	).Debug("Subscribing to all events")

	b.cmdPools.Subscribe(s)
	b.statePools.Subscribe(s)
	b.defaultPools.Subscribe(s)

	return s.ID
}

func (b *InMemoryBus) UnsubscribeFromAllEvents(id ports.SubscriptionID) {
	b.logger.With(
		ports.Field{Key: "action", Value: "unsubscribe_all_events"},
		ports.Field{Key: "subscription_id", Value: id},
	).Debug("Unsubscribing from all events")

	b.cmdPools.Unsubscribe(id)
	b.statePools.Unsubscribe(id)
	b.defaultPools.Unsubscribe(id)
}

func (b *InMemoryBus) SubscribeWithRoutingKey(eventType ports.EventType, routingKey string, s *ports.Subscription) ports.SubscriptionID {
	b.logger.With(
		ports.Field{Key: "action", Value: "subscribe_with_routing_key"},
		ports.Field{Key: "event_type", Value: eventType},
		ports.Field{Key: "routing_key", Value: routingKey},
		ports.Field{Key: "subscription_id", Value: s.ID},
	).Debug("Subscribing to event with routing key")

	s.Filter = func(e ports.Event) bool {
		return e.RoutingKey == routingKey
	}

	return b.Subscribe(eventType, s)
}

func (b *InMemoryBus) SubscribeToAllWithRoutingKey(routingKey string, s *ports.Subscription) ports.SubscriptionID {
	b.logger.With(
		ports.Field{Key: "action", Value: "subscribe_all_with_routing_key"},
		ports.Field{Key: "routing_key", Value: routingKey},
		ports.Field{Key: "subscription_id", Value: s.ID},
	).Debug("Subscribing to all events with routing key")

	s.Filter = func(e ports.Event) bool {
		return e.RoutingKey == routingKey
	}

	return b.SubscribeToAllEvents(s)
}
