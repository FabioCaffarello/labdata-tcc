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

type ListAListAllBySourceAndProviderSchemaUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.SchemaRepositoryMock
	useCase  *ListAllBySourceAndProviderSchemaUseCase
}

func TestListAllBySourceAndProviderSchemaUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListAListAllBySourceAndProviderSchemaUseCaseSuite))
}

func (suite *ListAListAllBySourceAndProviderSchemaUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.SchemaRepositoryMock)
	suite.useCase = NewListAllBySourceAndProviderSchemaUseCase(suite.repoMock)
}

func (suite *ListAListAllBySourceAndProviderSchemaUseCaseSuite) TestExecuteWhenSuccess() {
	entitySchemas := []*entity.Schema{
		{
			ID:         "1",
			Service:    "service1",
			Source:     "source1",
			Provider:   "provider1",
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

	suite.repoMock.On("FindAllBySourceAndProvider", "provider1", "source1").Return(entitySchemas, nil)

	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:         "1",
			Service:    "service1",
			Source:     "source1",
			Provider:   "provider1",
			SchemaType: "type1",
			JsonSchema: shareddto.JsonSchemaDTO{
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

	output, err := suite.useCase.Execute("provider1", "source1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
}

func (suite *ListAListAllBySourceAndProviderSchemaUseCaseSuite) TestExecuteWhenError() {
	suite.repoMock.On("FindAllBySourceAndProvider", "provider1", "source1").Return(nil, fmt.Errorf("error"))

	output, err := suite.useCase.Execute("provider1", "source1")

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), []outputdto.SchemaDTO{}, output)
	suite.repoMock.AssertExpectations(suite.T())
}
