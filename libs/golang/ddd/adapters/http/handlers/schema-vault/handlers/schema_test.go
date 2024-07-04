package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"libs/golang/ddd/domain/entities/schema-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/schema-vault/repository"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	shareddto "libs/golang/ddd/dtos/schema-vault/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WebSchemaHandlerSuite struct {
	suite.Suite
	handler  *WebSchemaHandler
	repoMock *mockrepository.SchemaRepositoryMock
}

func TestWebSchemaHandlerSuite(t *testing.T) {
	suite.Run(t, new(WebSchemaHandlerSuite))
}

func (suite *WebSchemaHandlerSuite) SetupTest() {
	suite.repoMock = new(mockrepository.SchemaRepositoryMock)
	suite.handler = NewWebSchemaHandler(suite.repoMock)
}

// Tests for CreateSchema handler
func (suite *WebSchemaHandlerSuite) TestCreateSchemaWhenSuccess() {
	inputDTO := inputdto.SchemaDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		SchemaType: "test_schema_type",
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
			Required: []string{
				"field1",
			},
		},
	}

	expectedOutput := outputdto.SchemaDTO{
		ID:              "1",
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		SchemaType:      "test_schema_type",
		SchemaVersionID: "v1",
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
			Required: []string{
				"field1",
			},
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Schema")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Schema)
		arg.ID = "1"
		arg.SchemaVersionID = "v1"
		arg.CreatedAt = "2023-06-01 00:00:00"
		arg.UpdatedAt = "2023-06-01 00:00:00"
	})

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPost, "/schemas", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateSchema(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.SchemaDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebSchemaHandlerSuite) TestCreateSchemaWhenDecodingFails() {
	req := httptest.NewRequest(http.MethodPost, "/schemas", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateSchema(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "invalid character")
}

func (suite *WebSchemaHandlerSuite) TestCreateSchemaWhenRepositoryFails() {
	inputDTO := inputdto.SchemaDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		SchemaType: "test_schema_type",
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
			Required: []string{
				"field1",
			},
		},
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Schema")).Return(errors.New("repository error"))

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPost, "/schemas", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateSchema(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for UpdateSchema handler
func (suite *WebSchemaHandlerSuite) TestUpdateSchemaWhenSuccess() {
	inputDTO := inputdto.SchemaDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		SchemaType: "test_schema_type",
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
			Required: []string{
				"field1",
			},
		},
	}

	expectedOutput := outputdto.SchemaDTO{
		ID:              "1",
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		SchemaType:      "test_schema_type",
		SchemaVersionID: "v1",
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
			Required: []string{
				"field1",
			},
		},
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Schema")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Schema)
		arg.ID = "1"
		arg.SchemaVersionID = "v1"
		arg.CreatedAt = "2023-06-01 00:00:00"
		arg.UpdatedAt = "2023-06-01 00:00:00"
	})

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPut, "/schemas/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateSchema(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.SchemaDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebSchemaHandlerSuite) TestUpdateSchemaWhenDecodingFails() {
	req := httptest.NewRequest(http.MethodPut, "/schemas/1", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateSchema(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "invalid character")
}

