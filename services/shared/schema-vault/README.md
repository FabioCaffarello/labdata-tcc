# Schema Vault

Schema Vault is a Go-based application designed for managing schema data, health checks, and configuration management. It integrates with MongoDB to manage and process schema entries.

## Features

- Health check endpoint
- CRUD operations for schema data
- Dynamic routing for service, provider, and source-based queries

## Endpoints

### Health Check

- **GET /healthz**
  - Returns the health status of the application.

### Schema Management

- **POST /schema**
  - Creates a new schema entry.
  - **Body**: JSON object with schema details.

- **PUT /schema**
  - Updates an existing schema entry.
  - **Body**: JSON object with updated schema details.

- **GET /schema**
  - Lists all schema entries.

- **GET /schema/{id}**
  - Retrieves a schema entry by its ID.

- **DELETE /schema/{id}**
  - Deletes a schema entry by its ID.

- **GET /schema/provider/{provider}/service/{service}**
  - Lists schemas by service and provider.

- **GET /schema/provider/{provider}/source/{source}**
  - Lists schemas by source and provider.

- **GET /schema/provider/{provider}/service/{service}/source/{source}**
  - Lists schemas by service, source, and provider.

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
npx nx image services-shared-schema-vault
```

## Docker Compose Configuration

#### Schema Vault

- **Image**: fabiocaffarello/schema-vault:latest
- **Environment Variables**:
  - `MONGODB_USER`: MongoDB username
  - `MONGODB_PASSWORD`: MongoDB password
  - `MONGODB_HOST`: MongoDB host
  - `MONGODB_PORT`: MongoDB port
  - `MONGODB_DBNAME`: MongoDB database name
- **Ports**: 8001:8000
- **Healthcheck**: Checks Schema Vault health by calling the health endpoint.
