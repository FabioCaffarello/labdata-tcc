# events-router/event

`events-router/event` is a Go library designed for handling and dispatching events related to event order operations within the `events-router` domain. This library provides mechanisms for creating, registering, and processing events, with a focus on AMQP-based event handling.

## Features

- Create and manage `ErrorCreated` and `OrderedProcess` events.
- Register event handlers for specific events.
- Dispatch events to registered handlers concurrently.
- Notify systems of event occurrences.

## Usage

### Creating an ErrorCreated Event

The `ErrorCreated` struct represents an event when an error is created. Use the `NewErrorCreated` function to create a new instance of this event.

```go
package main

import (
	"fmt"
	"libs/golang/ddd/events/events-router/event"
)

func main() {
	errorCreatedEvent := event.NewErrorCreated()
	errorCreatedEvent.SetPayload(map[string]interface{}{
		"errorID":   "12345",
		"errorCode": "500",
	})
	fmt.Println("Event created:", errorCreatedEvent)
}
```

### Handling the ErrorCreated Event

The `ErrorCreatedHandler` struct handles events of type `ErrorCreated`. Implement the `NotifierInterface` to define how notifications should be sent when the event is handled.

```go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"libs/golang/ddd/events/events-router/event"
	"libs/golang/ddd/events/events-router/handler"
)

type Notifier struct{}

func (n *Notifier) Notify(message []byte, routingKey string) error {
	fmt.Println("Notification sent:", string(message))
	return nil
}

func main() {
	notifier := &Notifier{}
	handler := handler.NewErrorCreatedHandler(notifier)

	event := event.NewErrorCreated()
	event.SetPayload(map[string]interface{}{
		"errorID":   "12345",
		"errorCode": "500",
	})

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go handler.Handle(event, wg, "exchangeName", "routingKey")
	wg.Wait()
}
```

### Creating an OrderedProcess Event

The `OrderedProcess` struct represents an event when an order is processed. Use the `NewOrderedProcess` function to create a new instance of this event.

```go
package main

import (
	"fmt"
	"libs/golang/ddd/events/events-router/event"
)

func main() {
	orderedProcessEvent := event.NewOrderedProcess()
	orderedProcessEvent.SetPayload(map[string]interface{}{
		"orderID": "54321",
		"status":  "processed",
	})
	fmt.Println("Event created:", orderedProcessEvent)
}
```

### Handling the OrderedProcess Event

The `OrderedProcessHandler` struct handles events of type `OrderedProcess`. Implement the `NotifierInterface` to define how notifications should be sent when the event is handled.

```go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"libs/golang/ddd/events/events-router/event"
	"libs/golang/ddd/events/events-router/handler"
)

type Notifier struct{}

func (n *Notifier) Notify(message []byte, routingKey string) error {
	fmt.Println("Notification sent:", string(message))
	return nil
}

func main() {
	notifier := &Notifier{}
	handler := handler.NewOrderedProcessHandler(notifier)

	event := event.NewOrderedProcess()
	event.SetPayload(map[string]interface{}{
		"orderID": "54321",
		"status":  "processed",
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
	"libs/golang/ddd/events/events-router/event"
	"libs/golang/ddd/events/events-router/handler"
	events "libs/golang/shared/go-events/amqp_events"
)

func main() {
	dispatcher := events.NewEventDispatcher()
	notifier := &Notifier{}
	errorHandler := handler.NewErrorCreatedHandler(notifier)
	orderedHandler := handler.NewOrderedProcessHandler(notifier)

	err := dispatcher.Register("ErrorCreated", errorHandler)
	if err != nil {
		fmt.Println("Error registering error handler:", err)
	}

	err = dispatcher.Register("OrderedProcess", orderedHandler)
	if err != nil {
		fmt.Println("Error registering ordered handler:", err)
	}

	errorEvent := event.NewErrorCreated()
	errorEvent.SetPayload(map[string]interface{}{
		"errorID":   "12345",
		"errorCode": "500",
	})

	orderedEvent := event.NewOrderedProcess()
	orderedEvent.SetPayload(map[string]interface{}{
		"orderID": "54321",
		"status":  "processed",
	})

	err = dispatcher.Dispatch(errorEvent, "exchangeName", "routingKey")
	if err != nil {
		fmt.Println("Error dispatching error event:", err)
	}

	err = dispatcher.Dispatch(orderedEvent, "exchangeName", "routingKey")
	if err != nil {
		fmt.Println("Error dispatching ordered event:", err)
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
npx nx test libs-golang-ddd-events-events-router
```