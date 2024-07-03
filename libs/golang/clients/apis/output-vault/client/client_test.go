package client

import (
	"context"
	"encoding/json"
	inputdto "libs/golang/ddd/dtos/output-vault/input"
	outputdto "libs/golang/ddd/dtos/output-vault/output"
	shareddto "libs/golang/ddd/dtos/output-vault/shared"
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
		case r.URL.Path == "/output" && r.Method == http.MethodPost:
			var outputInput inputdto.OutputDTO
			if err := json.NewDecoder(r.Body).Decode(&outputInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			outputOutput := outputdto.OutputDTO{
				ID:        "1",
				Data:      outputInput.Data,
				Service:   outputInput.Service,
				Source:    outputInput.Source,
				Provider:  outputInput.Provider,
				Metadata:  outputInput.Metadata,
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(outputOutput)

		case r.URL.Path == "/output" && r.Method == http.MethodPut:
			var outputInput inputdto.OutputDTO
			if err := json.NewDecoder(r.Body).Decode(&outputInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			outputOutput := outputdto.OutputDTO{
				ID:        "1",
				Data:      outputInput.Data,
				Service:   outputInput.Service,
				Source:    outputInput.Source,
				Provider:  outputInput.Provider,
				Metadata:  outputInput.Metadata,
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(outputOutput)

		case r.URL.Path == "/output" && r.Method == http.MethodGet:
			outputList := []outputdto.OutputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Service:   "service1",
					Source:    "source1",
					Provider:  "provider1",
					Metadata:  shareddto.MetadataDTO{},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(outputList)

		case r.URL.Path == "/output/1" && r.Method == http.MethodGet:
			outputOutput := outputdto.OutputDTO{
				ID:        "1",
				Data:      map[string]interface{}{"key": "value"},
				Service:   "service1",
				Source:    "source1",
				Provider:  "provider1",
				Metadata:  shareddto.MetadataDTO{},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(outputOutput)

		case r.URL.Path == "/output/1" && r.Method == http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)

		case r.URL.Path == "/output/provider/provider1/service/service1" && r.Method == http.MethodGet:
			outputList := []outputdto.OutputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Service:   "service1",
					Source:    "source1",
					Provider:  "provider1",
					Metadata:  shareddto.MetadataDTO{},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(outputList)

		case r.URL.Path == "/output/provider/provider1/source/source1" && r.Method == http.MethodGet:
			outputList := []outputdto.OutputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Service:   "service1",
					Source:    "source1",
					Provider:  "provider1",
					Metadata:  shareddto.MetadataDTO{},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(outputList)

		case r.URL.Path == "/output/provider/provider1/service/service1/source/source1" && r.Method == http.MethodGet:
			outputList := []outputdto.OutputDTO{
				{
					ID:        "1",
					Data:      map[string]interface{}{"key": "value"},
					Service:   "service1",
					Source:    "source1",
					Provider:  "provider1",
					Metadata:  shareddto.MetadataDTO{},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(outputList)

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

func (suite *ClientTestSuite) TestCreateOutputWhenSuccess() {
	outputInput := inputdto.OutputDTO{
		Data:     map[string]interface{}{"key": "value"},
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Metadata: shareddto.MetadataDTO{},
	}

	expectedOutput := outputdto.OutputDTO{
		ID:        "1",
		Data:      map[string]interface{}{"key": "value"},
		Service:   "test_service",
		Source:    "test_source",
		Provider:  "test_provider",
		Metadata:  shareddto.MetadataDTO{},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	outputOutput, err := suite.client.CreateOutput(outputInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, outputOutput)
}

func (suite *ClientTestSuite) TestUpdateOutputWhenSuccess() {
	outputInput := inputdto.OutputDTO{
		Data:     map[string]interface{}{"key": "value"},
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		Metadata: shareddto.MetadataDTO{},
	}

	expectedOutput := outputdto.OutputDTO{
		ID:        "1",
		Data:      map[string]interface{}{"key": "value"},
		Service:   "test_service",
		Source:    "test_source",
		Provider:  "test_provider",
		Metadata:  shareddto.MetadataDTO{},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	outputOutput, err := suite.client.UpdateOutput(outputInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, outputOutput)
}

func (suite *ClientTestSuite) TestListAllOutputsWhenSuccess() {
	expectedOutput := []outputdto.OutputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Service:   "service1",
			Source:    "source1",
			Provider:  "provider1",
			Metadata:  shareddto.MetadataDTO{},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	outputOutput, err := suite.client.ListAllOutputs()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, outputOutput)
}

func (suite *ClientTestSuite) TestListOutputByIDWhenSuccess() {
	expectedOutput := outputdto.OutputDTO{
		ID:        "1",
		Data:      map[string]interface{}{"key": "value"},
		Service:   "service1",
		Source:    "source1",
		Provider:  "provider1",
		Metadata:  shareddto.MetadataDTO{},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	outputOutput, err := suite.client.ListOutputByID("1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, outputOutput)
}

func (suite *ClientTestSuite) TestDeleteOutputWhenSuccess() {
	err := suite.client.DeleteOutput("1")

	assert.Nil(suite.T(), err)
}

func (suite *ClientTestSuite) TestListOutputsByServiceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.OutputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Service:   "service1",
			Source:    "source1",
			Provider:  "provider1",
			Metadata:  shareddto.MetadataDTO{},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	outputOutput, err := suite.client.ListOutputsByServiceAndProvider("service1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, outputOutput)
}

func (suite *ClientTestSuite) TestListOutputsBySourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.OutputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Service:   "service1",
			Source:    "source1",
			Provider:  "provider1",
			Metadata:  shareddto.MetadataDTO{},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	outputOutput, err := suite.client.ListOutputsBySourceAndProvider("source1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, outputOutput)
}

func (suite *ClientTestSuite) TestListOutputsByServiceAndSourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.OutputDTO{
		{
			ID:        "1",
			Data:      map[string]interface{}{"key": "value"},
			Service:   "service1",
			Source:    "source1",
			Provider:  "provider1",
			Metadata:  shareddto.MetadataDTO{},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	outputOutput, err := suite.client.ListOutputsByServiceAndSourceAndProvider("service1", "source1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, outputOutput)
}
