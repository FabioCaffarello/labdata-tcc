package usecase

import (
	"fmt"
	"testing"

	"libs/golang/ddd/domain/entities/output-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/output-vault/repository"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	shareddto "libs/golang/ddd/dtos/output-vault/shared"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListAllOutputUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.OutputRepositoryMock
	useCase  *ListAllOutputUseCase
}

func TestListAllOutputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllOutputUseCaseSuite))
}

func (suite *ListAllOutputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.OutputRepositoryMock)
	suite.useCase = NewListAllOutputUseCase(suite.repoMock)
}

func (suite *ListAllOutputUseCaseSuite) TestExecuteWhenSuccess() {
	entityOutputs := []*entity.Output{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: entity.Metadata{
				InputID: "input1",
				Input: entity.Input{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAll").Return(entityOutputs, nil)

	expectedOutput := []outputdto.OutputDTO{
		{
			ID:       "1",
			Service:  "service1",
			Source:   "source1",
			Provider: "provider",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
			Metadata: shareddto.MetadataDTO{
				InputID: "input1",
				Input: shareddto.InputDTO{
					Data:                map[string]interface{}{"key": "value"},
					ProcessingID:        "processing1",
					ProcessingTimestamp: "2021-06-01 00:00:00",
				},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	output, err := suite.useCase.Execute()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllOutputUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAll").Return(nil, fmt.Errorf("error"))

	output, err := suite.useCase.Execute()

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.OutputDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
