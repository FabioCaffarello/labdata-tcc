package client

import (
	"context"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
	outputdto "libs/golang/ddd/dtos/schema-vault/output"
	"libs/golang/shared/go-request/requests"
	"net/http"
	"time"
)

var (
	apiTimeout     = 100 * time.Millisecond
	defaultHeaders = map[string]string{"Content-Type": "application/json"}
)

// Client represents the schema vault client.
type Client struct {
	ctx     context.Context
	baseURL string
	timeout time.Duration
}

// NewClient initializes a new schema vault client.
func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://schema-handler:8000",
		timeout: apiTimeout,
	}
}

// CreateSchema sends a request to create a new schema.
//
// Parameters:
//   - schemaInput: The schema data transfer object.
//
// Returns:
//   - outputdto.SchemaDTO: The created schema data transfer object.
//   - error: An error if the request fails.
func (c *Client) CreateSchema(schemaInput inputdto.SchemaDTO) (outputdto.SchemaDTO, error) {
	pathParams := []string{"schema"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, schemaInput, defaultHeaders, http.MethodPost)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	var schemaOutput outputdto.SchemaDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &schemaOutput, c.timeout)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	return schemaOutput, nil
}

// UpdateSchema sends a request to update an existing schema.
//
// Parameters:
//   - schemaInput: The schema data transfer object.
//
// Returns:
//   - outputdto.SchemaDTO: The updated schema data transfer object.
//   - error: An error if the request fails.
func (c *Client) UpdateSchema(schemaInput inputdto.SchemaDTO) (outputdto.SchemaDTO, error) {
	pathParams := []string{"schema"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, schemaInput, defaultHeaders, http.MethodPut)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	var schemaOutput outputdto.SchemaDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &schemaOutput, c.timeout)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	return schemaOutput, nil
}

// ListAllSchemas sends a request to retrieve all schemas.
//
// Returns:
//   - []outputdto.SchemaDTO: A slice of schema data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListAllSchemas() ([]outputdto.SchemaDTO, error) {
	pathParams := []string{"schema"}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var schemaList []outputdto.SchemaDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &schemaList, c.timeout)
	if err != nil {
		return nil, err
	}

	return schemaList, nil
}

// ListSchemaByID sends a request to retrieve a schema by its ID.
//
// Parameters:
//   - id: The ID of the schema.
//
// Returns:
//   - outputdto.SchemaDTO: The schema data transfer object.
//   - error: An error if the request fails.
func (c *Client) ListSchemaByID(id string) (outputdto.SchemaDTO, error) {
	pathParams := []string{"schema", id}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	var schemaOutput outputdto.SchemaDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &schemaOutput, c.timeout)
	if err != nil {
		return outputdto.SchemaDTO{}, err
	}

	return schemaOutput, nil
}

// DeleteSchema sends a request to delete a schema by its ID.
//
// Parameters:
//   - id: The ID of the schema.
//
// Returns:
//   - error: An error if the request fails.
func (c *Client) DeleteSchema(id string) error {
	pathParams := []string{"schema", id}

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

// ListSchemasByServiceAndProvider sends a request to retrieve schemas by service and provider.
//
// Parameters:
//   - service: The service name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.SchemaDTO: A slice of schema data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListSchemasByServiceAndProvider(service, provider string) ([]outputdto.SchemaDTO, error) {
	pathParams := []string{"schema", "provider", provider, "service", service}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var schemaList []outputdto.SchemaDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &schemaList, c.timeout)
	if err != nil {
		return nil, err
	}

	return schemaList, nil
}

// ListSchemasBySourceAndProvider sends a request to retrieve schemas by source and provider.
//
// Parameters:
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.SchemaDTO: A slice of schema data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListSchemasBySourceAndProvider(source, provider string) ([]outputdto.SchemaDTO, error) {
	pathParams := []string{"schema", "provider", provider, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var schemaList []outputdto.SchemaDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &schemaList, c.timeout)
	if err != nil {
		return nil, err
	}

	return schemaList, nil
}

// ListSchemasByServiceAndSourceAndProvider sends a request to retrieve schemas by service, source, and provider.
//
// Parameters:
//   - service: The service name.
//   - source: The source name.
//   - provider: The provider name.
//
// Returns:
//   - []outputdto.SchemaDTO: A slice of schema data transfer objects.
//   - error: An error if the request fails.
func (c *Client) ListSchemasByServiceAndSourceAndProvider(service, source, provider string) ([]outputdto.SchemaDTO, error) {
	pathParams := []string{"schema", "provider", provider, "service", service, "source", source}

	req, err := requests.CreateRequest(c.ctx, c.baseURL, pathParams, nil, nil, defaultHeaders, http.MethodGet)
	if err != nil {
		return nil, err
	}

	var schemaList []outputdto.SchemaDTO
	err = requests.SendRequest(c.ctx, req, requests.DefaultHTTPClient, &schemaList, c.timeout)
	if err != nil {
		return nil, err
	}

	return schemaList, nil
}
