package usecase

import (
	"fmt"
	"testing"

	"libs/golang/ddd/domain/entities/config-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/config-vault/repository"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListAllByServiceAndSourceAndProviderConfigUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.ConfigRepositoryMock
	useCase  *ListAllByServiceAndSourceAndProviderConfigUseCase
}

func TestListAllByServiceAndSourceAndProviderConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllByServiceAndSourceAndProviderConfigUseCaseSuite))
}

func (suite *ListAllByServiceAndSourceAndProviderConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewListAllByServiceAndSourceAndProviderConfigUseCase(suite.repoMock)
}

func (suite *ListAllByServiceAndSourceAndProviderConfigUseCaseSuite) TestExecuteWhenSuccess() {
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

	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "service1", "source1", "provider1").Return(entityConfigs, nil)

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

	output, err := suite.useCase.Execute("service1", "source1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllByServiceAndSourceAndProviderConfigUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "service1", "source1", "provider1").Return(nil, fmt.Errorf("error"))

	output, err := suite.useCase.Execute("service1", "source1", "provider1")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.ConfigDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
