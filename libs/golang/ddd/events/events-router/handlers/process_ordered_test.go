package handler

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"libs/golang/ddd/events/events-router/event"

	"github.com/stretchr/testify/suite"
)

// OrderedProcessEventHandlerSuite is the test suite for OrderedProcessEventHandler.
type OrderedProcessEventHandlerSuite struct {
	suite.Suite
	notifier     *MockRabbitMQNotifier
	eventHandler *OrderedProcessHandler
}

func TestOrderedProcessEventHandlerSuite(t *testing.T) {
	suite.Run(t, new(OrderedProcessEventHandlerSuite))
}

func (suite *OrderedProcessEventHandlerSuite) SetupTest() {
	suite.notifier = new(MockRabbitMQNotifier)
	suite.eventHandler = NewOrderedProcessHandler(suite.notifier)
}

// TestHandle tests the Handle method of OrderedProcessHandler.
func (suite *OrderedProcessEventHandlerSuite) TestHandle() {
	// Arrange
	testEvent := event.NewOrderedProcess()
	payload := map[string]string{"key": "value"}
	testEvent.SetPayload(payload)
	var wg sync.WaitGroup
	routingKey := "test-routing-key"

	// Expected JSON output
	jsonOutput, _ := json.Marshal(payload)
	suite.notifier.On("Notify", jsonOutput, routingKey).Return(nil)

	// Act
	wg.Add(1)
	suite.eventHandler.Handle(testEvent, &wg, routingKey)
	wg.Wait()

	// Assert
	suite.notifier.AssertExpectations(suite.T())
}

// TestHandleNotifyError tests the Handle method when Notify returns an error.
func (suite *OrderedProcessEventHandlerSuite) TestHandleNotifyError() {
	// Arrange
	testEvent := event.NewOrderedProcess()
	payload := map[string]string{"key": "value"}
	testEvent.SetPayload(payload)
	var wg sync.WaitGroup
	routingKey := "test-routing-key"

	// Expected JSON output
	jsonOutput, _ := json.Marshal(payload)
	suite.notifier.On("Notify", jsonOutput, routingKey).Return(fmt.Errorf("error"))

	// Act
	wg.Add(1)
	suite.eventHandler.Handle(testEvent, &wg, routingKey)
	wg.Wait()

	// Assert
	suite.notifier.AssertExpectations(suite.T())
}
