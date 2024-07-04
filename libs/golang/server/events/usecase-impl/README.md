# usecase-impl

`usecase-impl` is a Go library that provides the interface definition for use case protocols. This library is intended to define a standard way for processing message channels within an event-driven architecture.

## Features

- Define a standard interface for use case protocols.
- Ensure consistency in processing message channels across different implementations.

## Usage

### Implementing the UseCaseProtocol Interface

The `UseCaseProtocol` interface defines a single method, `ProcessMessageChannel`, which processes messages from a given channel.

To use this library, you need to implement the `UseCaseProtocol` interface in your own struct.


### Example Implementation

Here's an example of how you might implement the `UseCaseProtocol` interface:

```go
package main

import (
	"fmt"
	usecaseprotocol "libs/golang/server/events/usecase-impl/protocol"
)

type MyUseCase struct{}

func (uc *MyUseCase) ProcessMessageChannel(msgCh <-chan []byte, listenerTag string) {
	for msg := range msgCh {
		fmt.Printf("Processing message from %s: %s\n", listenerTag, string(msg))
		// Add your message processing logic here
	}
}

func main() {
	myUseCase := &MyUseCase{}

	// Simulate a message channel
	msgCh := make(chan []byte)
	go func() {
		msgCh <- []byte("Hello, world!")
		close(msgCh)
	}()

	myUseCase.ProcessMessageChannel(msgCh, "example-listener")
}
```
