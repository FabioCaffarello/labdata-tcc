package usecase

import (
	"fmt"
	"testing"

	"libs/golang/ddd/domain/entities/schema-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/schema-vault/repository"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListAllSchemaUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.SchemaRepositoryMock
	useCase  *ListAllSchemaUseCase
}

func TestListAllSchemaUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAllSchemaUseCaseSuite))
}

func (suite *ListAllSchemaUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.SchemaRepositoryMock)
	suite.useCase = NewListAllSchemaUseCase(suite.repoMock)
}

func (suite *ListAllSchemaUseCaseSuite) TestExecuteWhenSuccess() {
	entitySchemas := []*entity.Schema{
		{
			ID:         "1",
			Service:    "service1",
			Source:     "source1",
			Provider:   "provider",
			SchemaType: "type1",
			JsonSchema: entity.JsonSchema{
				Required: []string{"field1"},
				Properties: map[string]interface{}{
					"field1": map[string]interface{}{
						"type": "string",
					},
					"field2": map[string]interface{}{
						"type": "string",
					},
				},
				JsonType: "object",
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAll").Return(entitySchemas, nil)

	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:         "1",
			Service:    "service1",
			Source:     "source1",
			Provider:   "provider",
			SchemaType: "type1",
			JsonSchema: shareddto.JsonSchemaDTO{
				JsonType: "object",
				Properties: map[string]interface{}{
					"field1": map[string]interface{}{
						"type": "string",
					},
					"field2": map[string]interface{}{
						"type": "string",
					},
				},
				Required: []string{"field1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	output, err := suite.useCase.Execute()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *ListAllSchemaUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAll").Return(nil, fmt.Errorf("error"))

	output, err := suite.useCase.Execute()

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.SchemaDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
