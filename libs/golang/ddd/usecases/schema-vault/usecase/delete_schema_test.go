package usecase

import (
	"fmt"
	"testing"

	mockrepository "libs/golang/ddd/domain/repositories/database/mock/schema-vault/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeleteSchemaUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.SchemaRepositoryMock
	useCase  *DeleteSchemaUseCase
}

func TestDeleteSchemaUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DeleteSchemaUseCaseSuite))
}

func (suite *DeleteSchemaUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.SchemaRepositoryMock)
	suite.useCase = NewDeleteSchemaUseCase(suite.repoMock)
}

func (suite *DeleteSchemaUseCaseSuite) TestExecuteWhenSuccess() {
	schemaID := "test_id"
	suite.repoMock.On("Delete", schemaID).Return(nil)

	err := suite.useCase.Execute(schemaID)

	assert.Nil(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *DeleteSchemaUseCaseSuite) TestExecuteWhenError() {
	schemaID := "test_id"
	suite.repoMock.On("Delete", schemaID).Return(fmt.Errorf("Schema with ID: %s not found", schemaID))

	err := suite.useCase.Execute(schemaID)

	assert.NotNil(suite.T(), err)
	suite.repoMock.AssertExpectations(suite.T())
}
