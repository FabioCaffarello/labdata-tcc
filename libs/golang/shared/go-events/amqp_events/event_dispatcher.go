package amqpevents

import (
	"errors"
	"sync"
)

// ErrHandlerAlreadyRegistered is returned when an event handler is already registered for a specific event.
var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

// EventDispatcher is responsible for managing event handlers and dispatching events to them.
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

// NewEventDispatcher creates a new instance of EventDispatcher.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Dispatch sends an event to all registered handlers for the event's name.
// It uses goroutines and a WaitGroup to handle concurrent execution of handlers.
// routingKey are used for routing the event in the AMQP system.
func (ev *EventDispatcher) Dispatch(event EventInterface, routingKey string) error {
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg, routingKey)
		}
		wg.Wait()
	}
	return nil
}

// Register adds a new handler for a specific event name.
// It returns an error if the handler is already registered for the event.
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

// Has checks if a specific handler is registered for an event name.
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

// Remove deletes a handler for a specific event name.
// It returns nil if the handler is successfully removed or if the handler was not found.
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

// Clear removes all registered handlers.
func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}
