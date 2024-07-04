package eventmock

import (
	"time"

	events "libs/golang/shared/go-events/amqp_events"

	"github.com/stretchr/testify/mock"
)

// MockEvent is a mock implementation of EventInterface used for testing purposes.
type MockEvent struct {
	mock.Mock
}

// GetName returns the name of the mock event.
//
// Returns:
//   - A string representing the name of the event.
func (m *MockEvent) GetName() string {
	args := m.Called()
	return args.String(0)
}

// GetDateTime returns the date and time of the mock event.
//
// Returns:
//   - A time.Time value representing the date and time of the event.
func (m *MockEvent) GetDateTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

// GetPayload returns the payload of the mock event.
//
// Returns:
//   - An interface{} representing the payload of the event.
func (m *MockEvent) GetPayload() interface{} {
	args := m.Called()
	return args.Get(0)
}

// SetPayload sets the payload of the mock event.
//
// Parameters:
//   - payload: An interface{} representing the payload to be set.
func (m *MockEvent) SetPayload(payload interface{}) {
	m.Called(payload)
}

// MockEventDispatcher is a mock implementation of EventDispatcherInterface used for testing purposes.
type MockEventDispatcher struct {
	mock.Mock
}

// Register registers an event handler for a specific event name in the mock dispatcher.
//
// Parameters:
//   - eventName: A string representing the name of the event.
//   - handler: An instance of EventHandlerInterface to handle the event.
//
// Returns:
//   - An error if registration fails, or nil if successful.
func (m *MockEventDispatcher) Register(eventName string, handler events.EventHandlerInterface) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

// Dispatch dispatches an event using the mock dispatcher.
//
// Parameters:
//   - event: An instance of EventInterface representing the event to be dispatched.
//   - routingKey: A string representing the routing key for the event.
//
// Returns:
//   - An error if dispatching fails, or nil if successful.
func (m *MockEventDispatcher) Dispatch(event events.EventInterface, routingKey string) error {
	args := m.Called(event, routingKey)
	return args.Error(0)
}

// Remove removes an event handler for a specific event name in the mock dispatcher.
//
// Parameters:
//   - eventName: A string representing the name of the event.
//   - handler: An instance of EventHandlerInterface to be removed.
//
// Returns:
//   - An error if removal fails, or nil if successful.
func (m *MockEventDispatcher) Remove(eventName string, handler events.EventHandlerInterface) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

// Has checks if a specific event handler is registered for an event name in the mock dispatcher.
//
// Parameters:
//   - eventName: A string representing the name of the event.
//   - handler: An instance of EventHandlerInterface to check for registration.
//
// Returns:
//   - A boolean indicating whether the handler is registered.
func (m *MockEventDispatcher) Has(eventName string, handler events.EventHandlerInterface) bool {
	args := m.Called(eventName, handler)
	return args.Bool(0)
}

// Clear clears all event handlers in the mock dispatcher.
func (m *MockEventDispatcher) Clear() {
	m.Called()
}
