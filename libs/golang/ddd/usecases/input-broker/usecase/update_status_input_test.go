package usecase

import (
	"fmt"
	"testing"

	"libs/golang/ddd/domain/entities/input-broker/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/input-broker/repository"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UpdateStatusInputUseCaseSuite struct {
	suite.Suite
	repoMock   *mockrepository.InputRepositoryMock
	useCase    *UpdateStatusInputUseCase
	statusDTO  shareddto.StatusDTO
	inputProps entity.InputProps
}

func TestUpdateStatusInputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UpdateStatusInputUseCaseSuite))
}

func (suite *UpdateStatusInputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.InputRepositoryMock)
	suite.useCase = NewUpdateStatusInputUseCase(suite.repoMock)
	suite.statusDTO = shareddto.StatusDTO{
		Code:   200,
		Detail: "Success",
	}
	suite.inputProps = entity.InputProps{
		Provider: "TestProvider",
		Service:  "TestService",
		Source:   "TestSource",
		Data:     map[string]interface{}{"key": "value"},
	}
}

func (suite *UpdateStatusInputUseCaseSuite) TestExecuteWhenSuccess() {
	expectedInput, _ := entity.NewInput(suite.inputProps)
	suite.repoMock.On("FindByID", expectedInput.GetEntityID()).Return(expectedInput, nil)
	suite.repoMock.On("Update", expectedInput).Return(nil)

	output, err := suite.useCase.Execute(expectedInput.GetEntityID(), suite.statusDTO)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.statusDTO.Code, output.Status.Code)
	assert.Equal(suite.T(), suite.statusDTO.Detail, output.Status.Detail)

	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *UpdateStatusInputUseCaseSuite) TestExecuteError() {
	expectedInput, _ := entity.NewInput(suite.inputProps)
	suite.repoMock.On("FindByID", expectedInput.GetEntityID()).Return(nil, fmt.Errorf("database error"))

	output, err := suite.useCase.Execute(expectedInput.GetEntityID(), suite.statusDTO)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.InputDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
