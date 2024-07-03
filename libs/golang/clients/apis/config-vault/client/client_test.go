package client

import (
	"context"
	"encoding/json"
	inputdto "libs/golang/ddd/dtos/config-vault/input"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	shareddto "libs/golang/ddd/dtos/config-vault/shared"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
		case r.URL.Path == "/config" && r.Method == http.MethodPost:
			var configInput inputdto.ConfigDTO
			if err := json.NewDecoder(r.Body).Decode(&configInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			configOutput := outputdto.ConfigDTO{
				ID:              "1",
				Active:          true,
				Service:         "service1",
				Source:          "source1",
				Provider:        "provider1",
				ConfigVersionID: "v1",
				DependsOn: []shareddto.JobDependenciesDTO{
					{Service: "dep_service1", Source: "dep_source1"},
				},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(configOutput)

		case r.URL.Path == "/config" && r.Method == http.MethodPut:
			var configInput inputdto.ConfigDTO
			if err := json.NewDecoder(r.Body).Decode(&configInput); err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			configOutput := outputdto.ConfigDTO{
				ID:              "1",
				Active:          true,
				Service:         "service1",
				Source:          "source1",
				Provider:        "provider1",
				ConfigVersionID: "v1",
				DependsOn: []shareddto.JobDependenciesDTO{
					{Service: "dep_service1", Source: "dep_source1"},
				},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(configOutput)

		case r.URL.Path == "/config" && r.Method == http.MethodGet:
			configList := []outputdto.ConfigDTO{
				{
					ID:              "1",
					Active:          true,
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					ConfigVersionID: "v1",
					DependsOn: []shareddto.JobDependenciesDTO{
						{Service: "dep_service1", Source: "dep_source1"},
					},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(configList)

		case r.URL.Path == "/config/1" && r.Method == http.MethodGet:
			configOutput := outputdto.ConfigDTO{
				ID:              "1",
				Active:          true,
				Service:         "service1",
				Source:          "source1",
				Provider:        "provider1",
				ConfigVersionID: "v1",
				DependsOn: []shareddto.JobDependenciesDTO{
					{Service: "dep_service1", Source: "dep_source1"},
				},
				CreatedAt: "2023-06-01T00:00:00Z",
				UpdatedAt: "2023-06-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(configOutput)

		case r.URL.Path == "/config/1" && r.Method == http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)

		case r.URL.Path == "/config/provider/provider1/service/service1" && r.Method == http.MethodGet:
			configList := []outputdto.ConfigDTO{
				{
					ID:              "1",
					Active:          true,
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					ConfigVersionID: "v1",
					DependsOn: []shareddto.JobDependenciesDTO{
						{Service: "dep_service1", Source: "dep_source1"},
					},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(configList)

		case r.URL.Path == "/config/provider/provider1/source/source1" && r.Method == http.MethodGet:
			configList := []outputdto.ConfigDTO{
				{
					ID:              "1",
					Active:          true,
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					ConfigVersionID: "v1",
					DependsOn: []shareddto.JobDependenciesDTO{
						{Service: "dep_service1", Source: "dep_source1"},
					},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(configList)

		case r.URL.Path == "/config/provider/provider1/service/service1/active/true" && r.Method == http.MethodGet:
			configList := []outputdto.ConfigDTO{
				{
					ID:              "1",
					Active:          true,
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					ConfigVersionID: "v1",
					DependsOn: []shareddto.JobDependenciesDTO{
						{Service: "dep_service1", Source: "dep_source1"},
					},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(configList)

		case r.URL.Path == "/config/provider/provider1/service/service1/source/source1" && r.Method == http.MethodGet:
			configList := []outputdto.ConfigDTO{
				{
					ID:              "1",
					Active:          true,
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					ConfigVersionID: "v1",
					DependsOn: []shareddto.JobDependenciesDTO{
						{Service: "dep_service1", Source: "dep_source1"},
					},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(configList)

		case r.URL.Path == "/config/provider/provider1/dependencies/service/service1/source/source1" && r.Method == http.MethodGet:
			configList := []outputdto.ConfigDTO{
				{
					ID:              "1",
					Active:          true,
					Service:         "service1",
					Source:          "source1",
					Provider:        "provider1",
					ConfigVersionID: "v1",
					DependsOn: []shareddto.JobDependenciesDTO{
						{Service: "dep_service1", Source: "dep_source1"},
					},
					CreatedAt: "2023-06-01T00:00:00Z",
					UpdatedAt: "2023-06-01T00:00:00Z",
				},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(configList)

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

func (suite *ClientTestSuite) TestCreateConfigWhenSuccess() {
	configInput := inputdto.ConfigDTO{
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
		Service:         "service1",
		Source:          "source1",
		Provider:        "provider1",
		ConfigVersionID: "v1",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service1", Source: "dep_source1"},
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	configOutput, err := suite.client.CreateConfig(configInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}

func (suite *ClientTestSuite) TestCreateConfigWhenBadRequest() {
	// Mock server to respond with bad request
	suite.mockServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	})

	configInput := inputdto.ConfigDTO{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
	}

	configOutput, err := suite.client.CreateConfig(configInput)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.ConfigDTO{}, configOutput)
}

func (suite *ClientTestSuite) TestCreateConfigWhenTimeout() {
	// Mock server to delay response
	suite.mockServer.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		http.Error(w, "Timeout", http.StatusGatewayTimeout)
	})

	configInput := inputdto.ConfigDTO{
		Active:   true,
		Service:  "test_service",
		Source:   "test_source",
		Provider: "test_provider",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service", Source: "dep_source"},
		},
	}

	configOutput, err := suite.client.CreateConfig(configInput)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), outputdto.ConfigDTO{}, configOutput)
}

func (suite *ClientTestSuite) TestUpdateConfigWhenSuccess() {
	configInput := inputdto.ConfigDTO{
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
		Service:         "service1",
		Source:          "source1",
		Provider:        "provider1",
		ConfigVersionID: "v1",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service1", Source: "dep_source1"},
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	configOutput, err := suite.client.UpdateConfig(configInput)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}

func (suite *ClientTestSuite) TestListAllConfigsWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	configOutput, err := suite.client.ListAllConfigs()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}

func (suite *ClientTestSuite) TestListConfigByIDWhenSuccess() {
	expectedOutput := outputdto.ConfigDTO{
		ID:              "1",
		Active:          true,
		Service:         "service1",
		Source:          "source1",
		Provider:        "provider1",
		ConfigVersionID: "v1",
		DependsOn: []shareddto.JobDependenciesDTO{
			{Service: "dep_service1", Source: "dep_source1"},
		},
		CreatedAt: "2023-06-01T00:00:00Z",
		UpdatedAt: "2023-06-01T00:00:00Z",
	}

	configOutput, err := suite.client.ListConfigByID("1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}

func (suite *ClientTestSuite) TestDeleteConfigWhenSuccess() {
	err := suite.client.DeleteConfig("1")

	assert.Nil(suite.T(), err)
}

func (suite *ClientTestSuite) TestListConfigsByServiceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	configOutput, err := suite.client.ListConfigsByServiceAndProvider("service1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}

func (suite *ClientTestSuite) TestListConfigsBySourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	configOutput, err := suite.client.ListConfigsBySourceAndProvider("source1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}

func (suite *ClientTestSuite) TestListConfigsByServiceAndProviderAndActiveWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	configOutput, err := suite.client.ListConfigsByServiceAndProviderAndActive("service1", "provider1", "true")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}

func (suite *ClientTestSuite) TestListConfigsByServiceAndSourceAndProviderWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	configOutput, err := suite.client.ListConfigsByServiceAndSourceAndProvider("service1", "source1", "provider1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}

func (suite *ClientTestSuite) TestListConfigsByProviderAndDependenciesWhenSuccess() {
	expectedOutput := []outputdto.ConfigDTO{
		{
			ID:              "1",
			Active:          true,
			Service:         "service1",
			Source:          "source1",
			Provider:        "provider1",
			ConfigVersionID: "v1",
			DependsOn: []shareddto.JobDependenciesDTO{
				{Service: "dep_service1", Source: "dep_source1"},
			},
			CreatedAt: "2023-06-01T00:00:00Z",
			UpdatedAt: "2023-06-01T00:00:00Z",
		},
	}

	configOutput, err := suite.client.ListConfigsByProviderAndDependencies("provider1", "service1", "source1")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedOutput, configOutput)
}
