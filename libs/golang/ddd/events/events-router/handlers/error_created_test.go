package handler

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"libs/golang/ddd/events/events-router/event"

	"github.com/stretchr/testify/suite"
)

// ErrorCreatedEventHandlerSuite is the test suite for ErrorCreatedEventHandler.
type ErrorCreatedEventHandlerSuite struct {
	suite.Suite
	notifier     *MockRabbitMQNotifier
	eventHandler *ErrorCreatedHandler
}

func TestErrorCreatedEventHandlerSuite(t *testing.T) {
	suite.Run(t, new(ErrorCreatedEventHandlerSuite))
}

func (suite *ErrorCreatedEventHandlerSuite) SetupTest() {
	suite.notifier = new(MockRabbitMQNotifier)
	suite.eventHandler = NewErrorCreatedHandler(suite.notifier)
}

// TestHandle tests the Handle method of ErrorCreatedHandler.
func (suite *ErrorCreatedEventHandlerSuite) TestHandle() {
	// Arrange
	testEvent := event.NewErrorCreated()
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
func (suite *ErrorCreatedEventHandlerSuite) TestHandleNotifyError() {
	// Arrange
	testEvent := event.NewErrorCreated()
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
