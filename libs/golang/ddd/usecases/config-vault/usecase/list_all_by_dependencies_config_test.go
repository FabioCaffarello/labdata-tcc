package usecase

import (
	"fmt"
	"libs/golang/ddd/domain/entities/config-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/config-vault/repository"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListAllByDependsOnConfigUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.ConfigRepositoryMock
	useCase  *ListAllByDependsOnConfigUseCase
}

func TestListAllByDependsOnConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllByDependsOnConfigUseCaseSuite))
}

func (suite *ListAllByDependsOnConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewListAllByDependsOnConfigUseCase(suite.repoMock)
}

func (suite *ListAllByDependsOnConfigUseCaseSuite) TestExecuteWhenSuccess() {
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
	}

	dependsOn := map[string]interface{}{
		"service": "service1",
		"source":  "source1",
	}

	suite.repoMock.On("FindAllByDependsOn", dependsOn).Return(entityConfigs, nil)

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
	}

	output, err := suite.useCase.Execute("service1", "source1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllByDependsOnConfigUseCaseSuite) TestExecuteWhenError() {
	dependsOn := map[string]interface{}{
		"service": "service1",
		"source":  "source1",
	}

	suite.repoMock.On("FindAllByDependsOn", dependsOn).Return(nil, fmt.Errorf("error"))

	output, err := suite.useCase.Execute("service1", "source1")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.ConfigDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
