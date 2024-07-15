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

type ListAllByServiceAndProviderAndActiveConfigUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.ConfigRepositoryMock
	useCase  *ListAllByServiceAndProviderAndActiveConfigUseCase
}

func TestListAllByServiceAndProviderAndActiveConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllByServiceAndProviderAndActiveConfigUseCaseSuite))
}

func (suite *ListAllByServiceAndProviderAndActiveConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewListAllByServiceAndProviderAndActiveConfigUseCase(suite.repoMock)
}

func (suite *ListAllByServiceAndProviderAndActiveConfigUseCaseSuite) TestExecuteWhenSuccess() {
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

	suite.repoMock.On("FindAllByServiceAndProviderAndActive", "service1", "provider1", true).Return(entityConfigs, nil)

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

	output, err := suite.useCase.Execute("service1", "provider1", true)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllByServiceAndProviderAndActiveConfigUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAllByServiceAndProviderAndActive", "service1", "provider1", true).Return(nil, fmt.Errorf("error"))

	output, err := suite.useCase.Execute("service1", "provider1", true)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.ConfigDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
