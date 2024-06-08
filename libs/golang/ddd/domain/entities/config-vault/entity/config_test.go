package entity

import (
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
		DependsOn: []map[string]interface{}{
			{"service": "dep_service1", "source": "dep_source1"},
			{"service": "dep_service2", "source": "dep_source2"},
		},
		UpdatedAt: "2023-01-01 12:00:00",
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

	parsedDate, _ := time.Parse(dateLayout, configProps.UpdatedAt)
	assert.Equal(suite.T(), parsedDate, config.UpdatedAt)
	assert.NotZero(suite.T(), config.CreatedAt)
	assert.NotZero(suite.T(), config.ID)
}

func (suite *ConfigVaultConfigSuite) TestInvalidConfig() {
	configProps := ConfigProps{
		Active:    true,
		Service:   "",
		Source:    "test_source",
		Provider:  "test_provider",
		UpdatedAt: "2023-01-01 12:00:00",
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
		DependsOn: []map[string]interface{}{
			{"service": "dep_service1", "source": "dep_source1"},
		},
		UpdatedAt: "2023-01-01 12:00:00",
	}

	// fmt.Printf("Config Props: %+v\n", configProps)

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

	assert.IsType(suite.T(), []interface{}{}, doc["depends_on"])
	assert.Equal(suite.T(), 1, len(doc["depends_on"].([]interface{})))
}

func (suite *ConfigVaultConfigSuite) TestIsValid() {
	testdUUID, _ := uuid.GenerateUUIDFromMap(map[string]interface{}{"service": "test_service"})
	config := &Config{
		ID:              md5id.NewID(getIDData("test_service", "test_source", "test_provider")),
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		ConfigVersionID: testdUUID,
		CreatedAt:       time.Now(),
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
