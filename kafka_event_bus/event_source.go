package kafka_event_bus

import "context"

// This file is extracted from coin_server legacy code

// EventBus defines an interface to deal with events delivery & subscription
type EventBus interface {
	Send(ctx context.Context, e ...interface{}) error
	Subscribe(topic string, handler func(event EventMessage) error)
}

// EventMessage specifies event strcuture
type EventMessage struct {
	Type uint
	Data []byte
}

// EventWithKey specifies event with key implementation
type EventWithKey interface {
	EventKey() string
}

// MessageBuilder ...
type MessageBuilder func()
