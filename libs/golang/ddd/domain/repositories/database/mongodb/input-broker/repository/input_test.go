package repository

import (
	gomongodb "libs/golang/clients/resources/go-mongo/client"
	"libs/golang/ddd/domain/entities/input-broker/entity"
	mongowrapper "libs/golang/wrappers/resources/mongo-wrapper/wrapper"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InputBrokerMongoDBRepositorySuite struct {
	suite.Suite
	client     *mongo.Client
	wrapper    *mongowrapper.MongoDBWrapper
	input      *entity.Input
	inputProps entity.InputProps
}

func TestInputBrokerMongoDBRepositorySuite(t *testing.T) {
	suite.Run(t, new(InputBrokerMongoDBRepositorySuite))
}

var (
	databaseName = "testdb"
)

func (suite *InputBrokerMongoDBRepositorySuite) SetupSuite() {
	os.Setenv("MONGODB_USER", "testuser")
	os.Setenv("MONGODB_PASSWORD", "testpassword")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27021")
	os.Setenv("MONGODB_DBNAME", databaseName)

	suite.wrapper = mongowrapper.NewMongoDBWrapper()
	err := suite.wrapper.Init()
	assert.NoError(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) SetupTest() {
	clientWrapper, ok := suite.wrapper.GetClient().(*gomongodb.Client)
	assert.True(suite.T(), ok, "expected *gomongodb.Client, got %T", clientWrapper)
	suite.client = clientWrapper.Client

	suite.inputProps = entity.InputProps{
		Service:  "test-service",
		Source:   "test-source",
		Provider: "test-provider",
		Data: map[string]interface{}{
			"test-key":  "test-value",
			"test-key2": "test-value2",
		},
	}

	var err error
	suite.input, err = entity.NewInput(suite.inputProps)
	assert.Nil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TearDownTest() {
	err := suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TearDownSuite() {
	os.Unsetenv("MONGODB_USER")
	os.Unsetenv("MONGODB_PASSWORD")
	os.Unsetenv("MONGODB_HOST")
	os.Unsetenv("MONGODB_PORT")
	os.Unsetenv("MONGODB_DBNAME")
}

func (suite *InputBrokerMongoDBRepositorySuite) TestNewInputRepository() {
	repository := NewInputRepository(suite.client, databaseName)
	assert.NotNil(suite.T(), repository)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestCreateInput() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestCreateInputAlreadyExists() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	err = repository.Create(suite.input)
	assert.NotNil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestGetOneByID() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	input, err := repository.getOneByID(suite.input.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), input)
	assert.Equal(suite.T(), suite.input.Metadata.Service, input.Metadata.Service)
	assert.Equal(suite.T(), suite.input.Metadata.Source, input.Metadata.Source)
	assert.Equal(suite.T(), suite.input.Metadata.Provider, input.Metadata.Provider)
	assert.Equal(suite.T(), suite.input.Data, input.Data)
	assert.Equal(suite.T(), suite.input.Status.Code, input.Status.Code)
	assert.Equal(suite.T(), suite.input.Status.Detail, input.Status.Detail)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestGetOneByIDNotFound() {
	repository := NewInputRepository(suite.client, databaseName)
	input, err := repository.getOneByID(suite.input.GetEntityID())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), input)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestGetOneByIDInvalidID() {
	repository := NewInputRepository(suite.client, databaseName)
	input, err := repository.getOneByID("invalid-id")
	assert.Nil(suite.T(), input)
	assert.NotNil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindByID() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	input, err := repository.FindByID(suite.input.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), input)
	assert.Equal(suite.T(), suite.input.Metadata.Service, input.Metadata.Service)
	assert.Equal(suite.T(), suite.input.Metadata.Source, input.Metadata.Source)
	assert.Equal(suite.T(), suite.input.Metadata.Provider, input.Metadata.Provider)
	assert.Equal(suite.T(), suite.input.Data, input.Data)
	assert.Equal(suite.T(), suite.input.Status.Code, input.Status.Code)
	assert.Equal(suite.T(), suite.input.Status.Detail, input.Status.Detail)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAll() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	secDoc := suite.inputProps
	secDoc.Service = "test-service2"
	secInput, err := entity.NewInput(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(secInput)
	assert.Nil(suite.T(), err)

	inputs, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Len(suite.T(), inputs, 2)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllEmpty() {
	repository := NewInputRepository(suite.client, databaseName)
	inputs, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), inputs)
	assert.Equal(suite.T(), 0, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllError() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	_, err = repository.FindAll()
	assert.Nil(suite.T(), err)

	err = suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)

	_, err = repository.FindAll()
	assert.Nil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestUpdate() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	inputStored, err := repository.FindByID(suite.input.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputStored)
	assert.Equal(suite.T(), suite.input.Metadata.ProcessingID, inputStored.Metadata.ProcessingID)

	inputStored.Metadata.ProcessingID = "new-processing-id"
	err = repository.Update(inputStored)
	assert.Nil(suite.T(), err)

	inputUpdated, err := repository.FindByID(suite.input.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputUpdated)
	assert.Equal(suite.T(), inputStored.Metadata.ProcessingID, inputUpdated.Metadata.ProcessingID)
	assert.Equal(suite.T(), inputStored.Metadata.Service, inputUpdated.Metadata.Service)
	assert.Equal(suite.T(), inputStored.Metadata.Source, inputUpdated.Metadata.Source)
	assert.Equal(suite.T(), inputStored.Metadata.Provider, inputUpdated.Metadata.Provider)
	assert.Equal(suite.T(), inputStored.Data, inputUpdated.Data)
	assert.Equal(suite.T(), inputStored.Status.Code, inputUpdated.Status.Code)
	assert.Equal(suite.T(), inputStored.Status.Detail, inputUpdated.Status.Detail)
	assert.Equal(suite.T(), inputStored.CreatedAt, inputUpdated.CreatedAt)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestUpdateNotFound() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	suite.inputProps.Service = "new-service"
	newInput, err := entity.NewInput(suite.inputProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(newInput)
	assert.NotNil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestUpdateError() {
	repository := NewInputRepository(suite.client, databaseName)

	// Create initial input
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	// Retrieve the stored input
	inputStored, err := repository.FindByID(suite.input.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputStored)

	// Backup the original stored input
	inputStoredBackup := *inputStored

	// Update the processing timestamp and update in repository
	inputStored.SetProcessingTimestamp(time.Now().Add(-time.Hour))
	err = repository.Update(inputStored)
	assert.Nil(suite.T(), err)

	// Retrieve the updated input
	inputUpdated, err := repository.FindByID(suite.input.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputUpdated)

	// Assert that the ProcessingID and Service are the same
	assert.Equal(suite.T(), inputStored.Metadata.ProcessingID, inputUpdated.Metadata.ProcessingID)
	assert.Equal(suite.T(), inputStored.Metadata.Service, inputUpdated.Metadata.Service)

	// Assert that the ProcessingTimestamp has changed
	assert.NotEqual(suite.T(), inputStoredBackup.Metadata.ProcessingTimestamp, inputUpdated.Metadata.ProcessingTimestamp)

	// Drop the database to simulate a non-existent database scenario
	err = suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)

	// Try to update the input in the dropped database and expect an error
	err = repository.Update(inputStored)
	assert.NotNil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestDelete() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	err = repository.Delete(suite.input.GetEntityID())
	assert.Nil(suite.T(), err)

	input, err := repository.FindByID(suite.input.GetEntityID())
	assert.Nil(suite.T(), input)
	assert.NotNil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestDeleteNotFound() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Delete("non-existent-id")
	assert.NotNil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestDeleteError() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	err = repository.Delete(suite.input.GetEntityID())
	assert.Nil(suite.T(), err)

	input, err := repository.FindByID(suite.input.GetEntityID())
	assert.Nil(suite.T(), input)
	assert.NotNil(suite.T(), err)

	err = repository.Delete(suite.input.GetEntityID())
	assert.NotNil(suite.T(), err)
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFind() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	query := bson.M{"metadata.service": suite.input.Metadata.Service}

	inputs, err := repository.find(query)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 1, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindEmpty() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)
	query := bson.M{"metadata.service": "test-service2"}

	inputs, err := repository.find(query)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 0, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllByServiceAndProvider() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	secDoc := suite.inputProps
	secDoc.Source = "test-source2"
	secInput, err := entity.NewInput(secDoc)
	assert.Nil(suite.T(), err)
	err = repository.Create(secInput)
	assert.Nil(suite.T(), err)

	inputs, err := repository.FindAllByServiceAndProvider(suite.input.Metadata.Provider, suite.input.Metadata.Service)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 2, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllBySourceAndProvider() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	secDoc := suite.inputProps
	secDoc.Service = "test-service2"
	secInput, err := entity.NewInput(secDoc)
	assert.Nil(suite.T(), err)
	err = repository.Create(secInput)
	assert.Nil(suite.T(), err)

	inputs, err := repository.FindAllBySourceAndProvider(suite.input.Metadata.Provider, suite.input.Metadata.Source)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 2, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllByServiceAndSourceAndProvider() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	secDoc := suite.inputProps
	secDoc.Service = "test-service2"
	secInput, err := entity.NewInput(secDoc)
	assert.Nil(suite.T(), err)
	err = repository.Create(secInput)
	assert.Nil(suite.T(), err)

	inputs, err := repository.FindAllByServiceAndSourceAndProvider(suite.input.Metadata.Provider, suite.input.Metadata.Service, suite.input.Metadata.Source)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 1, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllByStatus() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	secDoc := suite.inputProps
	secDoc.Service = "test-service2"
	secInput, err := entity.NewInput(secDoc)
	assert.Nil(suite.T(), err)
	secInput.SetStatus(1, "Processing")
	err = repository.Create(secInput)
	assert.Nil(suite.T(), err)

	inputs, err := repository.FindAllByStatus(1)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 1, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllByStatusAndServiceAndProvider() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	secDoc := suite.inputProps
	secDoc.Service = "test-service2"
	secInput, err := entity.NewInput(secDoc)
	assert.Nil(suite.T(), err)
	secInput.SetStatus(1, "Processing")
	err = repository.Create(secInput)
	assert.Nil(suite.T(), err)

	inputs, err := repository.FindAllByStatusAndServiceAndProvider(suite.input.Metadata.Provider, suite.input.Metadata.Service, 0)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 1, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllByStatusAndSourceAndProvider() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	secDoc := suite.inputProps
	secDoc.Source = "test-source2"
	secInput, err := entity.NewInput(secDoc)
	assert.Nil(suite.T(), err)
	secInput.SetStatus(1, "Processing")
	err = repository.Create(secInput)
	assert.Nil(suite.T(), err)

	inputs, err := repository.FindAllByStatusAndSourceAndProvider(suite.input.Metadata.Provider, suite.input.Metadata.Source, 0)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 1, len(inputs))
}

func (suite *InputBrokerMongoDBRepositorySuite) TestFindAllByStatusAndServiceAndSourceAndProvider() {
	repository := NewInputRepository(suite.client, databaseName)
	err := repository.Create(suite.input)
	assert.Nil(suite.T(), err)

	secDoc := suite.inputProps
	secDoc.Service = "test-service2"
	secDoc.Source = "test-source2"
	secInput, err := entity.NewInput(secDoc)
	assert.Nil(suite.T(), err)
	secInput.SetStatus(1, "Processing")
	err = repository.Create(secInput)
	assert.Nil(suite.T(), err)

	inputs, err := repository.FindAllByStatusAndServiceAndSourceAndProvider(suite.input.Metadata.Provider, suite.input.Metadata.Service, suite.input.Metadata.Source, 0)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), inputs)
	assert.Equal(suite.T(), 1, len(inputs))
}
