package usecase

import (
	"fmt"
	"testing"

	"libs/golang/ddd/domain/entities/config-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/config-vault/repository"
	inputdto "libs/golang/ddd/dtos/config-vault/input"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"libs/golang/ddd/shared/type-tools/custom-types-converter/config-vault/converter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UpdateConfigUseCaseSuite struct {
	suite.Suite
	repoMock    *mockrepository.ConfigRepositoryMock
	useCase     *UpdateConfigUseCase
	inputDTO    inputdto.ConfigDTO
	configProps entity.ConfigProps
}

func TestUpdateConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UpdateConfigUseCaseSuite))
}

func (suite *UpdateConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewUpdateConfigUseCase(suite.repoMock)
	suite.inputDTO = inputdto.ConfigDTO{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
	}
	suite.configProps = entity.ConfigProps{
		Active:    suite.inputDTO.Active,
		Service:   suite.inputDTO.Service,
		Source:    suite.inputDTO.Source,
		Provider:  suite.inputDTO.Provider,
		DependsOn: converter.ConvertJobDependenciesDTOToMap(suite.inputDTO.DependsOn),
	}
}

func (suite *UpdateConfigUseCaseSuite) TestExecuteWhenSuccess() {
	expectedConfig, _ := entity.NewConfig(suite.configProps)
	suite.repoMock.On("Update", expectedConfig).Return(nil)

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.inputDTO.Active, output.Active)
	assert.Equal(suite.T(), suite.inputDTO.Service, output.Service)
	assert.Equal(suite.T(), suite.inputDTO.Source, output.Source)
	assert.Equal(suite.T(), suite.inputDTO.Provider, output.Provider)
	assert.Equal(suite.T(), suite.inputDTO.DependsOn[0].Service, output.DependsOn[0].Service)
	assert.Equal(suite.T(), suite.inputDTO.DependsOn[0].Source, output.DependsOn[0].Source)

	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *UpdateConfigUseCaseSuite) TestExecuteError() {
	expectedConfig, _ := entity.NewConfig(suite.configProps)
	suite.repoMock.On("Update", expectedConfig).Return(fmt.Errorf("Config with ID: %s does not exist", expectedConfig.ID))

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.ConfigDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
