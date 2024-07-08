package client

import (
	"context"
	"encoding/json"
	inputdto "libs/golang/ddd/dtos/input-broker/input"
	outputdto "libs/golang/ddd/dtos/input-broker/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	client     *Client
	mockServer *httptest.Server
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

func (suite *ClientSuite) SetupTest() {
	// Create a mock HTTP server
	suite.mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/input" && r.Method == http.MethodPost:
			var inputInput inputdto.InputDTO
			if err := json.NewDecoder(r.Body).Decode(&inputInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			inputOutput := outputdto.InputDTO{
				ID:        "1",
				Data:      inputInput.Data,
				Metadata:  shareddto.MetadataDTO{Provider: inputInput.Provider, Service: inputInput.Service, Source: inputInput.Source},
				Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(inputOutput)

		case r.URL.Path == "/input/1" && r.Method == http.MethodPut:
			var inputInput inputdto.InputDTO
			if err := json.NewDecoder(r.Body).Decode(&inputInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			inputOutput := outputdto.InputDTO{
				ID:        "1",
				Data:      inputInput.Data,
				Metadata:  shareddto.MetadataDTO{Provider: inputInput.Provider, Service: inputInput.Service, Source: inputInput.Source},
				Status:    shareddto.StatusDTO{Code: 200, Detail: "Updated"},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputOutput)

		case r.URL.Path == "/input/1" && r.Method == http.MethodDelete:
			w.WriteHeader(http.StatusOK)

		case r.URL.Path == "/input" && r.Method == http.MethodGet:
			inputs := []outputdto.InputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
					Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputs)

		case r.URL.Path == "/input/1" && r.Method == http.MethodGet:
			inputOutput := outputdto.InputDTO{
				ID:        "1",
				Data:      map[string]interface{}{"key": "value"},
				Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
				Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputOutput)

		case r.URL.Path == "/input/1/status" && r.Method == http.MethodPut:
			var statusDTO shareddto.StatusDTO
			if err := json.NewDecoder(r.Body).Decode(&statusDTO); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			inputOutput := outputdto.InputDTO{
				ID:        "1",
				Data:      map[string]interface{}{"key": "value"},
				Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
				Status:    shareddto.StatusDTO{Code: statusDTO.Code, Detail: statusDTO.Detail},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputOutput)

		case r.URL.Path == "/input/provider/test_provider/service/test_service" && r.Method == http.MethodGet:
			inputs := []outputdto.InputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
					Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputs)

		case r.URL.Path == "/input/provider/test_provider/source/test_source" && r.Method == http.MethodGet:
			inputs := []outputdto.InputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
					Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputs)

		case r.URL.Path == "/input/provider/test_provider/service/test_service/source/test_source" && r.Method == http.MethodGet:
			inputs := []outputdto.InputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
					Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputs)

		case r.URL.Path == "/input/provider/test_provider/status/200" && r.Method == http.MethodGet:
			inputs := []outputdto.InputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
					Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputs)

		case r.URL.Path == "/input/provider/test_provider/service/test_service/status/200" && r.Method == http.MethodGet:
			inputs := []outputdto.InputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
					Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputs)

		case r.URL.Path == "/input/provider/test_provider/source/test_source/status/200" && r.Method == http.MethodGet:
			inputs := []outputdto.InputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
					Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputs)

		case r.URL.Path == "/input/provider/test_provider/service/test_service/source/test_source/status/200" && r.Method == http.MethodGet:
			inputs := []outputdto.InputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
					Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(inputs)

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

func (suite *ClientSuite) TearDownTest() {
	suite.mockServer.Close()
}

func (suite *ClientSuite) TestCreateInputWhenSuccess() {
	inputInput := inputdto.InputDTO{
		Provider: "test_provider",
		Service:  "test_service",
		Source:   "test_source",
		Data:     map[string]interface{}{"key": "value"},
	}

	expectedOutput := outputdto.InputDTO{
		ID:        "1",
		Data:      map[string]interface{}{"key": "value"},
		Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
		Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	inputOutput, err := suite.client.CreateInput(inputInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputOutput)
}

func (suite *ClientSuite) TestUpdateInputWhenSuccess() {
	inputInput := inputdto.InputDTO{
		Provider: "test_provider",
		Service:  "test_service",
		Source:   "test_source",
		Data:     map[string]interface{}{"key": "value"},
	}

	expectedOutput := outputdto.InputDTO{
		ID:        "1",
		Data:      map[string]interface{}{"key": "value"},
		Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
		Status:    shareddto.StatusDTO{Code: 200, Detail: "Updated"},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	inputOutput, err := suite.client.UpdateInput("1", inputInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputOutput)
}

func (suite *ClientSuite) TestDeleteInputWhenSuccess() {
	err := suite.client.DeleteInput("1")

	assert.Nil(suite.T(), err)
}

func (suite *ClientSuite) TestListAllInputsWhenSuccess() {
	expectedOutput := []outputdto.InputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
			Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	inputs, err := suite.client.ListAllInputs()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputs)
}

func (suite *ClientSuite) TestGetInputByIDWhenSuccess() {
	expectedOutput := outputdto.InputDTO{
		ID:        "1",
		Data:      map[string]interface{}{"key": "value"},
		Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
		Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	inputOutput, err := suite.client.GetInputByID("1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputOutput)
}

func (suite *ClientSuite) TestUpdateInputStatusWhenSuccess() {
	statusDTO := shareddto.StatusDTO{
		Code:   1,
		Detail: "Updated Status",
	}

	expectedOutput := outputdto.InputDTO{
		ID:        "1",
		Data:      map[string]interface{}{"key": "value"},
		Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
		Status:    shareddto.StatusDTO{Code: 1, Detail: "Updated Status"},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	inputOutput, err := suite.client.UpdateInputStatus("1", statusDTO)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputOutput)
}

func (suite *ClientSuite) TestListInputsByServiceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.InputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
			Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	inputs, err := suite.client.ListInputsByServiceAndProvider("test_service", "test_provider")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputs)
}

func (suite *ClientSuite) TestListInputsBySourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.InputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
			Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	inputs, err := suite.client.ListInputsBySourceAndProvider("test_source", "test_provider")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputs)
}

func (suite *ClientSuite) TestListInputsByServiceAndSourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.InputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
			Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	inputs, err := suite.client.ListInputsByServiceAndSourceAndProvider("test_service", "test_source", "test_provider")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputs)
}

func (suite *ClientSuite) TestListInputsByStatusAndProviderWhenSuccess() {
	expectedOutput := []outputdto.InputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
			Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	inputs, err := suite.client.ListInputsByStatusAndProvider(200, "test_provider")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputs)
}

func (suite *ClientSuite) TestListInputsByStatusAndServiceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.InputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
			Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	inputs, err := suite.client.ListInputsByStatusAndServiceAndProvider(200, "test_service", "test_provider")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputs)
}

func (suite *ClientSuite) TestListInputsByStatusAndSourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.InputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
			Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	inputs, err := suite.client.ListInputsByStatusAndSourceAndProvider(200, "test_source", "test_provider")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputs)
}

func (suite *ClientSuite) TestListInputsByStatusAndServiceAndSourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.InputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Metadata:  shareddto.MetadataDTO{Provider: "test_provider", Service: "test_service", Source: "test_source"},
			Status:    shareddto.StatusDTO{Code: 200, Detail: "Success"},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	inputs, err := suite.client.ListInputsByStatusAndServiceAndSourceAndProvider(200, "test_service", "test_source", "test_provider")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, inputs)
}
