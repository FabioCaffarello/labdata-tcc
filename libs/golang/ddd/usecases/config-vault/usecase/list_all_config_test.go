package usecase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"libs/golang/ddd/domain/entities/config-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/config-vault/repository"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
)

type ListAllConfigUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.ConfigRepositoryMock
	useCase  *ListAllConfigUseCase
}

func TestListAllConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllConfigUseCaseSuite))
}

func (suite *ListAllConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewListAllConfigUseCase(suite.repoMock)
}

func (suite *ListAllConfigUseCaseSuite) TestExecutewhenSuccess() {
	entityConfigs := []*entity.Config{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			DependsOn: []entity.JobDependencies{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
		{
			ID:              "2",
			Active:          true,
			Service:         "service2",
			Source:          "source2",
			Provider:        "provider2",
			ConfigVersionID: "v2",
			DependsOn: []entity.JobDependencies{
				{Service: "dep_service2", Source: "dep_source2"},
			},
			CreatedAt: "2023-06-02T00:00:00Z",
			UpdatedAt: "2023-06-02T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAll").Return(entityConfigs, nil)

	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
		{
			ID:              "2",
			Active:          true,
			Service:         "service2",
			Source:          "source2",
			Provider:        "provider2",
			ConfigVersionID: "v2",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service2", Source: "dep_source2"},
			},
			CreatedAt: "2023-06-02T00:00:00Z",
			UpdatedAt: "2023-06-02T00:00:00Z",
		},
	}

	output, err := suite.useCase.Execute()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllConfigUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAll").Return(nil, fmt.Errorf("database error"))

	output, err := suite.useCase.Execute()

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.ConfigDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
