package usecase

import (
	"fmt"
	"testing"

	"libs/golang/ddd/domain/entities/input-broker/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/input-broker/repository"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	mockevent "libs/golang/ddd/events/event-mock/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateInputUseCaseSuite struct {
	suite.Suite
	repoMock       *mockrepository.InputRepositoryMock
	eventMock      *mockevent.MockEvent
	dispatcherMock *mockevent.MockEventDispatcher
	useCase        *CreateInputUseCase
	inputDTO       inputdto.InputDTO
	inputProps     entity.InputProps
}

func TestCreateInputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CreateInputUseCaseSuite))
}

func (suite *CreateInputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.InputRepositoryMock)
	suite.eventMock = new(mockevent.MockEvent)
	suite.dispatcherMock = new(mockevent.MockEventDispatcher)
	suite.useCase = NewCreateInputUseCase(suite.repoMock, suite.eventMock, suite.dispatcherMock)
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

func (suite *CreateInputUseCaseSuite) TestExecuteWhenSuccess() {
	expectedInput, _ := entity.NewInput(suite.inputProps)
	suite.repoMock.On("Create", expectedInput).Return(nil)
	suite.eventMock.On("SetPayload", mock.Anything).Return()
	suite.dispatcherMock.On("Dispatch", suite.eventMock, fmt.Sprintf("input.created.%s.%s", suite.inputDTO.Service, suite.inputDTO.Source)).Return(nil)

	output, err := suite.useCase.Execute(suite.inputDTO)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.inputDTO.Provider, output.Metadata.Provider)
	assert.Equal(suite.T(), suite.inputDTO.Service, output.Metadata.Service)
	assert.Equal(suite.T(), suite.inputDTO.Source, output.Metadata.Source)
	assert.Equal(suite.T(), suite.inputDTO.Data, output.Data)
	suite.repoMock.AssertExpectations(suite.T())
	suite.eventMock.AssertExpectations(suite.T())
	suite.dispatcherMock.AssertExpectations(suite.T())
}

func (suite *CreateInputUseCaseSuite) TestExecuteWhenErrorCreatingInput() {
	expectedInput, _ := entity.NewInput(suite.inputProps)
	suite.repoMock.On("Create", expectedInput).Return(fmt.Errorf("Input with ID: %s already exists", expectedInput.ID))

	input, err := suite.useCase.Execute(suite.inputDTO)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.InputDTO{}, input)
	suite.repoMock.AssertExpectations(suite.T())
	suite.eventMock.AssertNotCalled(suite.T(), "SetPayload", mock.Anything)
	suite.dispatcherMock.AssertNotCalled(suite.T(), "Dispatch", mock.Anything, mock.Anything)
}
