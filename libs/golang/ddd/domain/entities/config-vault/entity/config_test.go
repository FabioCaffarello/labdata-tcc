package entity

import (
	"testing"
	"time"

	md5id "libs/golang/shared/id/go-md5"
	uuid "libs/golang/shared/id/go-uuid"

	"github.com/stretchr/testify/suite"
)

type ConfigVaultConfigSuite struct {
	suite.Suite
}

func TestConfigVaultConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigVaultConfigSuite))
}

// Test cases

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
	suite.Nil(err)
	suite.NotNil(config)
	suite.Equal(configProps.Service, config.Service)
	suite.Equal(configProps.Source, config.Source)
	suite.Equal(configProps.Provider, config.Provider)
	suite.True(config.Active)
	suite.Equal(2, len(config.DependsOn))
	suite.Equal("dep_service1", config.DependsOn[0].Service)
	suite.Equal("dep_source1", config.DependsOn[0].Source)
	suite.Equal("dep_service2", config.DependsOn[1].Service)
	suite.Equal("dep_source2", config.DependsOn[1].Source)

	parsedDate, _ := time.Parse(dateLayout, configProps.UpdatedAt)
	suite.Equal(parsedDate, config.UpdatedAt)
	suite.NotZero(config.CreatedAt)
	suite.NotZero(config.ConfigVersionID)
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
	suite.NotNil(err)
	suite.Equal(ErrInvalidService, err)
}

func (suite *ConfigVaultConfigSuite) TestSetDependsOn() {
	config := &Config{}
	dependsOn := []JobDependencies{
		{Service: "service1", Source: "source1"},
		{Service: "service2", Source: "source2"},
	}
	config.SetDependsOn(dependsOn)
	suite.Equal(2, len(config.DependsOn))
	suite.Equal("service1", config.DependsOn[0].Service)
	suite.Equal("source1", config.DependsOn[0].Source)
	suite.Equal("service2", config.DependsOn[1].Service)
	suite.Equal("source2", config.DependsOn[1].Source)
}

// func (suite *ConfigVaultConfigSuite) TestToMap() {
// 	configProps := ConfigProps{
// 		Active:   true,
// 		Service:  "test_service",
// 		Source:   "test_source",
// 		Provider: "test_provider",
// 		DependsOn: []map[string]interface{}{
// 			{"service": "dep_service1", "source": "dep_source1"},
// 		},
// 		UpdatedAt: "2023-01-01 12:00:00",
// 	}

// 	config, _ := NewConfig(configProps)
// 	doc, err := config.ToMap()
// 	suite.Nil(err)
// 	suite.NotNil(doc)
// 	suite.Equal("test_service", doc["service"])
// 	suite.Equal("test_source", doc["source"])
// 	suite.Equal("test_provider", doc["provider"])
// 	suite.True(doc["active"].(bool))
// 	suite.Equal(1, len(doc["depends_on"].([]interface{})))
// }

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
	suite.Nil(err)

	// Test with missing ID
	config.ID = ""
	err = config.isValid()
	suite.NotNil(err)
	suite.Equal(ErrInvalidID, err)

	// Test with missing Service
	config.ID = md5id.NewID(getIDData("test_service", "test_source", "test_provider"))
	config.Service = ""
	err = config.isValid()
	suite.NotNil(err)
	suite.Equal(ErrInvalidService, err)
}
