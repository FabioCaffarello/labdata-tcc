package client

import (
	"context"
	"encoding/json"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	client     *Client
	mockServer *httptest.Server
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) SetupTest() {
	// Create a mock HTTP server
	suite.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/schema" && r.Method == http.MethodPost:
			var schemaInput inputdto.SchemaDTO
			if err := json.NewDecoder(r.Body).Decode(&schemaInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			schemaOutput := outputdto.SchemaDTO{
				ID:              "1",
				Service:         schemaInput.Service,
				Source:          schemaInput.Source,
				Provider:        schemaInput.Provider,
				SchemaType:      schemaInput.SchemaType,
				JsonSchema:      schemaInput.JsonSchema,
				SchemaVersionID: "v1",
				CreatedAt:       "2023-06-01T00:00:00Z",
				UpdatedAt:       "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(schemaOutput)

		case r.URL.Path == "/schema" && r.Method == http.MethodPut:
			var schemaInput inputdto.SchemaDTO
			if err := json.NewDecoder(r.Body).Decode(&schemaInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			schemaOutput := outputdto.SchemaDTO{
				ID:              "1",
				Service:         schemaInput.Service,
				Source:          schemaInput.Source,
				Provider:        schemaInput.Provider,
				SchemaType:      schemaInput.SchemaType,
				JsonSchema:      schemaInput.JsonSchema,
				SchemaVersionID: "v1",
				CreatedAt:       "2023-06-01T00:00:00Z",
				UpdatedAt:       "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(schemaOutput)

		case r.URL.Path == "/schema" && r.Method == http.MethodGet:
			schemaList := []outputdto.SchemaDTO{
				{
					ID:              "1",
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					SchemaType:      "type1",
					JsonSchema:      shareddto.JsonSchemaDTO{},
					SchemaVersionID: "v1",
					CreatedAt:       "2023-06-01T00:00:00Z",
					UpdatedAt:       "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(schemaList)

		case r.URL.Path == "/schema/1" && r.Method == http.MethodGet:
			schemaOutput := outputdto.SchemaDTO{
				ID:              "1",
				Service:         "service1",
				Source:          "source1",
				Provider:        "provider1",
				SchemaType:      "type1",
				JsonSchema:      shareddto.JsonSchemaDTO{},
				SchemaVersionID: "v1",
				CreatedAt:       "2023-06-01T00:00:00Z",
				UpdatedAt:       "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(schemaOutput)

		case r.URL.Path == "/schema/1" && r.Method == http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)

		case r.URL.Path == "/schema/provider/provider1/service/service1" && r.Method == http.MethodGet:
			schemaList := []outputdto.SchemaDTO{
				{
					ID:              "1",
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					SchemaType:      "type1",
					JsonSchema:      shareddto.JsonSchemaDTO{},
					SchemaVersionID: "v1",
					CreatedAt:       "2023-06-01T00:00:00Z",
					UpdatedAt:       "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(schemaList)

		case r.URL.Path == "/schema/provider/provider1/service/service1/source/source1" && r.Method == http.MethodGet:
			schemaList := []outputdto.SchemaDTO{
				{
					ID:              "1",
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					SchemaType:      "type1",
					JsonSchema:      shareddto.JsonSchemaDTO{},
					SchemaVersionID: "v1",
					CreatedAt:       "2023-06-01T00:00:00Z",
					UpdatedAt:       "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(schemaList)

		case r.URL.Path == "/schema/provider/provider1/source/source1" && r.Method == http.MethodGet:
			schemaList := []outputdto.SchemaDTO{
				{
					ID:              "1",
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					SchemaType:      "type1",
					JsonSchema:      shareddto.JsonSchemaDTO{},
					SchemaVersionID: "v1",
					CreatedAt:       "2023-06-01T00:00:00Z",
					UpdatedAt:       "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(schemaList)

		case r.URL.Path == "/schema/provider/provider1/service/service1/source/source1/schema-type/input" && r.Method == http.MethodGet:
			schema := outputdto.SchemaDTO{
				ID:         "1",
				Service:    "service1",
				Source:     "source1",
				Provider:   "provider1",
				SchemaType: "input",
				JsonSchema: shareddto.JsonSchemaDTO{
					Required: []string{"field1"},
					Properties: map[string]interface{}{
						"field1": map[string]interface{}{
							"type": "string",
						},
					},
					JsonType: "object",
				},
				SchemaVersionID: "v1",
				CreatedAt:       "2023-06-01T00:00:00Z",
				UpdatedAt:       "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(schema)

		case r.URL.Path == "/schema/validate" && r.Method == http.MethodPost:
			w.WriteHeader(http.StatusOK)

		default:
			http.NotFound(w, r)
		}
	}))

	// Initialize the client with the mock server's URL
	suite.client = &Client{
		ctx:     context.Background(),
		baseURL: suite.mockServer.URL,
		timeout: apiTimeout,
	}
}

func (suite *ClientTestSuite) TearDownTest() {
	suite.mockServer.Close()
}

func (suite *ClientTestSuite) TestCreateSchemaWhenSuccess() {
	schemaInput := inputdto.SchemaDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		SchemaType: "test_type",
		JsonSchema: shareddto.JsonSchemaDTO{},
	}

	expectedOutput := outputdto.SchemaDTO{
		ID:              "1",
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		SchemaType:      "test_type",
		JsonSchema:      shareddto.JsonSchemaDTO{},
		SchemaVersionID: "v1",
		CreatedAt:       "2023-06-01T00:00:00Z",
		UpdatedAt:       "2023-06-01T00:00:00Z",
	}

	schemaOutput, err := suite.client.CreateSchema(schemaInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, schemaOutput)
}

func (suite *ClientTestSuite) TestUpdateSchemaWhenSuccess() {
	schemaInput := inputdto.SchemaDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		SchemaType: "test_type",
		JsonSchema: shareddto.JsonSchemaDTO{},
	}

	expectedOutput := outputdto.SchemaDTO{
		ID:              "1",
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		SchemaType:      "test_type",
		JsonSchema:      shareddto.JsonSchemaDTO{},
		SchemaVersionID: "v1",
		CreatedAt:       "2023-06-01T00:00:00Z",
		UpdatedAt:       "2023-06-01T00:00:00Z",
	}

	schemaOutput, err := suite.client.UpdateSchema(schemaInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, schemaOutput)
}

func (suite *ClientTestSuite) TestListAllSchemasWhenSuccess() {
	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			SchemaType:      "type1",
			JsonSchema:      shareddto.JsonSchemaDTO{},
			SchemaVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}

	schemaOutput, err := suite.client.ListAllSchemas()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, schemaOutput)
}

func (suite *ClientTestSuite) TestListSchemaByIDWhenSuccess() {
	expectedOutput := outputdto.SchemaDTO{
		ID:              "1",
		Service:         "service1",
		Source:          "source1",
		Provider:        "provider1",
		SchemaType:      "type1",
		JsonSchema:      shareddto.JsonSchemaDTO{},
		SchemaVersionID: "v1",
		CreatedAt:       "2023-06-01T00:00:00Z",
		UpdatedAt:       "2023-06-01T00:00:00Z",
	}

	schemaOutput, err := suite.client.ListSchemaByID("1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, schemaOutput)
}

func (suite *ClientTestSuite) TestDeleteSchemaWhenSuccess() {
	err := suite.client.DeleteSchema("1")

	assert.Nil(suite.T(), err)
}

func (suite *ClientTestSuite) TestListSchemasByServiceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			SchemaType:      "type1",
			JsonSchema:      shareddto.JsonSchemaDTO{},
			SchemaVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}

	schemaOutput, err := suite.client.ListSchemasByServiceAndProvider("service1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, schemaOutput)
}

func (suite *ClientTestSuite) TestListSchemasBySourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			SchemaType:      "type1",
			JsonSchema:      shareddto.JsonSchemaDTO{},
			SchemaVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}

	schemaOutput, err := suite.client.ListSchemasBySourceAndProvider("source1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, schemaOutput)
}

func (suite *ClientTestSuite) TestListSchemasByServiceAndSourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			SchemaType:      "type1",
			JsonSchema:      shareddto.JsonSchemaDTO{},
			SchemaVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}

	schemaOutput, err := suite.client.ListSchemasByServiceAndSourceAndProvider("service1", "source1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, schemaOutput)
}

func (suite *ClientTestSuite) TestListSchemaByServiceAndSourceAndProviderAndSchemaTypeWhenSuccess() {
	expectedOutput := outputdto.SchemaDTO{
		ID:         "1",
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider1",
		SchemaType: "input",
		JsonSchema: shareddto.JsonSchemaDTO{
			Required: []string{"field1"},
			Properties: map[string]interface{}{
				"field1": map[string]interface{}{
					"type": "string",
				},
			},
			JsonType: "object",
		},
		SchemaVersionID: "v1",
		CreatedAt:       "2023-06-01T00:00:00Z",
		UpdatedAt:       "2023-06-01T00:00:00Z",
	}

	schemaOutput, err := suite.client.ListSchemaByServiceAndSourceAndProviderAndSchemaType("provider1", "service1", "source1", "input")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, schemaOutput)
}

func (suite *ClientTestSuite) TestValidateSchemaWhenSuccess() {
	schemaData := inputdto.SchemaDataDTO{
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider1",
		SchemaType: "input",
		Data: map[string]interface{}{
			"field1": "value1",
		},
	}

	err := suite.client.ValidateSchema(schemaData)

	assert.Nil(suite.T(), err)
}

func (suite *ClientTestSuite) TestValidateSchemaWhenBadRequest() {
	suite.mockServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/schema/validate" && r.Method == http.MethodPost {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		} else {
			http.NotFound(w, r)
		}
	})

	schemaData := inputdto.SchemaDataDTO{
		Service:    "service1",
		Source:     "source1",
		Provider:   "provider1",
		SchemaType: "input",
		Data: map[string]interface{}{
			"field1": "value1",
		},
	}

	err := suite.client.ValidateSchema(schemaData)

	assert.NotNil(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "Bad Request")
}

func (suite *ClientTestSuite) TestListSchemaByServiceAndSourceAndProviderAndSchemaTypeWhenNotFound() {
	suite.mockServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/schema/provider/provider1/service/service1/source/source1/schema-type/input" && r.Method == http.MethodGet {
			http.NotFound(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	schemaOutput, err := suite.client.ListSchemaByServiceAndSourceAndProviderAndSchemaType("provider1", "service1", "source1", "input")

	assert.NotNil(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "404")
	assert.Equal(suite.T(), outputdto.SchemaDTO{}, schemaOutput)
}
