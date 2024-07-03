package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	events "libs/golang/shared/go-events/amqp_events"
)

// ErrorCreatedHandler handles events of type ErrorCreated.
type ErrorCreatedHandler struct {
	Notifier NotifierInterface
}

// NewErrorCreatedHandler creates a new ErrorCreatedHandler.
func NewErrorCreatedHandler(notifier NotifierInterface) *ErrorCreatedHandler {
	return &ErrorCreatedHandler{
		Notifier: notifier,
	}
}

// Handle processes the event and sends a notification.
func (si *ErrorCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, routingKey string) {
	defer wg.Done()
	jsonOutput, _ := json.Marshal(event.GetPayload())
	err := si.Notifier.Notify(jsonOutput, routingKey)
	if err != nil {
		fmt.Println(err)
	}
}
