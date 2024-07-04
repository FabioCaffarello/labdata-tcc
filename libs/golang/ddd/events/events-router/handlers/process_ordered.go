package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	events "libs/golang/shared/go-events/amqp_events"
)

// OrderedProcessHandler handles events of type OrderedProcess.
type OrderedProcessHandler struct {
	Notifier NotifierInterface
}

// NewOrderedProcessHandler creates a new OrderedProcessHandler.
func NewOrderedProcessHandler(notifier NotifierInterface) *OrderedProcessHandler {
	return &OrderedProcessHandler{
		Notifier: notifier,
	}
}

// Handle processes the event and sends a notification.
func (si *OrderedProcessHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, routingKey string) {
	defer wg.Done()
	jsonOutput, _ := json.Marshal(event.GetPayload())
	err := si.Notifier.Notify(jsonOutput, routingKey)
	if err != nil {
		fmt.Println(err)
	}
}
