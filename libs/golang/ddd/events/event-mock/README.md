# event-mock

The `event-mock` library provides mock implementations of the `EventInterface` and `EventDispatcherInterface` from the `go-events` package. These mocks are useful for testing purposes, allowing you to simulate and verify the behavior of event handling in your applications.

## Usage

### MockEvent

`MockEvent` is a mock implementation of the `EventInterface`. It can be used to simulate events in your tests.

#### Methods

- `GetName() string`: Returns the name of the mock event.
- `GetDateTime() time.Time`: Returns the date and time of the mock event.
- `GetPayload() interface{}`: Returns the payload of the mock event.
- `SetPayload(payload interface{})`: Sets the payload of the mock event.

### MockEventDispatcher

`MockEventDispatcher` is a mock implementation of the `EventDispatcherInterface`. It can be used to simulate event dispatching in your tests.

#### Methods

- `Register(eventName string, handler events.EventHandlerInterface) error`: Registers an event handler for a specific event name.
- `Dispatch(event events.EventInterface, exchangeName string, routingKey string) error`: Dispatches an event.
- `Remove(eventName string, handler events.EventHandlerInterface) error`: Removes an event handler for a specific event name.
- `Has(eventName string, handler events.EventHandlerInterface) bool`: Checks if a specific event handler is registered for an event name.
- `Clear()`: Clears all event handlers.

## Example

Here is an example of how to use the `event-mock` library in your tests:

```go
package yourpackage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"libs/golang/ddd/events/event-mock/mock"
	events "libs/golang/shared/go-events/amqp_events"
)

type YourTestSuite struct {
	suite.Suite
	eventMock        *eventmock.MockEvent
	dispatcherMock   *eventmock.MockEventDispatcher
}

func (suite *YourTestSuite) SetupTest() {
	suite.eventMock = new(eventmock.MockEvent)
	suite.dispatcherMock = new(eventmock.MockEventDispatcher)
}

func (suite *YourTestSuite) TestYourFunction() {
	// Arrange
	suite.eventMock.On("GetName").Return("TestEvent")
	suite.eventMock.On("GetDateTime").Return(time.Now())
	suite.eventMock.On("GetPayload").Return("payload")
	suite.eventMock.On("SetPayload", mock.Anything).Return()

	suite.dispatcherMock.On("Dispatch", suite.eventMock, "exchange", "routingKey").Return(nil)

	// Act
	// Call your function that uses the event and dispatcher mocks

	// Assert
	suite.eventMock.AssertExpectations(suite.T())
	suite.dispatcherMock.AssertExpectations(suite.T())
}

func TestYourTestSuite(t *testing.T) {
	suite.Run(t, new(YourTestSuite))
}
```