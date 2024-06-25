package converter

import (
	"libs/golang/ddd/domain/entities/output-vault/entity"
	shareddto "libs/golang/ddd/dtos/output-vault/shared"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OutputConverterEntityToDTOSuite struct {
	suite.Suite
}

func TestOutputConverterEntityToDTOSuite(t *testing.T) {
	suite.Run(t, new(OutputConverterEntityToDTOSuite))
}

func (suite *OutputConverterEntityToDTOSuite) TestConvertMetadataEntityToDTO() {
	entity := entity.Metadata{
		InputID: "input-id",
		Input: entity.Input{
			Data:                map[string]interface{}{"key": "value"},
			ProcessingID:        "processing-id",
			ProcessingTimestamp: "2021-06-01T00:00:00Z",
		},
	}

	expected := shareddto.MetadataDTO{
		InputID: "input-id",
		Input: shareddto.InputDTO{
			Data:                map[string]interface{}{"key": "value"},
			ProcessingID:        "processing-id",
			ProcessingTimestamp: "2021-06-01T00:00:00Z",
		},
	}

	dtoMetadata := ConvertMetadataEntityToDTO(entity)

	suite.Equal(expected, dtoMetadata)
}

func (suite *OutputConverterEntityToDTOSuite) TestConvertMetadataEntityToDTOWhenEmpty() {
	entity := entity.Metadata{}

	expected := shareddto.MetadataDTO{}

	dtoMetadata := ConvertMetadataEntityToDTO(entity)

	suite.Equal(expected, dtoMetadata)
}
