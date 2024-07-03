# config-vault/client

`config-vault/client` is a Go library that provides a client for interacting with the configuration vault API. This library includes functionalities for creating, reading, updating, and deleting configuration entities via HTTP requests.

## Features

- Create, read, update, and delete configuration entities via HTTP requests.
- List configurations based on various attributes such as service, provider, source, and dependencies.
- Handles request creation, sending, and response processing.

## Usage

### Creating a Client

The `Client` struct provides methods to interact with the configuration vault API.

```go
package main

import (
    "log"
    "fmt"
    "config-vault/client"
    inputdto "libs/golang/ddd/dtos/config-vault/input"
)

func main() {
    cli := client.NewClient()

    // Create a new configuration
    configInput := inputdto.ConfigDTO{
        Active:   true,
        Service:  "example-service",
        Source:   "example-source",
        Provider: "example-provider",
        DependsOn: []inputdto.JobDependenciesDTO{
            {Service: "dep-service", Source: "dep-source"},
        },
    }

    configOutput, err := cli.CreateConfig(configInput)
    if err != nil {
        log.Fatalf("Failed to create config: %v", err)
    }
    fmt.Printf("Created config: %+v\n", configOutput)

    // Other client methods can be used similarly...
}
```

### Client Methods

#### CreateConfig

Creates a new configuration.

```go
func (c *Client) CreateConfig(configInput inputdto.ConfigDTO) (outputdto.ConfigDTO, error)
```

#### UpdateConfig

Updates an existing configuration.

```go
func (c *Client) UpdateConfig(configInput inputdto.ConfigDTO) (outputdto.ConfigDTO, error)
```

#### ListAllConfigs

Lists all configurations.

```go
func (c *Client) ListAllConfigs() ([]outputdto.ConfigDTO, error)
```

#### ListConfigByID

Gets a configuration by its ID.

```go
func (c *Client) ListConfigByID(id string) (outputdto.ConfigDTO, error)
```

#### DeleteConfig

Deletes a configuration by its ID.

```go
func (c *Client) DeleteConfig(id string) error
```

#### ListConfigsByServiceAndProvider

Lists configurations by service and provider.

```go
func (c *Client) ListConfigsByServiceAndProvider(service, provider string) ([]outputdto.ConfigDTO, error)
```

#### ListConfigsBySourceAndProvider

Lists configurations by source and provider.

```go
func (c *Client) ListConfigsBySourceAndProvider(source, provider string) ([]outputdto.ConfigDTO, error)
```

#### ListConfigsByServiceAndProviderAndActive

Lists configurations by service, provider, and active status.

```go
func (c *Client) ListConfigsByServiceAndProviderAndActive(service, provider, active string) ([]outputdto.ConfigDTO, error)
```

#### ListConfigsByServiceAndSourceAndProvider

Lists configurations by service, source, and provider.

```go
func (c *Client) ListConfigsByServiceAndSourceAndProvider(service, source, provider string) ([]outputdto.ConfigDTO, error)
```

#### ListConfigsByProviderAndDependencies

Lists configurations by provider and dependencies.

```go
func (c *Client) ListConfigsByProviderAndDependencies(provider, service, source string) ([]outputdto.ConfigDTO, error)
```

## Testing

To run the tests for the `client` package, use the following command:

```sh
npx nx test libs-golang-clients-apis-config-vault
```

## Error Handling

The client methods include error handling for various scenarios, such as:

- Invalid request body
- Request timeout
- Internal server errors during API interaction

These errors are handled and returned as appropriate Go errors.