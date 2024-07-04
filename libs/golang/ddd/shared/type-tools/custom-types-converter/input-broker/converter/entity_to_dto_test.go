package converter

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InputConverterEntityToDTOSuite struct {
	suite.Suite
}

func TestConfigConverterEntityToDTOSuite(t *testing.T) {
	suite.Run(t, new(InputConverterEntityToDTOSuite))
}

func (suite *InputConverterEntityToDTOSuite) TestConvertMetadataEntityToDTO() {
	metadata := entity.Metadata{
		Provider:            "test_provider",
		Service:             "test_service",
		Source:              "test_source",
		ProcessingID:        "test_processing_id",
		ProcessingTimestamp: "2023-07-02T12:34:56Z",
	}

	expectedMetadataDTO := shareddto.MetadataDTO{
		Provider:            "test_provider",
		Service:             "test_service",
		Source:              "test_source",
		ProcessingID:        "test_processing_id",
		ProcessingTimestamp: "2023-07-02T12:34:56Z",
	}

	result := ConvertMetadataEntityToDTO(metadata)
	assert.Equal(suite.T(), expectedMetadataDTO, result)
}

func (suite *InputConverterEntityToDTOSuite) TestConvertStatusEntityToDTO() {
	status := entity.Status{
		Code:   200,
		Detail: "OK",
	}

	expectedStatusDTO := shareddto.StatusDTO{
		Code:   200,
		Detail: "OK",
	}

	result := ConvertStatusEntityToDTO(status)
	assert.Equal(suite.T(), expectedStatusDTO, result)
}
