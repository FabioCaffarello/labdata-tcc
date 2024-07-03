package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	events "libs/golang/shared/go-events/amqp_events"
)

// InputCreatedHandler handles events of type InputCreated.
type InputCreatedHandler struct {
	Notifier NotifierInterface
}

// NewInputCreatedHandler creates a new InputCreatedHandler.
func NewInputCreatedHandler(notifier NotifierInterface) *InputCreatedHandler {
	return &InputCreatedHandler{
		Notifier: notifier,
	}
}

// Handle processes the event and sends a notification.
func (si *InputCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, routingKey string) {
	defer wg.Done()
	jsonOutput, _ := json.Marshal(event.GetPayload())
	err := si.Notifier.Notify(jsonOutput, routingKey)
	if err != nil {
		fmt.Println(err)
	}
}
