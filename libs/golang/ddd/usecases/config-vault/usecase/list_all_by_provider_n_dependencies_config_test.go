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

type ListAllByProviderAndDependsOnConfigUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.ConfigRepositoryMock
	useCase  *ListAllByProviderAndDependsOnConfigUseCase
}

func TestListAllByProviderAndDependsOnConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllByProviderAndDependsOnConfigUseCaseSuite))
}

func (suite *ListAllByProviderAndDependsOnConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewListAllByProviderAndDependsOnConfigUseCase(suite.repoMock)
}

func (suite *ListAllByProviderAndDependsOnConfigUseCaseSuite) TestExecuteWhenSuccess() {
	entityConfigs := []*entity.Config{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			JobParameters: entity.JobParameters{
				ParserModule: "parser_module1",
			},
			DependsOn: []entity.JobDependencies{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllByProviderAndDependsOn", "provider1", "dep_service1", "dep_source1").Return(entityConfigs, nil)

	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			JobParameters: shareddto.JobParametersDTO{
				ParserModule: "parser_module1",
			},
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	output, err := suite.useCase.Execute("provider1", "dep_service1", "dep_source1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllByProviderAndDependsOnConfigUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAllByProviderAndDependsOn", "provider1", "service1", "source1").Return(nil, fmt.Errorf("error"))

	output, err := suite.useCase.Execute("provider1", "service1", "source1")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.ConfigDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
