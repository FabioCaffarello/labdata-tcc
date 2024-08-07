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

type ListAllByStatusAndSourceAndProviderInputUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.InputRepositoryMock
	useCase  *ListAllByStatusAndSourceAndProviderInputUseCase
}

func TestListAllByStatusAndSourceAndProviderInputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllByStatusAndSourceAndProviderInputUseCaseSuite))
}

func (suite *ListAllByStatusAndSourceAndProviderInputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.InputRepositoryMock)
	suite.useCase = NewListAllByStatusAndSourceAndProviderInputUseCase(suite.repoMock)
}

func (suite *ListAllByStatusAndSourceAndProviderInputUseCaseSuite) TestExecuteWhenSuccess() {
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

	suite.repoMock.On("FindAllByStatusAndSourceAndProvider", "test_source", "test_provider", 0).Return(entityInputs, nil)

	expectedOutput := []outputdto.InputDTO{
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

	output, err := suite.useCase.Execute("test_provider", "test_source", 0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
}

func (suite *ListAllByStatusAndSourceAndProviderInputUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAllByStatusAndSourceAndProvider", "test_source", "test_provider", 0).Return(nil, fmt.Errorf("error"))

	output, err := suite.useCase.Execute("test_provider", "test_source", 0)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.InputDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
