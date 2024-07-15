package repository

import (
	gomongodb "libs/golang/clients/resources/go-mongo/client"
	"libs/golang/ddd/domain/entities/config-vault/entity"
	mongowrapper "libs/golang/wrappers/resources/mongo-wrapper/wrapper"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConfigVaultMongoDBRepositorySuite struct {
	suite.Suite
	client      *mongo.Client
	wrapper     *mongowrapper.MongoDBWrapper
	config      *entity.Config
	configProps entity.ConfigProps
}

var (
	databaseName = "testdb"
)

func (suite *ConfigVaultMongoDBRepositorySuite) SetupSuite() {
	os.Setenv("MONGODB_USER", "testuser")
	os.Setenv("MONGODB_PASSWORD", "testpassword")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27019")
	os.Setenv("MONGODB_DBNAME", databaseName)

	suite.wrapper = mongowrapper.NewMongoDBWrapper()
	err := suite.wrapper.Init()
	assert.NoError(suite.T(), err)

}

func (suite *ConfigVaultMongoDBRepositorySuite) SetupTest() {
	clientWrapper, ok := suite.wrapper.GetClient().(*gomongodb.Client)
	assert.True(suite.T(), ok, "expected *gomongodb.Client, got %T", clientWrapper)
	suite.client = clientWrapper.Client

	suite.configProps = entity.ConfigProps{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		JobParameters: map[string]interface{}{
			"parser_module": "test_parser_module",
		},
		DependsOn: []map[string]interface{}{
			{"service": "dep_service1", "source": "dep_source1"},
			{"service": "dep_service2", "source": "dep_source2"},
		},
	}

	var err error
	suite.config, err = entity.NewConfig(suite.configProps)
	assert.Nil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TearDownTest() {
	err := suite.client.Database(databaseName).Drop(nil)
	suite.NoError(err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TearDownSuite() {
	os.Unsetenv("MONGODB_USER")
	os.Unsetenv("MONGODB_PASSWORD")
	os.Unsetenv("MONGODB_HOST")
	os.Unsetenv("MONGODB_PORT")
	os.Unsetenv("MONGODB_DBNAME")
}

func TestConfigVaultMongoDBRepositorySuite(t *testing.T) {
	suite.Run(t, new(ConfigVaultMongoDBRepositorySuite))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestNewConfigRepository() {
	repository := NewConfigRepository(suite.client, databaseName)
	assert.NotNil(suite.T(), repository)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestCreateConfig() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestCreateConfigAlreadyExists() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	err = repository.Create(suite.config)
	assert.NotNil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestGetOneByID() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	config, err := repository.getOneByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), config)
	assert.Equal(suite.T(), suite.config.Service, config.Service)
	assert.Equal(suite.T(), suite.config.Source, config.Source)
	assert.Equal(suite.T(), suite.config.Provider, config.Provider)
	assert.Equal(suite.T(), suite.config.Active, config.Active)
	assert.Equal(suite.T(), 2, len(config.DependsOn))
	assert.Equal(suite.T(), "dep_service1", config.DependsOn[0].Service)
	assert.Equal(suite.T(), "dep_source1", config.DependsOn[0].Source)
	assert.Equal(suite.T(), "dep_service2", config.DependsOn[1].Service)
	assert.Equal(suite.T(), "dep_source2", config.DependsOn[1].Source)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestGetOneByIDNotFound() {
	repository := NewConfigRepository(suite.client, databaseName)
	config, err := repository.getOneByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), config)
	assert.NotNil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestGetOneByIDInvalidID() {
	repository := NewConfigRepository(suite.client, databaseName)
	config, err := repository.getOneByID("invalid_id")
	assert.Nil(suite.T(), config)
	assert.NotNil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindByID() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	config, err := repository.FindByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), config)
	assert.Equal(suite.T(), suite.config.Service, config.Service)
	assert.Equal(suite.T(), suite.config.Source, config.Source)
	assert.Equal(suite.T(), suite.config.Provider, config.Provider)
	assert.Equal(suite.T(), suite.config.Active, config.Active)
	assert.Equal(suite.T(), 2, len(config.DependsOn))
	assert.Equal(suite.T(), "dep_service1", config.DependsOn[0].Service)
	assert.Equal(suite.T(), "dep_source1", config.DependsOn[0].Source)
	assert.Equal(suite.T(), "dep_service2", config.DependsOn[1].Service)
	assert.Equal(suite.T(), "dep_source2", config.DependsOn[1].Source)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindAll() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	secDoc := suite.configProps
	secDoc.Service = "test_service2"
	secConfig, err := entity.NewConfig(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(secConfig)
	assert.Nil(suite.T(), err)

	configs, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 2, len(configs))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindAllEmpty() {
	repository := NewConfigRepository(suite.client, databaseName)
	configs, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), configs)
	assert.Equal(suite.T(), 0, len(configs))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindAllError() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	_, err = repository.FindAll()
	assert.Nil(suite.T(), err)

	err = suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)

	_, err = repository.FindAll()
	assert.Nil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestUpdate() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	configStored, err := repository.FindByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configStored)
	assert.True(suite.T(), configStored.Active)

	suite.configProps.Active = false
	config, err := entity.NewConfig(suite.configProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(config)
	assert.Nil(suite.T(), err)

	configUpdated, err := repository.FindByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configUpdated)
	assert.False(suite.T(), configUpdated.Active)
	assert.Equal(suite.T(), configUpdated.CreatedAt, configStored.CreatedAt)
	assert.NotEqual(suite.T(), configUpdated.ConfigVersionID, configStored.ConfigVersionID)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestUpdateNotFound() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	suite.configProps.Active = false
	suite.configProps.Service = "test_service2"
	config, err := entity.NewConfig(suite.configProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(config)
	assert.NotNil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestUpdateError() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	configStored, err := repository.FindByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configStored)
	assert.True(suite.T(), configStored.Active)

	suite.configProps.Active = false
	config, err := entity.NewConfig(suite.configProps)
	assert.Nil(suite.T(), err)

	err = repository.Update(config)
	assert.Nil(suite.T(), err)

	configUpdated, err := repository.FindByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configUpdated)
	assert.False(suite.T(), configUpdated.Active)
	assert.Equal(suite.T(), configUpdated.CreatedAt, configStored.CreatedAt)
	assert.NotEqual(suite.T(), configUpdated.ConfigVersionID, configStored.ConfigVersionID)

	err = suite.client.Database(databaseName).Drop(nil)
	assert.Nil(suite.T(), err)

	err = repository.Update(config)
	assert.NotNil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestDelete() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	err = repository.Delete(suite.config.GetEntityID())
	assert.Nil(suite.T(), err)

	config, err := repository.FindByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), config)
	assert.NotNil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestDeleteNotFound() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Delete(suite.config.GetEntityID())
	assert.NotNil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestDeleteError() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	err = repository.Delete(suite.config.GetEntityID())
	assert.Nil(suite.T(), err)

	config, err := repository.FindByID(suite.config.GetEntityID())
	assert.Nil(suite.T(), config)
	assert.NotNil(suite.T(), err)

	err = repository.Delete(suite.config.GetEntityID())
	assert.NotNil(suite.T(), err)
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFind() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	query := bson.M{"service": suite.config.Service}

	configs, err := repository.find(query)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 1, len(configs))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindEmpty() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)
	query := bson.M{"source": "test_source2"}

	configs, err := repository.find(query)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 0, len(configs))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindAllByServiceAndProvider() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	secDoc := suite.configProps
	secDoc.Source = "test_source2"
	secConfig, err := entity.NewConfig(secDoc)
	assert.Nil(suite.T(), err)
	err = repository.Create(secConfig)
	assert.Nil(suite.T(), err)

	configs, err := repository.FindAllByServiceAndProvider(suite.config.Provider, suite.config.Service)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 2, len(configs))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindAllBySourceAndProvider() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	secDoc := suite.configProps
	secDoc.Service = "test_service2"
	secConfig, err := entity.NewConfig(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(secConfig)
	assert.Nil(suite.T(), err)

	configs, err := repository.FindAllBySourceAndProvider(suite.config.Provider, suite.config.Source)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 2, len(configs))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindAllByServiceAndSourceAndProvider() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	secDoc := suite.configProps
	secDoc.Service = "test_service2"
	secConfig, err := entity.NewConfig(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(secConfig)
	assert.Nil(suite.T(), err)

	configs, err := repository.FindAllByServiceAndSourceAndProvider(suite.config.Service, suite.config.Source, suite.config.Provider)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 1, len(configs))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindAllByServiceAndProviderAndActive() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	secDoc := suite.configProps
	secDoc.Active = false
	secDoc.Source = "test_source2"
	secConfig, err := entity.NewConfig(secDoc)
	assert.Nil(suite.T(), err)

	err = repository.Create(secConfig)
	assert.Nil(suite.T(), err)

	configs, err := repository.FindAllByServiceAndProviderAndActive(suite.config.Service, suite.config.Provider, suite.config.Active)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 1, len(configs))
}

func (suite *ConfigVaultMongoDBRepositorySuite) TestFindAllByProviderAndDependsOn() {
	repository := NewConfigRepository(suite.client, databaseName)
	err := repository.Create(suite.config)
	assert.Nil(suite.T(), err)

	configs, err := repository.FindAllByProviderAndDependsOn("test_provider", "dep_service1", "dep_source1")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), configs)
	assert.Equal(suite.T(), 1, len(configs))
}
