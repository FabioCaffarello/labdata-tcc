package client

import (
	"context"
	inputdto "libs/golang/ddd/dtos/config-vault/input"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
	"libs/golang/shared/go-request/requests"
	"net/http"
	"time"
)

var (
	apiTimeout     = 100 * time.Millisecond
	defaultHeaders = map[string]string{"Content-Type": "application/json"}
)

// Client represents the configuration vault client.
type Client struct {
	ctx     context.Context
	baseURL string
	timeout time.Duration
}

// NewClient initializes a new configuration vault client.
func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://config-handler:8000",
		timeout: apiTimeout,
	}
}

// CreateConfig sends a request to create a new configuration.
//
// Parameters:
//   - configInput: The configuration data transfer object.
//
// Returns:
//   - outputdto.ConfigDTO: The created configuration data transfer object.
//   - error: An error if the request fails.
func (c *Client) CreateConfig(configInput inputdto.ConfigDTO) (outputdto.ConfigDTO, error) {
	pathParams := []string{"config"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, configInput, defaultHeaders, http.MethodPost)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	var configOutput outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configOutput, c.timeout)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	return configOutput, nil
}

// UpdateConfig sends a request to update an existing configuration.
//
// Parameters:
//   - configInput: The configuration data transfer object.
//
// Returns:
//   - outputdto.ConfigDTO: The updated configuration data transfer object.
//   - error: An error if the request fails.
func (c *Client) UpdateConfig(configInput inputdto.ConfigDTO) (outputdto.ConfigDTO, error) {
	pathParams := []string{"config"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, configInput, defaultHeaders, http.MethodPut)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	var configOutput outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configOutput, c.timeout)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	return configOutput, nil
}

// ListAllConfigs sends a request to retrieve all configurations.
//
// Returns:
//   - []outputdto.ConfigDTO: A slice of configuration data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListAllConfigs() ([]outputdto.ConfigDTO, error) {
	pathParams := []string{"config"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var configList []outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configList, c.timeout)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

// ListConfigByID sends a request to retrieve a configuration by its ID.
//
// Parameters:
//   - id: The ID of the configuration.
//
// Returns:
//   - outputdto.ConfigDTO: The configuration data transfer object.
//   - error: An error if the request fails.
func (c *Client) ListConfigByID(id string) (outputdto.ConfigDTO, error) {
	pathParams := []string{"config", id}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	var configOutput outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configOutput, c.timeout)
	if err != nil {
		return outputdto.ConfigDTO{}, err
	}

	return configOutput, nil
}

// DeleteConfig sends a request to delete a configuration by its ID.
//
// Parameters:
//   - id: The ID of the configuration.
//
// Returns:
//   - error: An error if the request fails.
func (c *Client) DeleteConfig(id string) error {
	pathParams := []string{"config", id}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodDelete)
	if err != nil {
		return err
	}

	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, nil, c.timeout)
	if err != nil {
		return err
	}

	return nil
}

// ListConfigsByServiceAndProvider sends a request to retrieve configurations by service and provider.
//
// Parameters:
//   - service: The service name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.ConfigDTO: A slice of configuration data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListConfigsByServiceAndProvider(service, provider string) ([]outputdto.ConfigDTO, error) {
	pathParams := []string{"config", "provider", provider, "service", service}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var configList []outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configList, c.timeout)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

// ListConfigsBySourceAndProvider sends a request to retrieve configurations by source and provider.
//
// Parameters:
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.ConfigDTO: A slice of configuration data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListConfigsBySourceAndProvider(source, provider string) ([]outputdto.ConfigDTO, error) {
	pathParams := []string{"config", "provider", provider, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var configList []outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configList, c.timeout)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

// ListConfigsByServiceAndProviderAndActive sends a request to retrieve configurations by service, provider, and active status.
//
// Parameters:
//   - service: The service name.
//   - provider: The provider name.
//   - active: The active status.
//
// Returns:
//   - []outputdto.ConfigDTO: A slice of configuration data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListConfigsByServiceAndProviderAndActive(service, provider, active string) ([]outputdto.ConfigDTO, error) {
	pathParams := []string{"config", "provider", provider, "service", service, "active", active}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var configList []outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configList, c.timeout)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

// ListConfigsByServiceAndSourceAndProvider sends a request to retrieve configurations by service, source, and provider.
//
// Parameters:
//   - service: The service name.
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.ConfigDTO: A slice of configuration data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListConfigsByServiceAndSourceAndProvider(service, source, provider string) ([]outputdto.ConfigDTO, error) {
	pathParams := []string{"config", "provider", provider, "service", service, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var configList []outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configList, c.timeout)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

// ListConfigsByProviderAndDependencies sends a request to retrieve configurations by provider and dependencies.
//
// Parameters:
//   - provider: The provider name.
//   - service: The service name.
//   - source: The source name.
//
// Returns:
//   - []outputdto.ConfigDTO: A slice of configuration data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListConfigsByProviderAndDependencies(provider, service, source string) ([]outputdto.ConfigDTO, error) {
	pathParams := []string{"config", "provider", provider, "dependencies", "service", service, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var configList []outputdto.ConfigDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &configList, c.timeout)
	if err != nil {
		return nil, err
	}

	return configList, nil
}
