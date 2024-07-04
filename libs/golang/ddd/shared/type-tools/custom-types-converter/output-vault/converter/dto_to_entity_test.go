package converter

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	shareddto "libs/golang/ddd/dtos/output-vault/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OutputConverterDTOToEntitySuite struct {
	suite.Suite
}

func TestOutputConverterDTOToEntitySuite(t *testing.T) {
	suite.Run(t, new(OutputConverterDTOToEntitySuite))
}

func (suite *OutputConverterDTOToEntitySuite) TestConvertMetadataDTOToEntity() {
	dto := shareddto.MetadataDTO{
		InputID: "input-id",
		Input: shareddto.InputDTO{
			Data:                map[string]interface{}{"key": "value"},
			ProcessingID:        "processing-id",
			ProcessingTimestamp: "2021-06-01T00:00:00Z",
		},
	}

	expected := entity.Metadata{
		InputID: "input-id",
		Input: entity.Input{
			Data:                map[string]interface{}{"key": "value"},
			ProcessingID:        "processing-id",
			ProcessingTimestamp: "2021-06-01T00:00:00Z",
		},
	}

	entityMetadata := ConvertMetadataDTOToEntity(dto)

	assert.Equal(suite.T(), expected, entityMetadata)
}

func (suite *OutputConverterDTOToEntitySuite) TestConvertMetadataDTOToMap() {
	dto := shareddto.MetadataDTO{
		InputID: "input-id",
		Input: shareddto.InputDTO{
			Data:                map[string]interface{}{"key": "value"},
			ProcessingID:        "processing-id",
			ProcessingTimestamp: "2021-06-01T00:00:00Z",
		},
	}

	expected := map[string]interface{}{
		"input_id": "input-id",
		"input": map[string]interface{}{
			"data":                 map[string]interface{}{"key": "value"},
			"processing_id":        "processing-id",
			"processing_timestamp": "2021-06-01T00:00:00Z",
		},
	}

	entityMetadata := ConvertMetadataDTOToMap(dto)

	assert.Equal(suite.T(), expected, entityMetadata)
}
