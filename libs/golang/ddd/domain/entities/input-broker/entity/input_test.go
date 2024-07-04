package entity

import (
	md5id "libs/golang/shared/id/go-md5"
	uuid "libs/golang/shared/id/go-uuid"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InputVaultConfigSuite struct {
	suite.Suite
}

func TestInputVaultConfigSuite(t *testing.T) {
	suite.Run(t, new(InputVaultConfigSuite))
}

func (suite *InputVaultConfigSuite) TestNewInput() {
	inputProps := InputProps{
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
	}

	input, err := NewInput(inputProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), input)
	assert.Equal(suite.T(), inputProps.Service, input.Metadata.Service)
	assert.Equal(suite.T(), inputProps.Source, input.Metadata.Source)
	assert.Equal(suite.T(), inputProps.Provider, input.Metadata.Provider)
	assert.Equal(suite.T(), inputProps.Data, input.Data)
	assert.NotEmpty(suite.T(), input.ID)
}

func (suite *InputVaultConfigSuite) TestInvalidInput() {
	inputProps := InputProps{
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Source:   "test_source",
		Provider: "test_provider",
	}

	_, err := NewInput(inputProps)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), ErrInvalidService, err)
}

func (suite *InputVaultConfigSuite) TestToMap() {
	inputProps := InputProps{
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
	}

	input, err := NewInput(inputProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), input)
	assert.NotEmpty(suite.T(), input.ID)

	doc, err := input.ToMap()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), doc)
	assert.IsType(suite.T(), map[string]interface{}{}, doc)

	assert.Equal(suite.T(), input.Metadata.Service, doc["metadata"].(map[string]interface{})["service"])
	assert.Equal(suite.T(), input.Metadata.Source, doc["metadata"].(map[string]interface{})["source"])
	assert.Equal(suite.T(), input.Metadata.Provider, doc["metadata"].(map[string]interface{})["provider"])
	assert.Equal(suite.T(), input.Data, doc["data"])
}

func (suite *InputVaultConfigSuite) TestIsValid() {
	idData := getIDData("test_service", "test_source", "test_provider", map[string]interface{}{"field1": "value1"})
	processingID, _ := uuid.GenerateUUIDFromMap(idData)
	input := &Input{
		ID: md5id.NewID(idData),
		Metadata: Metadata{
			Service:             "test_service",
			Source:              "test_source",
			Provider:            "test_provider",
			ProcessingID:        processingID,
			ProcessingTimestamp: "2023-06-01 00:00:00",
		},
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Status: Status{
			Code:   200,
			Detail: "OK",
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}

	err := input.isValid()
	assert.Nil(suite.T(), err)

	input.Metadata.Service = ""
	err = input.isValid()
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), ErrInvalidService, err)
}

func (suite *InputVaultConfigSuite) TestMapToEntity() {
	inputProps := InputProps{
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Service:  "test-service",
		Source:   "test-source",
		Provider: "test-provider",
	}

	input, err := NewInput(inputProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), input)

	doc, err := input.ToMap()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), doc)
	assert.IsType(suite.T(), map[string]interface{}{}, doc)

	newInput := &Input{}
	newInput, err = newInput.MapToEntity(doc)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), newInput)

	assert.Equal(suite.T(), input.Metadata.Service, newInput.Metadata.Service)
	assert.Equal(suite.T(), input.Metadata.Source, newInput.Metadata.Source)
	assert.Equal(suite.T(), input.Metadata.Provider, newInput.Metadata.Provider)
	assert.Equal(suite.T(), input.Data, newInput.Data)
	assert.Equal(suite.T(), input.ID, newInput.ID)
	assert.Equal(suite.T(), input.Status, newInput.Status)
	assert.Equal(suite.T(), input.CreatedAt, newInput.CreatedAt)
	assert.Equal(suite.T(), input.UpdatedAt, newInput.UpdatedAt)
}
