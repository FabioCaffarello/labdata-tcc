package usecase

import (
	"fmt"
	"testing"

	mockrepository "libs/golang/ddd/domain/repositories/database/mock/config-vault/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeleteConfigUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.ConfigRepositoryMock
	useCase  *DeleteConfigUseCase
}

func TestDeleteConfigUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DeleteConfigUseCaseSuite))
}

func (suite *DeleteConfigUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.useCase = NewDeleteConfigUseCase(suite.repoMock)
}

func (suite *DeleteConfigUseCaseSuite) TestExecuteWhenSuccess() {
	configID := "test_id"
	suite.repoMock.On("Delete", configID).Return(nil)

	err := suite.useCase.Execute(configID)

	assert.Nil(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *DeleteConfigUseCaseSuite) TestExecuteWhenError() {
	configID := "test_id"
	suite.repoMock.On("Delete", configID).Return(fmt.Errorf("Config with ID: %s not found", configID))

	err := suite.useCase.Execute(configID)

	assert.NotNil(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}
