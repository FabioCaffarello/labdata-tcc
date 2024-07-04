# event-server

`event-server` is a Go library that provides an event listener server for managing and executing event listeners. This library includes functionalities for creating and managing a server, starting listeners, and handling server shutdown.

## Features

- Create and configure an event listener server.
- Start event listeners and manage their execution.
- Gracefully stop the server and listeners.

## Usage

### Creating and Configuring the Event Listener Server

The `NewListenerServer` function creates a new `ListenerServer` instance with the specified event listener controller.

```go
package main

import (
    "log"
    eventListener "libs/golang/server/events/listener/listener"
    "libs/golang/server/events/event-server/server"
)

func main() {
    listenerController := eventListener.NewEventListenerController()
    listenerServer := server.NewListenerServer(listenerController)

    go listenerServer.Start()

    // Simulate server running for a period of time
    // ...

    listenerServer.Stop()
}
```

### Starting the Listener Server

The `Start` method begins the execution of the listener server. It starts all the listeners managed by the controller in separate goroutines.

```go
func main() {
    listenerController := eventListener.NewEventListenerController()
    listenerServer := server.NewListenerServer(listenerController)

    go listenerServer.Start()

    // Simulate server running for a period of time
    // ...

    listenerServer.Stop()
}
```

### Stopping the Listener Server

The `Stop` method stops the server by sending a signal to the quit channel, which gracefully shuts down the listeners.

```go
func main() {
    listenerController := eventListener.NewEventListenerController()
    listenerServer := server.NewListenerServer(listenerController)

    go listenerServer.Start()

    // Simulate server running for a period of time
    // ...

    listenerServer.Stop()
}
```

## Testing

To run the tests for the `event-server` package, use the following command:

```sh
npx nx test libs-golang-server-events-event-server
```