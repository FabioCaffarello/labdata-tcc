package usecase

import (
	"fmt"
	"testing"

	mockrepository "libs/golang/ddd/domain/repositories/database/mock/output-vault/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeleteOutputUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.OutputRepositoryMock
	useCase  *DeleteOutputUseCase
}

func TestDeleteOutputUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DeleteOutputUseCaseSuite))
}

func (suite *DeleteOutputUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.OutputRepositoryMock)
	suite.useCase = NewDeleteOutputUseCase(suite.repoMock)
}

func (suite *DeleteOutputUseCaseSuite) TestExecuteWhenSuccess() {
	outputID := "test_id"
	suite.repoMock.On("Delete", outputID).Return(nil)

	err := suite.useCase.Execute(outputID)

	assert.Nil(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *DeleteOutputUseCaseSuite) TestExecuteWhenError() {
	outputID := "test_id"
	suite.repoMock.On("Delete", outputID).Return(fmt.Errorf("Output with ID: %s not found", outputID))

	err := suite.useCase.Execute(outputID)

	assert.NotNil(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}
