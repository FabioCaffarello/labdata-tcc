package converter

import (
	"libs/golang/ddd/domain/entities/input-broker/entity"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InputConverterDTOToEntitySuite struct {
	suite.Suite
}

func TestInputConverterDTOToEntitySuite(t *testing.T) {
	suite.Run(t, new(InputConverterDTOToEntitySuite))
}

func (suite *InputConverterDTOToEntitySuite) TestConvertMetadataDTOToEntity() {
	metadataDTO := shareddto.MetadataDTO{
		Provider:            "test_provider",
		Service:             "test_service",
		Source:              "test_source",
		ProcessingID:        "test_processing_id",
		ProcessingTimestamp: "2023-07-02T12:34:56Z",
	}

	expectedMetadata := entity.Metadata{
		Provider:            "test_provider",
		Service:             "test_service",
		Source:              "test_source",
		ProcessingID:        "test_processing_id",
		ProcessingTimestamp: "2023-07-02T12:34:56Z",
	}

	result := ConvertMetadataDTOToEntity(metadataDTO)
	assert.Equal(suite.T(), expectedMetadata, result)
}

func (suite *InputConverterDTOToEntitySuite) TestConvertStatusDTOToEntity() {
	statusDTO := shareddto.StatusDTO{
		Code:   200,
		Detail: "OK",
	}

	expectedStatus := entity.Status{
		Code:   200,
		Detail: "OK",
	}

	result := ConvertStatusDTOToEntity(statusDTO)
	assert.Equal(suite.T(), expectedStatus, result)
}
