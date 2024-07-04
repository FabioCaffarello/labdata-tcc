package usecase

import (
	"fmt"
	"testing"

	"libs/golang/ddd/domain/entities/input-broker/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/input-broker/repository"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListAllByServiceAndProviderInputUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.InputRepositoryMock
	useCase  *ListAllByServiceAndProviderInputUseCase
}

func TestListAllByServiceAndProviderInputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllByServiceAndProviderInputUseCaseSuite))
}

func (suite *ListAllByServiceAndProviderInputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.InputRepositoryMock)
	suite.useCase = NewListAllByServiceAndProviderInputUseCase(suite.repoMock)
}

func (suite *ListAllByServiceAndProviderInputUseCaseSuite) TestExecuteWhenSuccess() {
	entityInputs := []*entity.Input{
		{
			ID: "test_id",
			Metadata: entity.Metadata{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: entity.Status{
				Code:   0,
				Detail: "test_detail",
			},
			Data: map[string]interface{}{
				"key": "value",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return(entityInputs, nil)

	expectedInput := []outputdto.InputDTO{
		{
			ID: "test_id",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service",
				Source:              "test_source",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id",
				ProcessingTimestamp: "2023-06-01 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   0,
				Detail: "test_detail",
			},
			Data: map[string]interface{}{
				"key": "value",
			},
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	output, err := suite.useCase.Execute("test_provider", "test_service")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedInput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllByServiceAndProviderInputUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return(nil, fmt.Errorf("No inputs found for provider: %s and service: %s", "test_provider", "test_service"))

	output, err := suite.useCase.Execute("test_provider", "test_service")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.InputDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
