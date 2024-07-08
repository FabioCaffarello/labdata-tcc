package usecase

import (
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/schema-vault/repository"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCaseSuite struct {
	suite.Suite
	repoMock *mockrepository.SchemaRepositoryMock
	useCase  *ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase
}

func TestListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ListOneByIDSchemaUseCaseSuite))
}

func (suite *ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCaseSuite) SetupTest() {
	suite.repoMock = new(mockrepository.SchemaRepositoryMock)
	suite.useCase = NewListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCase(suite.repoMock)
}

func (suite *ListOneByServiceAndSourceAndProviderAndSchemaTypeSchemaUseCaseSuite) TestExecuteWhenSuccess() {
	entitySchemas := []*entity.Schema{
		{
			ID:         "1",
			Service:    "service1",
			Source:     "source1",
			Provider:   "provider",
			SchemaType: "input",
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
		{
			ID:         "2",
			Service:    "service1",
			Source:     "source1",
			Provider:   "provider",
			SchemaType: "output",
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

	suite.repoMock.On("FindOneByServiceAndSourceAndProviderAndSchemaType", "provider", "service1", "source1", "input").Return(entitySchemas[0], nil)

	expectedOutput := outputdto.SchemaDTO{
		ID:         "1",
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider",
		SchemaType: "input",
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
	}

	output, err := suite.useCase.Execute("provider", "service1", "source1", "input")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, output)
	suite.repoMock.AssertExpectations(suite.T())
}
