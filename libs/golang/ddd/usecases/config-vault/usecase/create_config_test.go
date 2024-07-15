package usecase

import (
	"fmt"
	"libs/golang/ddd/domain/entities/config-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/config-vault/repository"
	inputdto "libs/golang/ddd/dtos/config-vault/input"

	outputdto "libs/golang/ddd/dtos/config-vault/output"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CreateConfigUseCaseSuite struct {
	suite.Suite
	repoMock    *mockrepository.ConfigRepositoryMock
	useCase     *CreateConfigUseCase
	inputDTO    inputdto.ConfigDTO
	configProps entity.ConfigProps
}

func TestCreateConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CreateConfigUseCaseSuite))
}

func (suite *CreateConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewCreateConfigUseCase(suite.repoMock)
	suite.inputDTO = inputdto.ConfigDTO{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		JobParameters: shareddto.JobParametersDTO{
			ParserModule: "test_parser_module",
		},
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
	}
	suite.configProps = entity.ConfigProps{
		Active:        suite.inputDTO.Active,
		Service:       suite.inputDTO.Service,
		Source:        suite.inputDTO.Source,
		Provider:      suite.inputDTO.Provider,
		JobParameters: converter.ConvertJobParametersDTOToMap(suite.inputDTO.JobParameters),
		DependsOn:     converter.ConvertJobDependenciesDTOToMap(suite.inputDTO.DependsOn),
	}
}

func (suite *CreateConfigUseCaseSuite) TestExecuteWhenSuccess() {
	expectedConfig, _ := entity.NewConfig(suite.configProps)
	suite.repoMock.On("Create", expectedConfig).Return(nil)

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.inputDTO.Active, output.Active)
	assert.Equal(suite.T(), suite.inputDTO.Service, output.Service)
	assert.Equal(suite.T(), suite.inputDTO.Source, output.Source)
	assert.Equal(suite.T(), suite.inputDTO.Provider, output.Provider)
	assert.Equal(suite.T(), suite.inputDTO.DependsOn[0].Service, output.DependsOn[0].Service)
	assert.Equal(suite.T(), suite.inputDTO.DependsOn[0].Source, output.DependsOn[0].Source)
	assert.Equal(suite.T(), suite.inputDTO.JobParameters.ParserModule, output.JobParameters.ParserModule)

	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *CreateConfigUseCaseSuite) TestExecuteError() {
	expectedConfig, _ := entity.NewConfig(suite.configProps)
	suite.repoMock.On("Create", expectedConfig).Return(fmt.Errorf("Config with ID: %s already exists", expectedConfig.ID))

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.ConfigDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
