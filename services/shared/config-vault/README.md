# Config Vault

Config Vault is a Go-based application designed for managing configurations and health checks.

## Features

- Health check endpoint
- CRUD operations for configurations
- Dynamic routing for service and provider-based queries

## Endpoints

### Health Check

- **GET /healthz**
  - Returns the health status of the application.

### Configuration Management

- **POST /config**
  - Creates a new configuration.
  - **Body**: JSON object with configuration details.

- **PUT /config**
  - Updates an existing configuration.
  - **Body**: JSON object with updated configuration details.

- **GET /config**
  - Lists all configurations.

- **GET /config/{id}**
  - Retrieves a configuration by its ID.

- **DELETE /config/{id}**
  - Deletes a configuration by its ID.

- **GET /config/provider/{provider}/service/{service}**
  - Lists configurations by service and provider.

- **GET /config/provider/{provider}/source/{source}**
  - Lists configurations by source and provider.

- **GET /config/provider/{provider}/service/{service}/active/{active}**
  - Lists configurations by service, provider, and active status.

- **GET /config/provider/{provider}/service/{service}/source/{source}**
  - Lists configurations by service, source, and provider.

- **GET /config/provider/{provider}/dependencies/service/{service}/source/{source}**
  - Lists configurations by provider and dependencies.


## Building and Deploying

The application can be built and deployed using the provided Nx targets. The following targets are defined:

- **serve**: Runs the Go application.
- **test**: Runs the tests.
- **lint**: Lints the code.
- **tidy**: Tidies the Go modules.
- **godoc**: Generates Go documentation.
- **wire**: Runs Google Wire for dependency injection.
- **build**: Builds the Go application for a Linux environment.
- **image**: Builds the Docker image.

Example command to build the Docker image:
```bash
npx nx image services-shared-config-vault
```

## Docker Compose Configuration

#### Config Vault

- **Image**: fabiocaffarello/config-vault:latest
- **Environment Variables**:
  - `MONGODB_USER`: MongoDB username
  - `MONGODB_PASSWORD`: MongoDB password
  - `MONGODB_HOST`: MongoDB host
  - `MONGODB_PORT`: MongoDB port
  - `MONGODB_DBNAME`: MongoDB database name
- **Ports**: 8000:8000
- **Healthcheck**: Checks Config Vault health by calling the health endpoint.
