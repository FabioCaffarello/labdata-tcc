# events-router/usecase

`events-router/usecase` is a Go library that provides the implementation of use cases for event-driven processing. This library includes functionalities for pre-processing input messages, handling errors, and dispatching processed orders.

## Features

- Pre-process input messages.
- Handle and dispatch error events.
- Dispatch processed orders to the appropriate channels.

## Usage

### Creating and Configuring the PreProcessingUseCase

The `NewPreProcessingUseCase` function creates a new `PreProcessingUseCase` instance with the specified event order repository, error event, process order event, and event dispatcher.

```go
package main

import (
	"log"
	"libs/golang/ddd/domain/entities/events-router/entity"
	"libs/golang/ddd/usecases/events-router/usecase"
	events "libs/golang/shared/go-events/amqp_events"
)

func main() {
	eventOrderRepository := entity.NewEventOrderRepository()
	errorCreated := events.NewEvent("ErrorCreated")
	processOrderCreated := events.NewEvent("ProcessOrderCreated")
	eventDispatcher := events.NewEventDispatcher()

	preProcessingUseCase := usecase.NewPreProcessingUseCase(
		eventOrderRepository,
		errorCreated,
		processOrderCreated,
		eventDispatcher,
	)

	// Simulate processing a message channel
	msgCh := make(chan []byte)
	go func() {
		msgCh <- []byte(`{"metadata": {"service": "exampleService", "source": "exampleSource", "provider": "exampleProvider", "processing_id": "12345"}, "data": {"key": "value"}}`)
		close(msgCh)
	}()

	preProcessingUseCase.ProcessMessageChannel(msgCh, "example-listener")
}
```

### Processing Messages

The `ProcessMessageChannel` method processes messages from the provided channel and dispatches them for further processing.

```go
func main() {
	eventOrderRepository := entity.NewEventOrderRepository()
	errorCreated := events.NewEvent("ErrorCreated")
	processOrderCreated := events.NewEvent("ProcessOrderCreated")
	eventDispatcher := events.NewEventDispatcher()

	preProcessingUseCase := usecase.NewPreProcessingUseCase(
		eventOrderRepository,
		errorCreated,
		processOrderCreated,
		eventDispatcher,
	)

	// Simulate processing a message channel
	msgCh := make(chan []byte)
	go func() {
		msgCh <- []byte(`{"metadata": {"service": "exampleService", "source": "exampleSource", "provider": "exampleProvider", "processing_id": "12345"}, "data": {"key": "value"}}`)
		close(msgCh)
	}()

	preProcessingUseCase.ProcessMessageChannel(msgCh, "example-listener")
}
```

## Testing

To run the tests for the `usecase` package, use the following command:

```sh
npx nx test libs-golang-ddd-usecases-events-router
```]