package usecase

import (
	"libs/golang/ddd/domain/entities/config-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/config-vault/repository"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListOneByIDConfigUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.ConfigRepositoryMock
	useCase  *ListOneByIDConfigUseCase
}

func TestListOneByIDConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListOneByIDConfigUseCaseSuite))
}

func (suite *ListOneByIDConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewListOneByIDConfigUseCase(suite.repoMock)
}

func (suite *ListOneByIDConfigUseCaseSuite) TestExecutewhenSuccess() {
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
		{
			ID:              "2",
			Active:          true,
			Service:         "service2",
			Source:          "source2",
			Provider:        "provider2",
			ConfigVersionID: "v2",
			JobParameters: entity.JobParameters{
				ParserModule: "parser_module2",
			},
			DependsOn: []entity.JobDependencies{
				{Service: "dep_service2", Source: "dep_source2"},
			},
			CreatedAt: "2023-06-02T00:00:00Z",
			UpdatedAt: "2023-06-02T00:00:00Z",
		},
	}

	suite.repoMock.On("FindByID", "1").Return(entityConfigs[0], nil)

	expectedOuput := outputdto.ConfigDTO{
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
	}

	output, err := suite.useCase.Execute("1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOuput, output)
	suite.repoMock.AssertExpectations(suite.T())

}
