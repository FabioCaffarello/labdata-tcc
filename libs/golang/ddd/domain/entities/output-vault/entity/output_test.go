package entity

import (
	md5id "libs/golang/shared/id/go-md5"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OutputVaultConfigSuite struct {
	suite.Suite
}

func TestOutputVaultConfigSuite(t *testing.T) {
	suite.Run(t, new(OutputVaultConfigSuite))
}

func (suite *OutputVaultConfigSuite) TestNewOutput() {
	outputProps := OutputProps{
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

	output, err := NewOutput(outputProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
	assert.Equal(suite.T(), outputProps.Service, output.Service)
	assert.Equal(suite.T(), outputProps.Source, output.Source)
	assert.Equal(suite.T(), outputProps.Provider, output.Provider)
	assert.Equal(suite.T(), outputProps.Data, output.Data)
	assert.Equal(suite.T(), outputProps.Metadata["input_id"], output.Metadata.InputID)
}

func (suite *OutputVaultConfigSuite) TestInvalidOutput() {
	outputProps := OutputProps{
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Metadata: map[string]interface{}{
			"input": map[string]interface{}{
				"data": map[string]interface{}{
					"input1": "value1",
				},
				"processing_id":        "processing_id",
				"processing_timestamp": "2023-06-01 00:00:00",
			},
		},
	}

	_, err := NewOutput(outputProps)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), ErrInvalidInputID, err)
}

func (suite *OutputVaultConfigSuite) TestToMap() {
	outputProps := OutputProps{
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

	output, err := NewOutput(outputProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)

	doc, err := output.ToMap()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), doc)
	assert.IsType(suite.T(), map[string]interface{}{}, doc)

	assert.Equal(suite.T(), output.Service, doc["service"])
	assert.Equal(suite.T(), output.Source, doc["source"])
	assert.Equal(suite.T(), output.Provider, doc["provider"])
	assert.Equal(suite.T(), output.Data, doc["data"])
	assert.Equal(suite.T(), output.Metadata.InputID, doc["metadata"].(map[string]interface{})["input_id"])
	assert.Equal(suite.T(), output.Metadata.Input.ProcessingID, doc["metadata"].(map[string]interface{})["input"].(map[string]interface{})["processing_id"])
	assert.Equal(suite.T(), output.Metadata.Input.ProcessingTimestamp, doc["metadata"].(map[string]interface{})["input"].(map[string]interface{})["processing_timestamp"])
	assert.Equal(suite.T(), output.Metadata.Input.Data, doc["metadata"].(map[string]interface{})["input"].(map[string]interface{})["data"])
}

func (suite *OutputVaultConfigSuite) TestIsValid() {
	idData := getIDData("test_service", "test_source", "test_provider", map[string]interface{}{"field1": "value1"})
	output := &Output{
		ID:       md5id.NewID(idData),
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Data:     map[string]interface{}{"field1": "value1"},
		Metadata: Metadata{
			InputID: "input_id",
			Input: Input{
				ProcessingID:        "processing_id",
				ProcessingTimestamp: "2023-06-01 00:00:00",
				Data: map[string]interface{}{
					"input1": "value1",
				},
			},
		},
		CreatedAt: time.Now().Format(dateLayout),
		UpdatedAt: time.Now().Format(dateLayout),
	}

	err := output.isValid()
	assert.Nil(suite.T(), err)

	output.Service = ""
	err = output.isValid()
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), ErrInvalidService, err)
}

func (suite *OutputVaultConfigSuite) TestMapToEntity() {
	outputProps := OutputProps{
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

	output, err := NewOutput(outputProps)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)

	doc, err := output.ToMap()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), doc)
	assert.IsType(suite.T(), map[string]interface{}{}, doc)

	newOuput := &Output{}
	newOuput, err = newOuput.MapToEntity(doc)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), newOuput)

	assert.Equal(suite.T(), output.ID, newOuput.ID)
	assert.Equal(suite.T(), output.Service, newOuput.Service)
	assert.Equal(suite.T(), output.Source, newOuput.Source)
	assert.Equal(suite.T(), output.Provider, newOuput.Provider)
	assert.Equal(suite.T(), output.Data, newOuput.Data)
	assert.Equal(suite.T(), output.Metadata.InputID, newOuput.Metadata.InputID)
	assert.Equal(suite.T(), output.Metadata.Input.ProcessingID, newOuput.Metadata.Input.ProcessingID)
	assert.Equal(suite.T(), output.Metadata.Input.ProcessingTimestamp, newOuput.Metadata.Input.ProcessingTimestamp)
	assert.Equal(suite.T(), output.Metadata.Input.Data, newOuput.Metadata.Input.Data)
}
