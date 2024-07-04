package listener

import (
	usecaseprotocol "libs/golang/server/events/usecase-impl/protocol"
	"sync"

	"errors"
)

// Listener represents a consumer and its associated use case protocol.
type Listener struct {
	consumer        ConsumerInterface
	usecaseProtocol usecaseprotocol.UseCaseProtocol
}

// EventListener manages a collection of listeners.
type EventListener struct {
	listeners map[string]*Listener
	mu        sync.RWMutex
}

// NewEventListener creates a new instance of EventListener.
//
// Returns:
//   - A new instance of EventListener.
func NewEventListener() *EventListener {
	return &EventListener{
		listeners: make(map[string]*Listener),
	}
}

// AddListener adds a new listener to the EventListener.
//
// Parameters:
//   - consumer: The consumer to be added.
//   - usecaseProtocol: The use case protocol associated with the consumer.
//
// Returns:
//   - An error if the listener already exists, otherwise nil.
func (c *EventListener) AddListener(consumer ConsumerInterface, usecaseProtocol usecaseprotocol.UseCaseProtocol) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.listeners[consumer.GetListenerTag()]
	if ok {
		return errors.New("Listener already exists")
	}
	c.listeners[consumer.GetListenerTag()] = &Listener{
		consumer:        consumer,
		usecaseProtocol: usecaseProtocol,
	}
	return nil
}

// RemoveListener removes a listener from the EventListener.
//
// Parameters:
//   - listenerTag: The tag of the listener to be removed.
//
// Returns:
//   - An error if the listener is not found, otherwise nil.
func (c *EventListener) RemoveListener(listenerTag string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.listeners[listenerTag]
	if !ok {
		return errors.New("Listener not found")
	}
	delete(c.listeners, listenerTag)
	return nil
}

// GetListener retrieves a listener by its tag.
//
// Parameters:
//   - listenerTag: The tag of the listener to be retrieved.
//
// Returns:
//   - The listener if found, otherwise an error.
func (c *EventListener) GetListener(listenerTag string) (*Listener, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	listener, ok := c.listeners[listenerTag]
	if !ok {
		return nil, errors.New("Listener not found")
	}
	return listener, nil
}

// GetListeners retrieves all listeners.
//
// Returns:
//   - A map of all listeners.
func (c *EventListener) GetListeners() map[string]*Listener {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.listeners
}

// StartListener starts a listener by its tag.
//
// Parameters:
//   - listenerTag: The tag of the listener to be started.
//
// Returns:
//   - An error if the listener is not found, otherwise nil.
func (c *EventListener) StartListener(listenerTag string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	listener, ok := c.listeners[listenerTag]
	if !ok {
		return errors.New("Listener not found")
	}
	go func(listener *Listener) {
		go listener.consumer.Consume()
		listener.usecaseProtocol.ProcessMessageChannel(listener.consumer.GetMsgCh(), listenerTag)
	}(listener)
	return nil
}