func (suite *WebSchemaHandlerSuite) TestUpdateSchemaWhenRepositoryFails() {
	inputDTO := inputdto.SchemaDTO{
		Service:    "test_service",
		Source:     "test_source",
		Provider:   "test_provider",
		SchemaType: "test_schema_type",
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
			Required: []string{
				"field1",
			},
		},
	}

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Schema")).Return(errors.New("repository error"))

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest(http.MethodPut, "/schemas/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateSchema(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for DeleteSchema handler
func (suite *WebSchemaHandlerSuite) TestDeleteSchemaWhenSuccess() {
	suite.repoMock.On("Delete", "1").Return(nil)

	req := httptest.NewRequest("DELETE", "/schemas/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteSchema(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	assert.Equal(suite.T(), "Schema deleted successfully", rr.Body.String())
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebSchemaHandlerSuite) TestDeleteSchemaWhenIDNotProvided() {
	req := httptest.NewRequest("DELETE", "/schema/", nil)
	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteSchema(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "ID is required")
}

func (suite *WebSchemaHandlerSuite) TestDeleteSchemaWhenRepositoryFails() {
	suite.repoMock.On("Delete", "1").Return(errors.New("repository error"))

	req := httptest.NewRequest("DELETE", "/schemas/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteSchema(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListAllSchemas handler
func (suite *WebSchemaHandlerSuite) TestListAllSchemasWhenSuccess() {
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
		{
			ID:         "2",
			Service:    "service2",
			Source:     "source2",
			Provider:   "provider",
			SchemaType: "type2",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
		{
			ID:         "2",
			Service:    "service2",
			Source:     "source2",
			Provider:   "provider",
			SchemaType: "type2",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/schemas", nil)
	rr := httptest.NewRecorder()

	suite.handler.ListAllSchemas(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.SchemaDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebSchemaHandlerSuite) TestListAllSchemasWhenRepositoryFails() {
	suite.repoMock.On("FindAll").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest(http.MethodGet, "/schemas", nil)
	rr := httptest.NewRecorder()

	suite.handler.ListAllSchemas(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListSchemaByID handler
func (suite *WebSchemaHandlerSuite) TestListSchemaByIDWhenSuccess() {
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
		{
			ID:         "2",
			Service:    "service2",
			Source:     "source2",
			Provider:   "provider",
			SchemaType: "type2",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindByID", "1").Return(entitySchemas[0], nil)

	expectedOutput := outputdto.SchemaDTO{
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
		CreatedAt: "2023-06-01 00:00:00",
		UpdatedAt: "2023-06-01 00:00:00",
	}

	req := httptest.NewRequest(http.MethodGet, "/schemas/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemaByID(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.SchemaDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebSchemaHandlerSuite) TestListSchemaByIDWhenIDNotProvided() {
	req := httptest.NewRequest(http.MethodGet, "/schema/", nil)
	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemaByID(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "ID is required")
}

func (suite *WebSchemaHandlerSuite) TestListSchemaByIDWhenRepositoryFails() {
	suite.repoMock.On("FindByID", "1").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest(http.MethodGet, "/schemas/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemaByID(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListSchemasByServiceAndProvider handler
func (suite *WebSchemaHandlerSuite) TestListSchemasByServiceAndProviderWhenSuccess() {
	entitySchemas := []*entity.Schema{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider",
			SchemaType:      "type1",
			SchemaVersionID: "v1",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return(entitySchemas, nil)

	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider",
			SchemaType:      "type1",
			SchemaVersionID: "v1",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	req := httptest.NewRequest("GET", "/schemas/service/test_service/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemasByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.SchemaDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebSchemaHandlerSuite) TestListSchemasByServiceAndProviderWhenServiceOrProviderNotProvided() {
	req := httptest.NewRequest("GET", "/schemas/service/test_service", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemasByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Service and provider are required")
}

func (suite *WebSchemaHandlerSuite) TestListSchemasByServiceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/schemas/service/test_service/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemasByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListSchemasBySourceAndProvider handler
func (suite *WebSchemaHandlerSuite) TestListSchemasBySourceAndProviderWhenSuccess() {
	entitySchemas := []*entity.Schema{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider",
			SchemaType:      "type1",
			SchemaVersionID: "v1",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllBySourceAndProvider", "test_provider", "test_source").Return(entitySchemas, nil)

	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider",
			SchemaType:      "type1",
			SchemaVersionID: "v1",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	req := httptest.NewRequest("GET", "/schemas/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemasBySourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.SchemaDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebSchemaHandlerSuite) TestListSchemasBySourceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllBySourceAndProvider", "test_provider", "test_source").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/schemas/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemasBySourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListSchemasByServiceAndSourceAndProvider handler
func (suite *WebSchemaHandlerSuite) TestListSchemasByServiceAndSourceAndProviderWhenSuccess() {
	entitySchemas := []*entity.Schema{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider",
			SchemaType:      "type1",
			SchemaVersionID: "v1",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "test_service", "test_source", "test_provider").Return(entitySchemas, nil)

	expectedOutput := []outputdto.SchemaDTO{
		{
			ID:              "1",
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider",
			SchemaType:      "type1",
			SchemaVersionID: "v1",
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
			CreatedAt: "2023-06-01 00:00:00",
			UpdatedAt: "2023-06-01 00:00:00",
		},
	}

	req := httptest.NewRequest("GET", "/schemas/service/test_service/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemasByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.SchemaDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebSchemaHandlerSuite) TestListSchemasByServiceAndSourceAndProviderWhenServiceOrSourceOrProviderNotProvided() {
	req := httptest.NewRequest("GET", "/schemas/service/test_service/source/test_source", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemasByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Service, source and provider are required")
}

func (suite *WebSchemaHandlerSuite) TestListSchemasByServiceAndSourceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "test_service", "test_source", "test_provider").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/schemas/service/test_service/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListSchemasByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}
