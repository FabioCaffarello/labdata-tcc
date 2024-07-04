package repository

import (
	"fmt"
	"libs/golang/clients/resources/go-docdb/client"
	"libs/golang/database/go-docdb/database"
	"libs/golang/ddd/domain/entities/events-router/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// EventOrderRepositorySuite is a test suite for the EventOrderRepository.
type EventOrderRepositorySuite struct {
	suite.Suite
	repo   *EventOrderRepository
	client *client.Client
}

func TestEventOrderRepositorySuite(t *testing.T) {
	suite.Run(t, new(EventOrderRepositorySuite))
}

func (suite *EventOrderRepositorySuite) SetupTest() {
	// Initialize the in-memory database client
	db := database.NewInMemoryDocBD("test_database")
	suite.client = client.NewClient(db)
	suite.repo = NewEventOrderRepository(suite.client, "test_database")
}

func (suite *EventOrderRepositorySuite) TearDownTest() {
	suite.client = nil
	suite.repo = nil
}

func (suite *EventOrderRepositorySuite) TestNewEventOrderRepository() {
	assert.NotNil(suite.T(), suite.repo)
	assert.Equal(suite.T(), "events-order", suite.repo.collectionName)
}

func (suite *EventOrderRepositorySuite) TestGetOneByIDSuccess() {
	// Insert a document to retrieve
	doc := map[string]interface{}{
		"_id":           "9b97f68f63f3faa91d2d6558428f1863",
		"service":       "test_service",
		"source":        "test_source",
		"provider":      "test_provider",
		"stage":         "test_stage",
		"processing_id": "xyz789",
		"data": map[string]interface{}{
			"key": "value",
		},
	}
	err := suite.repo.client.InsertOne(suite.repo.collectionName, doc)
	assert.Nil(suite.T(), err)

	// Retrieve the document by ID
	result, err := suite.repo.getOneByID("9b97f68f63f3faa91d2d6558428f1863")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "test_service", result.Service)
	assert.Equal(suite.T(), "test_source", result.Source)
	assert.Equal(suite.T(), "test_provider", result.Provider)
	assert.Equal(suite.T(), "xyz789", string(result.ProcessingID))
	assert.Equal(suite.T(), map[string]interface{}{"key": "value"}, result.Data)
}

func (suite *EventOrderRepositorySuite) TestGetOneByIDError() {
	// Try to retrieve a non-existing document
	result, err := suite.repo.getOneByID("invalid_id")
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), "document not found", err.Error())
}

func (suite *EventOrderRepositorySuite) TestFindByID() {
	// Insert a document to retrieve
	doc := map[string]interface{}{
		"_id":           "9b97f68f63f3faa91d2d6558428f1863",
		"service":       "test_service",
		"source":        "test_source",
		"provider":      "test_provider",
		"stage":         "test_stage",
		"processing_id": "xyz789",
		"data": map[string]interface{}{
			"key": "value",
		},
	}
	err := suite.repo.client.InsertOne(suite.repo.collectionName, doc)
	assert.Nil(suite.T(), err)

	// Retrieve the document by ID
	result, err := suite.repo.FindByID("9b97f68f63f3faa91d2d6558428f1863")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "test_service", result.Service)
	assert.Equal(suite.T(), "test_source", result.Source)
	assert.Equal(suite.T(), "test_provider", result.Provider)
	assert.Equal(suite.T(), "xyz789", string(result.ProcessingID))
	assert.Equal(suite.T(), map[string]interface{}{"key": "value"}, result.Data)
}

func (suite *EventOrderRepositorySuite) TestCreateSuccess() {
	// Create an event order
	props := entity.EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "test_provider",
		Stage:        "test_stage",
		ProcessingID: "xyz789",
		Data: map[string]interface{}{
			"key": "value",
		},
	}
	eventOrder, err := entity.NewEventOrder(props)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), eventOrder)

	// Save the event order
	err = suite.repo.Create(eventOrder)
	assert.Nil(suite.T(), err)

	// Retrieve the saved event order by ID
	result, err := suite.repo.FindByID(eventOrder.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "test_service", result.Service)
	assert.Equal(suite.T(), "test_source", result.Source)
	assert.Equal(suite.T(), "test_provider", result.Provider)
	assert.Equal(suite.T(), "xyz789", string(result.ProcessingID))
	assert.Equal(suite.T(), map[string]interface{}{"key": "value"}, result.Data)
}

func (suite *EventOrderRepositorySuite) TestCreateDuplicate() {
	// Create an event order
	props := entity.EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "test_provider",
		Stage:        "test_stage",
		ProcessingID: "xyz789",
		Data: map[string]interface{}{
			"key": "value",
		},
	}
	eventOrder, err := entity.NewEventOrder(props)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), eventOrder)

	// Save the event order
	err = suite.repo.Create(eventOrder)
	assert.Nil(suite.T(), err)

	// Try to save the same event order again
	err = suite.repo.Create(eventOrder)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), fmt.Sprintf("Event order with ID %s already exists", eventOrder.GetEntityID()), err.Error())
}

