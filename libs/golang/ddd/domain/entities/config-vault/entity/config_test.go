package entity

import (
	"log"
	"testing"
	"time"

	md5id "libs/golang/shared/id/go-md5"
	uuid "libs/golang/shared/id/go-uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigVaultConfigSuite struct {
	suite.Suite
}

func TestConfigVaultConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigVaultConfigSuite))
}

func (suite *ConfigVaultConfigSuite) TestNewConfig() {
	configProps := ConfigProps{
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

	config, err := NewConfig(configProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), config)
	assert.Equal(suite.T(), configProps.Service, config.Service)
	assert.Equal(suite.T(), configProps.Source, config.Source)
	assert.Equal(suite.T(), configProps.Provider, config.Provider)
	assert.Equal(suite.T(), configProps.Active, config.Active)
	assert.Equal(suite.T(), 2, len(config.DependsOn))
	assert.Equal(suite.T(), "dep_service1", config.DependsOn[0].Service)
	assert.Equal(suite.T(), "dep_source1", config.DependsOn[0].Source)
	assert.Equal(suite.T(), "dep_service2", config.DependsOn[1].Service)
	assert.Equal(suite.T(), "dep_source2", config.DependsOn[1].Source)
	assert.NotZero(suite.T(), config.CreatedAt)
	assert.NotZero(suite.T(), config.ID)
}

func (suite *ConfigVaultConfigSuite) TestInvalidConfig() {
	configProps := ConfigProps{
		Active:   true,
		Service:  "",
		Source:   "test_source",
		Provider: "test_provider",
		JobParameters: map[string]interface{}{
			"parser_module": "test_parser_module",
		},
	}

	_, err := NewConfig(configProps)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), ErrInvalidService, err)
}

func (suite *ConfigVaultConfigSuite) TestSetDependsOn() {
	config := &Config{}
	dependsOn := []JobDependencies{
		{Service: "service1", Source: "source1"},
		{Service: "service2", Source: "source2"},
	}
	config.SetDependsOn(dependsOn)
	assert.Equal(suite.T(), 2, len(config.DependsOn))
	assert.Equal(suite.T(), "service1", config.DependsOn[0].Service)
	assert.Equal(suite.T(), "source1", config.DependsOn[0].Source)
	assert.Equal(suite.T(), "service2", config.DependsOn[1].Service)
	assert.Equal(suite.T(), "source2", config.DependsOn[1].Source)
}

func (suite *ConfigVaultConfigSuite) TestToMap() {
	configProps := ConfigProps{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		JobParameters: map[string]interface{}{
			"parser_module": "test_parser_module",
		},
		DependsOn: []map[string]interface{}{
			{"service": "dep_service1", "source": "dep_source1"},
		},
	}

	config, err := NewConfig(configProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), config)

	doc, err := config.ToMap()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), doc)
	assert.IsType(suite.T(), map[string]interface{}{}, doc)

	assert.Equal(suite.T(), string(config.ID), doc["_id"])
	assert.Equal(suite.T(), string(config.ConfigVersionID), doc["config_version_id"])
	assert.Equal(suite.T(), config.Service, doc["service"])
	assert.Equal(suite.T(), config.Source, doc["source"])
	assert.Equal(suite.T(), config.Provider, doc["provider"])
	assert.Equal(suite.T(), config.Active, doc["active"])

	assert.IsType(suite.T(), []map[string]interface{}{}, doc["depends_on"])
	dependsOn := doc["depends_on"].([]map[string]interface{})
	assert.Equal(suite.T(), 1, len(dependsOn))
}

func (suite *ConfigVaultConfigSuite) TestIsValid() {
	testdUUID, _ := uuid.GenerateUUIDFromMap(map[string]interface{}{"service": "test_service"})
	config := &Config{
		ID:              md5id.NewID(getIDData("test_service", "test_source", "test_provider")),
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		ConfigVersionID: testdUUID,
		CreatedAt:       time.Now().Format(dateLayout),
	}

	err := config.isValid()
	assert.Nil(suite.T(), err)

	// Test with missing ID
	config.ID = ""
	err = config.isValid()
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), ErrInvalidID, err)

	// Test with missing Service
	config.ID = md5id.NewID(getIDData("test_service", "test_source", "test_provider"))
	config.Service = ""
	err = config.isValid()
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), ErrInvalidService, err)
}

func (suite *ConfigVaultConfigSuite) TestMapToEntity() {
	configProps := ConfigProps{
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

	// Create a new config entity
	config, err := NewConfig(configProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), config)

	// Convert config entity to map
	doc, err := config.ToMap()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), doc)

	// Convert map back to config entity
	newConfig, err := config.MapToEntity(doc)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), newConfig)

	// Assert the properties
	assert.Equal(suite.T(), config.ID, newConfig.ID)
	assert.Equal(suite.T(), config.Active, newConfig.Active)
	assert.Equal(suite.T(), config.Service, newConfig.Service)
	assert.Equal(suite.T(), config.Source, newConfig.Source)
	assert.Equal(suite.T(), config.Provider, newConfig.Provider)
	assert.Equal(suite.T(), config.ConfigVersionID, newConfig.ConfigVersionID)
	assert.Equal(suite.T(), config.CreatedAt, newConfig.CreatedAt)
	assert.Equal(suite.T(), config.UpdatedAt, newConfig.UpdatedAt)
	assert.Equal(suite.T(), config.DependsOn, newConfig.DependsOn)
	log.Printf("config.JobParameters: %v", config.JobParameters)
	assert.Equal(suite.T(), config.JobParameters, newConfig.JobParameters)
}
