package usecase

import (
	"fmt"
	"testing"

	"libs/golang/ddd/domain/entities/input-broker/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/input-broker/repository"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UpdateInputUseCaseSuite struct {
	suite.Suite
	repoMock   *mockrepository.InputRepositoryMock
	useCase    *UpdateInputUseCase
	inputDTO   inputdto.InputDTO
	inputProps entity.InputProps
}

func TestUpdateInputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UpdateInputUseCaseSuite))
}

func (suite *UpdateInputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.InputRepositoryMock)
	suite.useCase = NewUpdateInputUseCase(suite.repoMock)
	suite.inputDTO = inputdto.InputDTO{
		Provider: "test_provider",
		Service:  "test_service",
		Source:   "test_source",
		Data:     map[string]interface{}{"key": "value"},
	}
	suite.inputProps = entity.InputProps{
		Provider: suite.inputDTO.Provider,
		Service:  suite.inputDTO.Service,
		Source:   suite.inputDTO.Source,
		Data:     suite.inputDTO.Data,
	}
}

func (suite *UpdateInputUseCaseSuite) TestExecuteWhenSuccess() {
	expectedInput, _ := entity.NewInput(suite.inputProps)
	suite.repoMock.On("Update", expectedInput).Return(nil)

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.inputDTO.Provider, output.Metadata.Provider)
	assert.Equal(suite.T(), suite.inputDTO.Service, output.Metadata.Service)
	assert.Equal(suite.T(), suite.inputDTO.Source, output.Metadata.Source)
	assert.Equal(suite.T(), suite.inputDTO.Data, output.Data)

	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *UpdateInputUseCaseSuite) TestExecuteError() {
	expectedInput, _ := entity.NewInput(suite.inputProps)
	suite.repoMock.On("Update", expectedInput).Return(fmt.Errorf("Input with ID: %s does not exist", expectedInput.ID))

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.InputDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
