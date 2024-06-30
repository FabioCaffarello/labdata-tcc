package handler

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"libs/golang/ddd/events/input-broker/event"

	"github.com/stretchr/testify/suite"
)

// InputCreatedEventHandlerSuite is the test suite for InputCreatedEventHandler.
type InputCreatedEventHandlerSuite struct {
	suite.Suite
	notifier     *MockRabbitMQNotifier
	eventHandler *InputCreatedHandler
}

func TestInputCreatedEventHandlerSuite(t *testing.T) {
	suite.Run(t, new(InputCreatedEventHandlerSuite))
}

func (suite *InputCreatedEventHandlerSuite) SetupTest() {
	suite.notifier = new(MockRabbitMQNotifier)
	suite.eventHandler = NewInputCreatedHandler(suite.notifier)
}

// TestHandle tests the Handle method of InputCreatedHandler.
func (suite *InputCreatedEventHandlerSuite) TestHandle() {
	// Arrange
	testEvent := event.NewInputCreated()
	payload := map[string]string{"key": "value"}
	testEvent.SetPayload(payload)
	var wg sync.WaitGroup
	exchangeName := "test-exchange"
	routingKey := "test-routing-key"

	// Expected JSON output
	jsonOutput, _ := json.Marshal(payload)
	suite.notifier.On("Notify", jsonOutput, routingKey).Return(nil)

	// Act
	wg.Add(1)
	suite.eventHandler.Handle(testEvent, &wg, exchangeName, routingKey)
	wg.Wait()

	// Assert
	suite.notifier.AssertExpectations(suite.T())
}

// TestHandleNotifyError tests the Handle method when Notify returns an error.
func (suite *InputCreatedEventHandlerSuite) TestHandleNotifyError() {
	// Arrange
	testEvent := event.NewInputCreated()
	payload := map[string]string{"key": "value"}
	testEvent.SetPayload(payload)
	var wg sync.WaitGroup
	exchangeName := "test-exchange"
	routingKey := "test-routing-key"

	// Expected JSON output
	jsonOutput, _ := json.Marshal(payload)
	suite.notifier.On("Notify", jsonOutput, routingKey).Return(fmt.Errorf("error"))

	// Act
	wg.Add(1)
	suite.eventHandler.Handle(testEvent, &wg, exchangeName, routingKey)
	wg.Wait()

	// Assert
	suite.notifier.AssertExpectations(suite.T())
}
