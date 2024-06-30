# go-events

The `go-events` library is a Go package designed for managing and dispatching events using the AMQP protocol. It provides a mechanism to register event handlers, dispatch events, and ensure concurrent execution of handlers through goroutines and sync.WaitGroup.

## Features

- Register event handlers for specific events.
- Dispatch events to registered handlers concurrently.
- Manage event handlers (add, check, remove).
- Clear all registered handlers.

## Usage

### Create Event Dispatcher

The `NewEventDispatcher` function creates a new instance of `EventDispatcher`.

```go
package main

import (
	"fmt"
	events "libs/golang/shared/go-events/amqp_events"
)

func main() {
	dispatcher := events.NewEventDispatcher()
	fmt.Println("Event Dispatcher created:", dispatcher)
}
```

### Register Event Handler

The `Register` method adds a new handler for a specific event name. It returns an error if the handler is already registered for the event.

```go
err := dispatcher.Register("event.name", handler)
if err != nil {
    fmt.Println("Error registering handler:", err)
}
```

### Dispatch Event

The `Dispatch` method sends an event to all registered handlers for the event's name. It uses goroutines and a `sync.WaitGroup` to handle concurrent execution of handlers.

```go
err := dispatcher.Dispatch(event, "exchangeName", "routingKey")
if err != nil {
    fmt.Println("Error dispatching event:", err)
}
```

### Check Handler Registration

The `Has` method checks if a specific handler is registered for an event name.

```go
if dispatcher.Has("event.name", handler) {
    fmt.Println("Handler is registered")
} else {
    fmt.Println("Handler is not registered")
}
```

### Remove Event Handler

The `Remove` method deletes a handler for a specific event name. It returns nil if the handler is successfully removed or if the handler was not found.

```go
err := dispatcher.Remove("event.name", handler)
if err != nil {
    fmt.Println("Error removing handler:", err)
}
```

### Clear All Handlers

The `Clear` method removes all registered handlers.

```go
dispatcher.Clear()
fmt.Println("All handlers cleared")
```

## Interfaces

### EventInterface

Defines the methods that an event should implement.

```go
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}
```

### EventHandlerInterface

Defines the method that an event handler should implement.

```go
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string)
}
```

### EventDispatcherInterface

Defines the methods that an event dispatcher should implement.

```go
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface, exchangeName string, routingKey string) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}
```

## Testing

To run the tests for the `go-events` package, use the following command:

```sh
npx nx test libs-golang-shared-go-events
```
