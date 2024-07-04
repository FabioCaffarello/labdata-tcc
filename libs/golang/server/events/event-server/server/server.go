package server

import (
	eventListener "libs/golang/server/events/listener/listener"
	"log"
)

// ListenerServer represents a server that manages event listeners.
type ListenerServer struct {
	controller *eventListener.EventListener // Controller for managing event listeners.
	quitCh     chan struct{}                // Channel to signal the server to stop.
}

// NewListenerServer creates a new instance of ListenerServer.
//
// Parameters:
//   - controller: The event listener controller.
//
// Returns:
//   - A new instance of ListenerServer.
func NewListenerServer(
	controller *eventListener.EventListener,
) *ListenerServer {
	return &ListenerServer{
		controller: controller,
		quitCh:     make(chan struct{}),
	}
}

// Start begins the execution of the listener server.
//
// This method starts all the listeners managed by the controller in separate goroutines.
// It runs a main loop that waits for a signal on the quitCh channel to shut down the server.
func (es *ListenerServer) Start() {
	for listenerTag := range es.controller.GetListeners() {
		go es.controller.StartListener(listenerTag)
	}
mainloop:
	for {
		select {
		case <-es.quitCh:
			log.Println("Shutting down listener server")
			break mainloop
		}
	}
}
