package usecase

import (
	"fmt"
	"testing"

	mockrepository "libs/golang/ddd/domain/repositories/database/mock/input-broker/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeleteInputUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.InputRepositoryMock
	useCase  *DeleteInputUseCase
}

func TestDeleteInputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DeleteInputUseCaseSuite))
}

func (suite *DeleteInputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.InputRepositoryMock)
	suite.useCase = NewDeleteInputUseCase(suite.repoMock)
}

func (suite *DeleteInputUseCaseSuite) TestExecuteWhenSuccess() {
	inputID := "test_id"
	suite.repoMock.On("Delete", inputID).Return(nil)

	err := suite.useCase.Execute(inputID)

	assert.Nil(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *DeleteInputUseCaseSuite) TestExecuteWhenError() {
	inputID := "test_id"
	suite.repoMock.On("Delete", inputID).Return(fmt.Errorf("Input with ID: %s not found", inputID))

	err := suite.useCase.Execute(inputID)

	assert.NotNil(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}
