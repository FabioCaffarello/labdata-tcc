package repository

import (
	gomongodb "libs/golang/clients/resources/go-mongo/client"
	"libs/golang/ddd/domain/entities/output-vault/entity"
	mongowrapper "libs/golang/wrappers/resources/mongo-wrapper/wrapper"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OutputVaultMongoDBRepositorySuite struct {
	suite.Suite
	client      *mongo.Client
	wrapper     *mongowrapper.MongoDBWrapper
	output      *entity.Output
	outputProps entity.OutputProps
}

var (
	databaseName = "testdb"
)

func (suite *OutputVaultMongoDBRepositorySuite) SetupSuite() {
	os.Setenv("MONGODB_USER", "testuser")
	os.Setenv("MONGODB_PASSWORD", "testpassword")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27020")
	os.Setenv("MONGODB_DBNAME", databaseName)

	suite.wrapper = mongowrapper.NewMongoDBWrapper()
	err := suite.wrapper.Init()
	assert.NoError(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) SetupTest() {
	clientWrapper, ok := suite.wrapper.GetClient().(*gomongodb.Client)
	assert.True(suite.T(), ok, "expected *gomongodb.Client, got %T", clientWrapper)
	suite.client = clientWrapper.Client

	suite.outputProps = entity.OutputProps{
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Metadata: map[string]interface{}{
			"input_id": "input_id",
			"input": map[string]interface{}{
				"data": map[string]interface{}{
					"input1": "value1",
				},
				"processing_id":        "processing_id",
				"processing_timestamp": "2023-06-01 00:00:00",
			},
		},
	}

	var err error
	suite.output, err = entity.NewOutput(suite.outputProps)
	assert.Nil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TearDownTest() {
	err := suite.client.Database(databaseName).Drop(nil)
	suite.NoError(err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TearDownSuite() {
	os.Unsetenv("MONGODB_USER")
	os.Unsetenv("MONGODB_PASSWORD")
	os.Unsetenv("MONGODB_HOST")
	os.Unsetenv("MONGODB_PORT")
	os.Unsetenv("MONGODB_DBNAME")
}

func TestOutputVaultMongoDBRepositorySuite(t *testing.T) {
	suite.Run(t, new(OutputVaultMongoDBRepositorySuite))
}

func (suite *OutputVaultMongoDBRepositorySuite) TestNewOutputRepository() {
	repository := NewOutputRepository(suite.client, databaseName)
	assert.NotNil(suite.T(), repository)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestCreateOutput() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestCreateOutputAlreadyExists() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	err = repository.Create(suite.output)
	assert.NotNil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestGetOneByID() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	output, err := repository.getOneByID(suite.output.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
	assert.Equal(suite.T(), suite.output.Service, output.Service)
	assert.Equal(suite.T(), suite.output.Source, output.Source)
	assert.Equal(suite.T(), suite.output.Provider, output.Provider)
	assert.Equal(suite.T(), suite.output.Data, output.Data)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestGetOneByIDNotFound() {
	repository := NewOutputRepository(suite.client, databaseName)
	output, err := repository.getOneByID(suite.output.GetEntityID())
	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestGetOneByIDInvalidID() {
	repository := NewOutputRepository(suite.client, databaseName)
	output, err := repository.getOneByID("invalid_id")
	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFindByID() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	output, err := repository.FindByID(suite.output.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
	assert.Equal(suite.T(), suite.output.Service, output.Service)
	assert.Equal(suite.T(), suite.output.Source, output.Source)
	assert.Equal(suite.T(), suite.output.Provider, output.Provider)
	assert.Equal(suite.T(), suite.output.Data, output.Data)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFindAll() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	secDoc := suite.outputProps
	secDoc.Service = "test_service2"
	secOutput, err := entity.NewOutput(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(secOutput)
	assert.Nil(suite.T(), err)

	outputs, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), outputs)
	assert.Len(suite.T(), outputs, 2)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFindAllEmpty() {
	repository := NewOutputRepository(suite.client, databaseName)
	outputs, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), outputs)
	assert.Equal(suite.T(), 0, len(outputs))
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFindAllError() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	_, err = repository.FindAll()
	assert.Nil(suite.T(), err)

	err = suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)

	_, err = repository.FindAll()
	assert.Nil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestUpdate() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	outputStored, err := repository.FindByID(suite.output.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), outputStored)

	suite.outputProps.Metadata = map[string]interface{}{
		"input_id": "input_id",
		"input": map[string]interface{}{
			"data": map[string]interface{}{
				"input1": "value1",
			},
			"processing_id":        "processing_id_updated",
			"processing_timestamp": "2023-06-01 00:00:00",
		},
	}
	output, err := entity.NewOutput(suite.outputProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(output)
	assert.Nil(suite.T(), err)

	outputUpdated, err := repository.FindByID(output.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), outputUpdated)
	assert.Equal(suite.T(), "processing_id_updated", outputUpdated.Metadata.Input.ProcessingID)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestUpdateNotFound() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	suite.outputProps.Metadata = map[string]interface{}{
		"input_id": "input_id",
		"input": map[string]interface{}{
			"data": map[string]interface{}{
				"input1": "value1",
			},
			"processing_id":        "processing_id_updated",
			"processing_timestamp": "2023-06-01 00:00:00",
		},
	}
	suite.outputProps.Service = "test-service-updated"
	output, err := entity.NewOutput(suite.outputProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(output)
	assert.NotNil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestUpdateError() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	outputStored, err := repository.FindByID(suite.output.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), outputStored)

	suite.outputProps.Metadata = map[string]interface{}{
		"input_id": "input_id",
		"input": map[string]interface{}{
			"data": map[string]interface{}{
				"input1": "value1",
			},
			"processing_id":        "processing_id_updated",
			"processing_timestamp": "2023-06-01 00:00:00",
		},
	}
	output, err := entity.NewOutput(suite.outputProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(output)
	assert.Nil(suite.T(), err)

	err = suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)

	err = repository.Update(output)
	assert.NotNil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestDelete() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	err = repository.Delete(suite.output.GetEntityID())
	assert.Nil(suite.T(), err)

	output, err := repository.FindByID(suite.output.GetEntityID())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), output)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestDeleteNotFound() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	err = repository.Delete(suite.output.GetEntityID())
	assert.Nil(suite.T(), err)

	output, err := repository.FindByID(suite.output.GetEntityID())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), output)

	err = repository.Delete(suite.output.GetEntityID())
	assert.NotNil(suite.T(), err)
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFind() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	query := bson.M{"service": suite.output.Service}

	outputs, err := repository.find(query)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), outputs)
	assert.Equal(suite.T(), 1, len(outputs))
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFindEmpty() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	query := bson.M{"service": "invalid-service"}

	outputs, err := repository.find(query)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), outputs)
	assert.Equal(suite.T(), 0, len(outputs))
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFindAllByServiceAndProvider() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	secDoc := suite.outputProps
	secDoc.Source = "test_source2"
	secOutput, err := entity.NewOutput(secDoc)
	assert.Nil(suite.T(), err)
	err = repository.Create(secOutput)
	assert.Nil(suite.T(), err)

	configs, err := repository.FindAllByServiceAndProvider(suite.output.Provider, suite.output.Service)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 2, len(configs))
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFindAllBySourceAndProvider() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	secDoc := suite.outputProps
	secDoc.Source = "test_source2"
	secOutput, err := entity.NewOutput(secDoc)
	assert.Nil(suite.T(), err)
	err = repository.Create(secOutput)
	assert.Nil(suite.T(), err)

	schemas, err := repository.FindAllBySourceAndProvider(suite.output.Provider, suite.output.Source)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemas)
	assert.Equal(suite.T(), 1, len(schemas))
}

func (suite *OutputVaultMongoDBRepositorySuite) TestFindAllByServiceAndSourceAndProvider() {
	repository := NewOutputRepository(suite.client, databaseName)
	err := repository.Create(suite.output)
	assert.Nil(suite.T(), err)

	secDoc := suite.outputProps
	secDoc.Source = "test_source2"
	secOutput, err := entity.NewOutput(secDoc)
	assert.Nil(suite.T(), err)
	err = repository.Create(secOutput)
	assert.Nil(suite.T(), err)

	schemas, err := repository.FindAllByServiceAndSourceAndProvider(suite.output.Provider, suite.output.Service, suite.output.Source)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemas)
	assert.Equal(suite.T(), 1, len(schemas))
}
