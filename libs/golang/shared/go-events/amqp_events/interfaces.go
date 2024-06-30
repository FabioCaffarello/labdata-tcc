package amqpevents

import (
	"sync"
	"time"
)

// EventInterface defines the methods that an event should implement.
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

// EventHandlerInterface defines the method that an event handler should implement.
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string)
}

// EventListenerInterface defines the method that an event listener should implement.
type EventListenerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

// EventDispatcherInterface defines the methods that an event dispatcher should implement.
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface, exchangeName string, routingKey string) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}
