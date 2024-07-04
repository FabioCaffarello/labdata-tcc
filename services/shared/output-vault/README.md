# Output Vault

Output Vault is a Go-based application designed for managing output data, health checks, and configuration management. It integrates with MongoDB to manage and process output entries.

## Features

- Health check endpoint
- CRUD operations for output data
- Dynamic routing for service, provider, and source-based queries

## Endpoints

### Health Check

- **GET /healthz**
  - Returns the health status of the application.

### Output Management

- **POST /output**
  - Creates a new output entry.
  - **Body**: JSON object with output details.

- **PUT /output**
  - Updates an existing output entry.
  - **Body**: JSON object with updated output details.

- **GET /output**
  - Lists all output entries.

- **GET /output/{id}**
  - Retrieves an output entry by its ID.

- **DELETE /output/{id}**
  - Deletes an output entry by its ID.

- **GET /output/provider/{provider}/service/{service}**
  - Lists outputs by service and provider.

- **GET /output/provider/{provider}/source/{source}**
  - Lists outputs by source and provider.

- **GET /output/provider/{provider}/service/{service}/source/{source}**
  - Lists outputs by service, source, and provider.

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
npx nx image services-shared-output-vault
```

## Docker Compose Configuration

#### Output Vault

- **Image**: fabiocaffarello/output-vault:latest
- **Environment Variables**:
  - `MONGODB_USER`: MongoDB username
  - `MONGODB_PASSWORD`: MongoDB password
  - `MONGODB_HOST`: MongoDB host
  - `MONGODB_PORT`: MongoDB port
  - `MONGODB_DBNAME`: MongoDB database name
- **Ports**: 8002:8000
- **Healthcheck**: Checks Output Vault health by calling the health endpoint.
