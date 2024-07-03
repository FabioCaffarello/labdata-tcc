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

type ListAllByStatusAndProviderInputUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.InputRepositoryMock
	useCase  *ListAllByStatusAndProviderInputUseCase
}

func TestListAllByStatusAndProviderInputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllByStatusAndProviderInputUseCaseSuite))
}

func (suite *ListAllByStatusAndProviderInputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.InputRepositoryMock)
	suite.useCase = NewListAllByStatusAndProviderInputUseCase(suite.repoMock)
}

func (suite *ListAllByStatusAndProviderInputUseCaseSuite) TestExecuteWhenSuccess() {
	entityInputs := []*entity.Input{
		{
			ID: "test_id2",
			Metadata: entity.Metadata{
				Service:             "test_service2",
				Source:              "test_source2",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id2",
				ProcessingTimestamp: "2023-06-02 00:00:00",
			},
			Status: entity.Status{
				Code:   200,
				Detail: "test_detail2",
			},
			Data: map[string]interface{}{
				"key2": "value2",
			},
			CreatedAt: "2023-06-02 00:00:00",
			UpdatedAt: "2023-06-02 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByStatusAndProvider", "test_provider", 200).Return(entityInputs, nil)

	expectedOutput := []outputdto.InputDTO{
		{
			ID: "test_id2",
			Metadata: shareddto.MetadataDTO{
				Service:             "test_service2",
				Source:              "test_source2",
				Provider:            "test_provider",
				ProcessingID:        "test_processing_id2",
				ProcessingTimestamp: "2023-06-02 00:00:00",
			},
			Status: shareddto.StatusDTO{
				Code:   200,
				Detail: "test_detail2",
			},
			Data: map[string]interface{}{
				"key2": "value2",
			},
			CreatedAt: "2023-06-02 00:00:00",
			UpdatedAt: "2023-06-02 00:00:00",
		},
	}

	output, err := suite.useCase.Execute("test_provider", 200)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllByStatusAndProviderInputUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAllByStatusAndProvider", "test_provider", 200).Return(nil, fmt.Errorf("database error"))

	output, err := suite.useCase.Execute("test_provider", 200)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.InputDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
