# input-broker/event

`input-broker/event` is a Go library designed for handling and dispatching events related to input operations. This library provides mechanisms for creating, registering, and processing events, with a focus on AMQP-based event handling.

## Features

- Create and manage `InputCreated` events.
- Register event handlers for specific events.
- Dispatch events to registered handlers concurrently.
- Notify systems of event occurrences.


## Usage

### Creating an InputCreated Event

The `InputCreated` struct represents an event when a new input is created. Use the `NewInputCreated` function to create a new instance of this event.

```go
package main

import (
	"fmt"
	"libs/golang/ddd/events/input-broker/event"
)

func main() {
	inputCreatedEvent := event.NewInputCreated()
	inputCreatedEvent.SetPayload(map[string]interface{}{
		"inputID": "12345",
		"status":  "created",
	})
	fmt.Println("Event created:", inputCreatedEvent)
}
```

### Handling the InputCreated Event

The `InputCreatedHandler` struct handles events of type `InputCreated`. Implement the `NotifierInterface` to define how notifications should be sent when the event is handled.

```go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"libs/golang/ddd/events/input-broker/event"
	"libs/golang/ddd/events/input-broker/handler"
)

type Notifier struct{}

func (n *Notifier) Notify(message []byte, routingKey string) error {
	fmt.Println("Notification sent:", string(message))
	return nil
}

func main() {
	notifier := &Notifier{}
	handler := handler.NewInputCreatedHandler(notifier)

	event := event.NewInputCreated()
	event.SetPayload(map[string]interface{}{
		"inputID": "12345",
		"status":  "created",
	})

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go handler.Handle(event, wg, "exchangeName", "routingKey")
	wg.Wait()
}
```

### Registering Event Handlers

Register event handlers to handle specific events using the `go-events` library.

```go
package main

import (
	"fmt"
	"libs/golang/ddd/events/input-broker/event"
	"libs/golang/ddd/events/input-broker/handler"
	events "libs/golang/shared/go-events/amqp_events"
)

func main() {
	dispatcher := events.NewEventDispatcher()
	notifier := &Notifier{}
	handler := handler.NewInputCreatedHandler(notifier)

	err := dispatcher.Register("InputCreated", handler)
	if err != nil {
		fmt.Println("Error registering handler:", err)
	}

	event := event.NewInputCreated()
	event.SetPayload(map[string]interface{}{
		"inputID": "12345",
		"status":  "created",
	})

	err = dispatcher.Dispatch(event, "exchangeName", "routingKey")
	if err != nil {
		fmt.Println("Error dispatching event:", err)
	}
}
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

### NotifierInterface

Defines the methods that a notifier should implement.

```go
type NotifierInterface interface {
	Notify(message []byte, routingKey string) error
}
```

## Testing

To run the tests for the `event` package, use the following command:

```sh
npx nx test libs-golang-ddd-events-input-broker
```
