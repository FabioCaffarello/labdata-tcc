# Input Broker

Input Broker is a Go-based application designed for handling input data, health checks, and event dispatching. It integrates with MongoDB and RabbitMQ to manage and process input messages.

## Features

- Health check endpoint
- Input data processing
- Event dispatching using RabbitMQ

## Endpoints

### Health Check

- **GET /healthz**
  - Returns the health status of the application.

### Input Management

- **POST /input**
  - Creates a new input entry.
  - **Body**: JSON object with input details.

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
npx nx image services-shared-input-broker
```

## Docker Compose Configuration

#### Input Broker

- **Image**: fabiocaffarello/input-broker:latest
- **Environment Variables**:
  - `MONGODB_USER`: MongoDB username
  - `MONGODB_PASSWORD`: MongoDB password
  - `MONGODB_HOST`: MongoDB host
  - `MONGODB_PORT`: MongoDB port
  - `MONGODB_DBNAME`: MongoDB database name
  - `RABBITMQ_USER`: RabbitMQ username
  - `RABBITMQ_PASSWORD`: RabbitMQ password
  - `RABBITMQ_HOST`: RabbitMQ host
  - `RABBITMQ_PORT`: RabbitMQ port
  - `RABBITMQ_PROTOCOL`: RabbitMQ protocol
  - `RABBITMQ_EXCHANGE_NAME`: RabbitMQ exchange name
  - `RABBITMQ_EXCHANGE_TYPE`: RabbitMQ exchange type
- **Ports**: 8003:8000
- **Healthcheck**: Checks Input Broker health by calling the health endpoint.