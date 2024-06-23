package repository

import (
	"os"
	"testing"

	gomongodb "libs/golang/clients/resources/go-mongo/client"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"libs/golang/ddd/domain/entities/schema-vault/entity"
	mongowrapper "libs/golang/wrappers/resources/mongo-wrapper/wrapper"
)

type SchemaRepositoryTestSuite struct {
	suite.Suite
	client      *mongo.Client
	wrapper     *mongowrapper.MongoDBWrapper
	schema      *entity.Schema
	schemaProps entity.SchemaProps
}

var (
	databaseName = "testdb"
)

func (suite *SchemaRepositoryTestSuite) SetupSuite() {
	os.Setenv("MONGODB_USER", "testuser")
	os.Setenv("MONGODB_PASSWORD", "testpassword")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DBNAME", databaseName)

	suite.wrapper = mongowrapper.NewMongoDBWrapper()
	err := suite.wrapper.Init()
	assert.NoError(suite.T(), err)
}

func (suite *SchemaRepositoryTestSuite) SetupTest() {
	clientWrapper, ok := suite.wrapper.GetClient().(*gomongodb.Client)
	assert.True(suite.T(), ok, "expected *gomongodb.Client, got %T", clientWrapper)
	suite.client = clientWrapper.Client

	suite.schemaProps = entity.SchemaProps{
		Service:    "test-service",
		Source:     "test-source",
		Provider:   "test-provider",
		SchemaType: "test-schema-type",
		JsonSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
				"field2": map[string]interface{}{
					"type": "string",
				},
			},
			"required": []interface{}{
				"field1",
			},
		},
	}

	var err error
	suite.schema, err = entity.NewSchema(suite.schemaProps)
	assert.Nil(suite.T(), err)
}

func (suite *SchemaRepositoryTestSuite) TearDownTest() {
	err := suite.client.Database(databaseName).Drop(nil)
	suite.NoError(err)
}

func (suite *SchemaRepositoryTestSuite) TearDownSuite() {
	os.Unsetenv("MONGODB_USER")
	os.Unsetenv("MONGODB_PASSWORD")
	os.Unsetenv("MONGODB_HOST")
	os.Unsetenv("MONGODB_PORT")
	os.Unsetenv("MONGODB_DBNAME")
}

func TestSchemaRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SchemaRepositoryTestSuite))
}

func (suite *SchemaRepositoryTestSuite) TestNewSchemaRepository() {
	repository := NewSchemaRepository(suite.client, databaseName)
	assert.NotNil(suite.T(), repository)
}

func (suite *SchemaRepositoryTestSuite) TestCreateSchema() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)
}

func (suite *SchemaRepositoryTestSuite) TestSchemaAlreadyExists() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	err = repository.Create(suite.schema)
	assert.NotNil(suite.T(), err)
}

func (suite *SchemaRepositoryTestSuite) TestGetOneByID() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	schema, err := repository.getOneByID(suite.schema.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.schema.Service, schema.Service)
	assert.Equal(suite.T(), suite.schema.Source, schema.Source)
	assert.Equal(suite.T(), suite.schema.Provider, schema.Provider)
	assert.Equal(suite.T(), suite.schema.SchemaType, schema.SchemaType)
	assert.Equal(suite.T(), suite.schema.JsonSchema, schema.JsonSchema)
}

func (suite *SchemaRepositoryTestSuite) TestGetOneByIDNotFound() {
	repository := NewSchemaRepository(suite.client, databaseName)
	schema, err := repository.getOneByID(suite.schema.GetEntityID())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), schema)
}

func (suite *SchemaRepositoryTestSuite) TestGetOneByIDInvalidID() {
	repository := NewSchemaRepository(suite.client, databaseName)
	schema, err := repository.getOneByID("invalid_id")
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), schema)
}

func (suite *SchemaRepositoryTestSuite) TestFindByID() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	schema, err := repository.FindByID(suite.schema.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.schema.Service, schema.Service)
	assert.Equal(suite.T(), suite.schema.Source, schema.Source)
	assert.Equal(suite.T(), suite.schema.Provider, schema.Provider)
	assert.Equal(suite.T(), suite.schema.SchemaType, schema.SchemaType)
	assert.Equal(suite.T(), suite.schema.JsonSchema, schema.JsonSchema)
}

func (suite *SchemaRepositoryTestSuite) TestFindAll() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	secDoc := suite.schemaProps
	secDoc.Service = "test_service2"
	seSchema, err := entity.NewSchema(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(seSchema)
	assert.Nil(suite.T(), err)

	schemas, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemas)
	assert.Equal(suite.T(), 2, len(schemas))
}

