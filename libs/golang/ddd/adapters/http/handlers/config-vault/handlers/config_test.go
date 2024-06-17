package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"libs/golang/ddd/domain/entities/config-vault/entity"
	mockrepository "libs/golang/ddd/domain/repositories/database/mock/config-vault/repository"
	inputdto "libs/golang/ddd/dtos/config-vault/input"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WebConfigHandlerSuite struct {
	suite.Suite
	handler  *WebConfigHandler
	repoMock *mockrepository.ConfigRepositoryMock
}

func TestWebConfigHandlerSuite(t *testing.T) {
	suite.Run(t, new(WebConfigHandlerSuite))
}

func (suite *WebConfigHandlerSuite) SetupTest() {
	suite.repoMock = new(mockrepository.ConfigRepositoryMock)
	suite.handler = NewWebConfigHandler(suite.repoMock)
}

// Tests for CreateConfig handler
func (suite *WebConfigHandlerSuite) TestCreateConfigWhenSuccess() {
	inputDTO := inputdto.ConfigDTO{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
	}

	expectedOutput := outputdto.ConfigDTO{
		ID:              "1",
		Active:          true,
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		ConfigVersionID: "v1",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Config")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Config)
		arg.ID = "1"
		arg.ConfigVersionID = "v1"
		arg.CreatedAt = "2023-06-01T00:00:00Z"
		arg.UpdatedAt = "2023-06-01T00:00:00Z"
	})

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest("POST", "/configs", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateConfig(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestCreateConfigWhenDecodingFails() {
	req := httptest.NewRequest("POST", "/configs", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateConfig(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "invalid character")
}

func (suite *WebConfigHandlerSuite) TestCreateConfigWhenRepositoryFails() {
	inputDTO := inputdto.ConfigDTO{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
	}

	suite.repoMock.On("Create", mock.AnythingOfType("*entity.Config")).Return(errors.New("repository error"))

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest("POST", "/configs", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.CreateConfig(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for UpdateConfig handler
func (suite *WebConfigHandlerSuite) TestUpdateConfigWhenSuccess() {
	inputDTO := inputdto.ConfigDTO{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
	}

	expectedOutput := outputdto.ConfigDTO{
		ID:              "1",
		Active:          true,
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		ConfigVersionID: "v1",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Config")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entity.Config)
		arg.ID = "1"
		arg.ConfigVersionID = "v1"
		arg.CreatedAt = "2023-06-01T00:00:00Z"
		arg.UpdatedAt = "2023-06-01T00:00:00Z"
	})

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest("PUT", "/configs", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateConfig(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestUpdateConfigWhenDecodingFails() {
	req := httptest.NewRequest("PUT", "/configs", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateConfig(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "invalid character")
}

func (suite *WebConfigHandlerSuite) TestUpdateConfigWhenRepositoryFails() {
	inputDTO := inputdto.ConfigDTO{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
	}

	suite.repoMock.On("Update", mock.AnythingOfType("*entity.Config")).Return(errors.New("repository error"))

	jsonBody, _ := json.Marshal(inputDTO)
	req := httptest.NewRequest("PUT", "/configs", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	suite.handler.UpdateConfig(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for DeleteConfig handler
func (suite *WebConfigHandlerSuite) TestDeleteConfigWhenSuccess() {
	suite.repoMock.On("Delete", "1").Return(nil)

	req := httptest.NewRequest("DELETE", "/configs/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteConfig(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	assert.Equal(suite.T(), "Config deleted successfully", rr.Body.String())
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestDeleteConfigWhenIDNotProvided() {
	req := httptest.NewRequest("DELETE", "/configs", nil)
	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteConfig(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "ID is required")
}

func (suite *WebConfigHandlerSuite) TestDeleteConfigWhenRepositoryFails() {
	suite.repoMock.On("Delete", "1").Return(errors.New("repository error"))

	req := httptest.NewRequest("DELETE", "/configs/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.DeleteConfig(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListAllConfigs handler
func (suite *WebConfigHandlerSuite) TestListAllConfigsWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service1",
			Source:          "test_source1",
			Provider:        "test_provider1",
			ConfigVersionID: "v1",
			DependsOn:       []shareddto.JobDependenciesDTO{},
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
		{
			ID:              "2",
			Active:          true,
			Service:         "test_service2",
			Source:          "test_source2",
			Provider:        "test_provider2",
			ConfigVersionID: "v2",
			DependsOn:       []shareddto.JobDependenciesDTO{},
			CreatedAt:       "2023-06-02T00:00:00Z",
			UpdatedAt:       "2023-06-02T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAll").Return([]*entity.Config{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service1",
			Source:          "test_source1",
			Provider:        "test_provider1",
			ConfigVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
		{
			ID:              "2",
			Active:          true,
			Service:         "test_service2",
			Source:          "test_source2",
			Provider:        "test_provider2",
			ConfigVersionID: "v2",
			CreatedAt:       "2023-06-02T00:00:00Z",
			UpdatedAt:       "2023-06-02T00:00:00Z",
		},
	}, nil)

	req := httptest.NewRequest("GET", "/configs", nil)
	rr := httptest.NewRecorder()

	suite.handler.ListAllConfigs(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestListAllConfigsWhenRepositoryFails() {
	suite.repoMock.On("FindAll").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/configs", nil)
	rr := httptest.NewRecorder()

	suite.handler.ListAllConfigs(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListConfigByID handler
func (suite *WebConfigHandlerSuite) TestListConfigByIDWhenSuccess() {
	expectedOutput := outputdto.ConfigDTO{
		ID:              "1",
		Active:          true,
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		ConfigVersionID: "v1",
		DependsOn:       []shareddto.JobDependenciesDTO{},
		CreatedAt:       "2023-06-01T00:00:00Z",
		UpdatedAt:       "2023-06-01T00:00:00Z",
	}

	suite.repoMock.On("FindByID", "1").Return(&entity.Config{
		ID:              "1",
		Active:          true,
		Service:         "test_service",
		Source:          "test_source",
		Provider:        "test_provider",
		ConfigVersionID: "v1",
		CreatedAt:       "2023-06-01T00:00:00Z",
		UpdatedAt:       "2023-06-01T00:00:00Z",
	}, nil)

	req := httptest.NewRequest("GET", "/config/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigByID(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestListConfigByIDWhenIDNotProvided() {
	req := httptest.NewRequest("GET", "/config/", nil)
	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigByID(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "ID is required")
}

func (suite *WebConfigHandlerSuite) TestListConfigByIDWhenRepositoryFails() {
	suite.repoMock.On("FindByID", "1").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/config/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigByID(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListConfigsByServiceAndProvider handler
func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			DependsOn:       []shareddto.JobDependenciesDTO{},
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return([]*entity.Config{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}, nil)

	req := httptest.NewRequest("GET", "/configs/service/test_service/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndProviderWhenServiceOrProviderNotProvided() {
	req := httptest.NewRequest("GET", "/configs/service/test_service", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Service and provider are required")
}

func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllByServiceAndProvider", "test_provider", "test_service").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/configs/service/test_service/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListConfigsBySourceAndProvider handler
func (suite *WebConfigHandlerSuite) TestListConfigsBySourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			DependsOn:       []shareddto.JobDependenciesDTO{},
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllBySourceAndProvider", "test_provider", "test_source").Return([]*entity.Config{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}, nil)

	req := httptest.NewRequest("GET", "/configs/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsBySourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestListConfigsBySourceAndProviderWhenSourceOrProviderNotProvided() {
	req := httptest.NewRequest("GET", "/configs/source/test_source", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsBySourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Source and provider are required")
}

func (suite *WebConfigHandlerSuite) TestListConfigsBySourceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllBySourceAndProvider", "test_provider", "test_source").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/configs/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsBySourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListConfigsByServiceAndSourceAndProvider handler
func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndSourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			DependsOn:       []shareddto.JobDependenciesDTO{},
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "test_service", "test_source", "test_provider").Return([]*entity.Config{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}, nil)

	req := httptest.NewRequest("GET", "/configs/service/test_service/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndSourceAndProviderWhenServiceSourceOrProviderNotProvided() {
	req := httptest.NewRequest("GET", "/configs/service/test_service/source/test_source", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Service, source, and provider are required")
}

func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndSourceAndProviderWhenRepositoryFails() {
	suite.repoMock.On("FindAllByServiceAndSourceAndProvider", "test_service", "test_source", "test_provider").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/configs/service/test_service/source/test_source/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("source", "test_source")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndSourceAndProvider(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListConfigsByServiceAndProviderAndActive handler
func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndProviderAndActiveWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			DependsOn:       []shareddto.JobDependenciesDTO{},
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllByServiceAndProviderAndActive", "test_service", "test_provider", true).Return([]*entity.Config{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			CreatedAt:       "2023-06-01T00:00:00Z",
			UpdatedAt:       "2023-06-01T00:00:00Z",
		},
	}, nil)

	req := httptest.NewRequest("GET", "/configs/service/test_service/provider/test_provider/active/true", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("active", "true")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndProviderAndActive(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndProviderAndActiveWhenServiceProviderOrActiveNotProvided() {
	req := httptest.NewRequest("GET", "/configs/service/test_service/provider/test_provider", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndProviderAndActive(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Service, provider, and active status are required")
}

func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndProviderAndActiveWhenInvalidActiveStatus() {
	req := httptest.NewRequest("GET", "/configs/service/test_service/provider/test_provider/active/invalid", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("active", "invalid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndProviderAndActive(rr, req)

	assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Invalid active status value")
}

func (suite *WebConfigHandlerSuite) TestListConfigsByServiceAndProviderAndActiveWhenRepositoryFails() {
	suite.repoMock.On("FindAllByServiceAndProviderAndActive", "test_service", "test_provider", true).Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/configs/service/test_service/provider/test_provider/active/true", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("service", "test_service")
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("active", "true")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByServiceAndProviderAndActive(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}

// Tests for ListConfigsByProviderAndDependencies handler
func (suite *WebConfigHandlerSuite) TestListConfigsByProviderAndDependenciesWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service", Source: "dep_source"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	suite.repoMock.On("FindAllByProviderAndDependsOn", "test_provider", "dep_service", "dep_source").Return([]*entity.Config{
		{
			ID:              "1",
			Active:          true,
			Service:         "test_service",
			Source:          "test_source",
			Provider:        "test_provider",
			ConfigVersionID: "v1",
			DependsOn: []entity.JobDependencies{
				{Service: "dep_service", Source: "dep_source"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}, nil)

	req := httptest.NewRequest("GET", "/configs/dependencies/provider/test_provider/service/dep_service/source/dep_source", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("service", "dep_service")
	rctx.URLParams.Add("source", "dep_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByProviderAndDependencies(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)

	var actualOutput []outputdto.ConfigDTO
	err := json.NewDecoder(rr.Body).Decode(&actualOutput)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
	suite.repoMock.AssertExpectations(suite.T())
}

func (suite *WebConfigHandlerSuite) TestListConfigsByDependenciesWhenNoDependenciesProvided() {
	tests := []struct {
		url string
	}{
		{"/configs/dependencies/provider//service//source/"},
		{"/configs/dependencies/provider//service/dep_service/source/"},
		{"/configs/dependencies/provider//service//source/dep_source"},
		{"/configs/dependencies/provider/test_provider/service//source/"},
		{"/configs/dependencies/provider/test_provider/service/dep_service/source/"},
		{"/configs/dependencies/provider/test_provider/service//source/dep_source"},
	}

	for _, tc := range tests {
		req := httptest.NewRequest("GET", tc.url, nil)
		rctx := chi.NewRouteContext()
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		rr := httptest.NewRecorder()

		suite.handler.ListConfigsByProviderAndDependencies(rr, req)

		assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
		assert.Contains(suite.T(), rr.Body.String(), "At least one dependency is required")
	}
}

func (suite *WebConfigHandlerSuite) TestListConfigsByProviderAndDependenciesWhenRepositoryFails() {
	suite.repoMock.On("FindAllByProviderAndDependsOn", "test_provider", "dep_service", "dep_source").Return(nil, errors.New("repository error"))

	req := httptest.NewRequest("GET", "/configs/dependencies/provider/test_provider/service/dep_service/source/dep_source", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("provider", "test_provider")
	rctx.URLParams.Add("service", "dep_service")
	rctx.URLParams.Add("source", "dep_source")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	suite.handler.ListConfigsByProviderAndDependencies(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "repository error")
	suite.repoMock.AssertExpectations(suite.T())
}
