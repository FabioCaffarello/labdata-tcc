package entity

import (
	"testing"

	uuid "libs/golang/shared/id/go-uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventsRouterEventOrderSuite struct {
	suite.Suite
}

func TestEventsRouterEventOrderSuite(t *testing.T) {
	suite.Run(t, new(EventsRouterEventOrderSuite))
}

func (suite *EventsRouterEventOrderSuite) TestNewEventOrderWhenSuccess() {
	props := EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "test_provider",
		ProcessingID: "xyz789",
		Data:         map[string]interface{}{"key": "value"},
	}

	eventOrder, err := NewEventOrder(props)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "test_service", eventOrder.Service)
	assert.Equal(suite.T(), "test_source", eventOrder.Source)
	assert.Equal(suite.T(), "test_provider", eventOrder.Provider)
	assert.Equal(suite.T(), uuid.ID("xyz789"), eventOrder.ProcessingID)
	assert.Equal(suite.T(), map[string]interface{}{"key": "value"}, eventOrder.Data)
}

func (suite *EventsRouterEventOrderSuite) TestNewEventOrderWhenInvalidService() {
	props := EventOrderProps{
		Service:      "",
		Source:       "test_source",
		Provider:     "test_provider",
		ProcessingID: "xyz789",
		Data:         map[string]interface{}{"key": "value"},
	}

	eventOrder, err := NewEventOrder(props)

	assert.Nil(suite.T(), eventOrder)
	assert.Equal(suite.T(), ErrInvalidService, err)
}

func (suite *EventsRouterEventOrderSuite) TestNewEventOrderWhenInvalidSource() {
	props := EventOrderProps{
		Service:      "test_service",
		Source:       "",
		Provider:     "test_provider",
		ProcessingID: "xyz789",
		Data:         map[string]interface{}{"key": "value"},
	}

	eventOrder, err := NewEventOrder(props)

	assert.Nil(suite.T(), eventOrder)
	assert.Equal(suite.T(), ErrInvalidSource, err)
}

func (suite *EventsRouterEventOrderSuite) TestNewEventOrderWhenInvalidProvider() {
	props := EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "",
		ProcessingID: "xyz789",
		Data:         map[string]interface{}{"key": "value"},
	}

	eventOrder, err := NewEventOrder(props)

	assert.Nil(suite.T(), eventOrder)
	assert.Equal(suite.T(), ErrInvalidProvider, err)
}

func (suite *EventsRouterEventOrderSuite) TestNewEventOrderWhenInvalidProcessingID() {
	props := EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "test_provider",
		ProcessingID: "",
		Data:         map[string]interface{}{"key": "value"},
	}

	eventOrder, err := NewEventOrder(props)

	assert.Nil(suite.T(), eventOrder)
	assert.Equal(suite.T(), ErrInvalidProcessingID, err)
}

func (suite *EventsRouterEventOrderSuite) TestGetEntityID() {
	props := EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "test_provider",
		ProcessingID: "xyz789",
		Data:         map[string]interface{}{"key": "value"},
	}

	eventOrder, _ := NewEventOrder(props)
	entityID := eventOrder.GetEntityID()

	assert.NotEmpty(suite.T(), entityID)
}

func (suite *EventsRouterEventOrderSuite) TestToMap() {
	props := EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "test_provider",
		ProcessingID: "xyz789",
		Data:         map[string]interface{}{"key": "value"},
	}
	eventOrder, err := NewEventOrder(props)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), eventOrder)
	assert.NotEmpty(suite.T(), eventOrder.ID)

	doc, err := eventOrder.ToMap()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), doc)
	assert.IsType(suite.T(), map[string]interface{}{}, doc)

	assert.Equal(suite.T(), string(eventOrder.ID), doc["_id"])
	assert.Equal(suite.T(), eventOrder.Service, doc["service"])
	assert.Equal(suite.T(), eventOrder.Source, doc["source"])
	assert.Equal(suite.T(), eventOrder.Provider, doc["provider"])
	assert.Equal(suite.T(), string(eventOrder.ProcessingID), doc["processing_id"])
	assert.Equal(suite.T(), eventOrder.Data, doc["data"])
}

func (suite *EventsRouterEventOrderSuite) TestMapToEntity() {
	eventOrderProps := EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "test_provider",
		ProcessingID: "xyz789",
		Data:         map[string]interface{}{"key": "value"},
	}

	eventOrder, err := NewEventOrder(eventOrderProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), eventOrder)

	doc, err := eventOrder.ToMap()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), doc)
	assert.IsType(suite.T(), map[string]interface{}{}, doc)

	newEventOrder := &EventOrder{}
	newEventOrder, err = newEventOrder.MapToEntity(doc)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), newEventOrder)

	assert.Equal(suite.T(), eventOrder.ID, newEventOrder.ID)
	assert.Equal(suite.T(), eventOrder.Service, newEventOrder.Service)
	assert.Equal(suite.T(), eventOrder.Source, newEventOrder.Source)
	assert.Equal(suite.T(), eventOrder.Provider, newEventOrder.Provider)
	assert.Equal(suite.T(), eventOrder.ProcessingID, newEventOrder.ProcessingID)
	assert.Equal(suite.T(), eventOrder.Data, newEventOrder.Data)
}
