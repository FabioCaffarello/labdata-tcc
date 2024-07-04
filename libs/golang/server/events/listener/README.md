# listener

`listener` is a Go library that provides an event listener management system. This library includes functionalities for creating and managing event listeners, adding and removing listeners, and starting listeners to process messages.

## Features

- Create and configure event listeners.
- Add and remove listeners.
- Start listeners to consume and process messages.
- Thread-safe management of listeners.

## Usage

### Creating and Configuring the Event Listener

The `NewEventListener` function creates a new `EventListener` instance.

```go
package main

import (
	"log"
	"libs/golang/server/events/listener/listener"
	usecaseprotocol "libs/golang/server/events/usecase-impl/protocol"
)

func main() {
	eventListener := listener.NewEventListener()

	// Add a listener (pseudo-code, implement your own ConsumerInterface and UseCaseProtocol)
	consumer := &YourConsumer{}
	usecaseProtocol := &YourUseCaseProtocol{}
	err := eventListener.AddListener(consumer, usecaseProtocol)
	if err != nil {
		log.Fatalf("Failed to add listener: %v", err)
	}

	// Start the listener
	err = eventListener.StartListener(consumer.GetListenerTag())
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
}
```

### Adding a Listener

The `AddListener` method adds a new listener to the `EventListener`.

```go
func main() {
	eventListener := listener.NewEventListener()

	// Add a listener (pseudo-code, implement your own ConsumerInterface and UseCaseProtocol)
	consumer := &YourConsumer{}
	usecaseProtocol := &YourUseCaseProtocol{}
	err := eventListener.AddListener(consumer, usecaseProtocol)
	if err != nil {
		log.Fatalf("Failed to add listener: %v", err)
	}
}
```

### Removing a Listener

The `RemoveListener` method removes a listener from the `EventListener`.

```go
func main() {
	eventListener := listener.NewEventListener()

	// Remove a listener by tag
	err := eventListener.RemoveListener("your-listener-tag")
	if err != nil {
		log.Fatalf("Failed to remove listener: %v", err)
	}
}
```

### Starting a Listener

The `StartListener` method starts a listener by its tag.

```go
func main() {
	eventListener := listener.NewEventListener()

	// Start a listener by tag
	err := eventListener.StartListener("your-listener-tag")
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
}
```

## Testing

To run the tests for the `listener` package, use the following command:

```sh
npx nx test libs-golang-server-events-listener
```