func (suite *SchemaRepositoryTestSuite) TestFindAllEmpty() {
	repository := NewSchemaRepository(suite.client, databaseName)
	schemas, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemas)
	assert.Equal(suite.T(), 0, len(schemas))
}

func (suite *SchemaRepositoryTestSuite) TestFindAllError() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	_, err = repository.FindAll()
	assert.Nil(suite.T(), err)

	err = suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)

	_, err = repository.FindAll()
	assert.Nil(suite.T(), err)
}

func (suite *SchemaRepositoryTestSuite) TestUpdate() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	schemaStored, err := repository.FindByID(suite.schema.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemaStored)

	suite.schemaProps.SchemaType = "test-schema-type-updated"
	schema, err := entity.NewSchema(suite.schemaProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(schema)
	assert.Nil(suite.T(), err)

	schemaUpdated, err := repository.FindByID(suite.schema.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemaUpdated)
	assert.Equal(suite.T(), schema.SchemaType, schemaUpdated.SchemaType)
}

func (suite *SchemaRepositoryTestSuite) TestUpdateNotFound() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	suite.schemaProps.SchemaType = "test-schema-type-updated"
	suite.schemaProps.Service = "test-service-updated"
	schema, err := entity.NewSchema(suite.schemaProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(schema)
	assert.NotNil(suite.T(), err)
}

func (suite *SchemaRepositoryTestSuite) TestUpdateError() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	schemaStored, err := repository.FindByID(suite.schema.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemaStored)

	suite.schemaProps.SchemaType = "test-schema-type-updated"
	schema, err := entity.NewSchema(suite.schemaProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(schema)
	assert.Nil(suite.T(), err)

	err = suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)

	err = repository.Update(schema)
	assert.NotNil(suite.T(), err)
}

func (suite *SchemaRepositoryTestSuite) TestDelete() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	err = repository.Delete(suite.schema.GetEntityID())
	assert.Nil(suite.T(), err)

	schema, err := repository.FindByID(suite.schema.GetEntityID())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), schema)
}

func (suite *SchemaRepositoryTestSuite) TestDeleteNotFound() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	err = repository.Delete(suite.schema.GetEntityID())
	assert.Nil(suite.T(), err)

	schema, err := repository.FindByID(suite.schema.GetEntityID())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), schema)

	err = repository.Delete(suite.schema.GetEntityID())
	assert.NotNil(suite.T(), err)
}

func (suite *SchemaRepositoryTestSuite) TestFind() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	query := bson.M{"service": suite.schema.Service}

	schemas, err := repository.find(query)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemas)
	assert.Equal(suite.T(), 1, len(schemas))
}

func (suite *SchemaRepositoryTestSuite) TestFindEmpty() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	query := bson.M{"service": "invalid-service"}

	schemas, err := repository.find(query)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemas)
	assert.Equal(suite.T(), 0, len(schemas))
}

func (suite *SchemaRepositoryTestSuite) TestFindAllByServiceAndProvider() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	secDoc := suite.schemaProps
	secDoc.Source = "test_source2"
	secSchema, err := entity.NewSchema(secDoc)
	assert.Nil(suite.T(), err)
	err = repository.Create(secSchema)
	assert.Nil(suite.T(), err)

	configs, err := repository.FindAllByServiceAndProvider(suite.schema.Provider, suite.schema.Service)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 2, len(configs))
}

func (suite *SchemaRepositoryTestSuite) TestFindAllBySourceAndProvider() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	secDoc := suite.schemaProps
	secDoc.Provider = "test_provider2"
	seSchema, err := entity.NewSchema(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(seSchema)
	assert.Nil(suite.T(), err)

	schemas, err := repository.FindAllBySourceAndProvider(suite.schema.Provider, suite.schema.Source)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemas)
	assert.Equal(suite.T(), 1, len(schemas))
}

func (suite *SchemaRepositoryTestSuite) TestFindAllByServiceAndSourceAndProvider() {
	repository := NewSchemaRepository(suite.client, databaseName)
	err := repository.Create(suite.schema)
	assert.Nil(suite.T(), err)

	secDoc := suite.schemaProps
	secDoc.Provider = "test_provider2"
	seSchema, err := entity.NewSchema(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(seSchema)
	assert.Nil(suite.T(), err)

	schemas, err := repository.FindAllByServiceAndSourceAndProvider(suite.schema.Service, suite.schema.Source, suite.schema.Provider)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), schemas)
	assert.Equal(suite.T(), 1, len(schemas))
}