func (suite *EventOrderRepositorySuite) TestCreateInvalid() {
	// Create an invalid event order with missing service
	props := entity.EventOrderProps{
		Source:       "test_source",
		Provider:     "test_provider",
		Stage:        "test_stage",
		ProcessingID: "xyz789",
		Data: map[string]interface{}{
			"key": "value",
		},
	}
	eventOrder, err := entity.NewEventOrder(props)
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), eventOrder)

	// Try to save the invalid event order
	err = suite.repo.Create(eventOrder)
	assert.NotNil(suite.T(), err)
}

func (suite *EventOrderRepositorySuite) TestCreateWithoutData() {
	// Create an event order with no data
	props := entity.EventOrderProps{
		Service:      "test_service",
		Source:       "test_source",
		Provider:     "test_provider",
		Stage:        "test_stage",
		ProcessingID: "xyz789",
		Data:         nil,
	}
	eventOrder, err := entity.NewEventOrder(props)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), eventOrder)

	// Save the event order
	err = suite.repo.Create(eventOrder)
	assert.Nil(suite.T(), err)

	// Retrieve the saved event order by ID
	result, err := suite.repo.FindByID(eventOrder.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "test_service", result.Service)
	assert.Equal(suite.T(), "test_source", result.Source)
	assert.Equal(suite.T(), "test_provider", result.Provider)
	assert.Equal(suite.T(), "xyz789", string(result.ProcessingID))
	assert.Nil(suite.T(), result.Data)
}

func (suite *EventOrderRepositorySuite) TestFindAll() {
	// Insert multiple documents to retrieve
	doc1 := map[string]interface{}{
		"_id":           "9b97f68f63f3faa91d2d6558428f1863",
		"service":       "test_service1",
		"source":        "test_source1",
		"provider":      "test_provider1",
		"stage":         "test_stage1",
		"processing_id": "xyz7891",
		"data": map[string]interface{}{
			"key1": "value1",
		},
	}
	doc2 := map[string]interface{}{
		"_id":           "9b97f68f63f3faa91d2d6558428f1864",
		"service":       "test_service2",
		"source":        "test_source2",
		"provider":      "test_provider2",
		"stage":         "test_stage2",
		"processing_id": "xyz7892",
		"data": map[string]interface{}{
			"key2": "value2",
		},
	}
	err := suite.repo.client.InsertOne(suite.repo.collectionName, doc1)
	assert.Nil(suite.T(), err)
	err = suite.repo.client.InsertOne(suite.repo.collectionName, doc2)
	assert.Nil(suite.T(), err)

	// Retrieve all documents
	results, err := suite.repo.FindAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), results)
	assert.Equal(suite.T(), 2, len(results))

	// Check first document
	assert.Equal(suite.T(), "test_service1", results[0].Service)
	assert.Equal(suite.T(), "test_source1", results[0].Source)
	assert.Equal(suite.T(), "test_provider1", results[0].Provider)
	assert.Equal(suite.T(), "xyz7891", string(results[0].ProcessingID))
	assert.Equal(suite.T(), map[string]interface{}{"key1": "value1"}, results[0].Data)

	// Check second document
	assert.Equal(suite.T(), "test_service2", results[1].Service)
	assert.Equal(suite.T(), "test_source2", results[1].Source)
	assert.Equal(suite.T(), "test_provider2", results[1].Provider)
	assert.Equal(suite.T(), "xyz7892", string(results[1].ProcessingID))
	assert.Equal(suite.T(), map[string]interface{}{"key2": "value2"}, results[1].Data)
}

func (suite *EventOrderRepositorySuite) TestDeleteSuccess() {
	// Insert a document to delete
	doc := map[string]interface{}{
		"_id":           "9b97f68f63f3faa91d2d6558428f1863",
		"service":       "test_service",
		"source":        "test_source",
		"provider":      "test_provider",
		"stage":         "test_stage",
		"processing_id": "xyz789",
		"data": map[string]interface{}{
			"key": "value",
		},
	}
	err := suite.repo.client.InsertOne(suite.repo.collectionName, doc)
	assert.Nil(suite.T(), err)

	// Delete the document by ID
	err = suite.repo.Delete("9b97f68f63f3faa91d2d6558428f1863")
	assert.Nil(suite.T(), err)

	// Try to retrieve the deleted document by ID
	result, err := suite.repo.FindByID("9b97f68f63f3faa91d2d6558428f1863")
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), "document not found", err.Error())
}

func (suite *EventOrderRepositorySuite) TestDeleteNonExisting() {
	// Try to delete a non-existing document by ID
	err := suite.repo.Delete("non_existing_id")
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "document not found", err.Error())
}
