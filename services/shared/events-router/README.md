# Events Router

Events Router is a Go-based application designed for managing event routing, processing, and dispatching. It integrates with RabbitMQ and an in-memory document database to manage and process events.

## Features

- Event routing and processing
- Event dispatching using RabbitMQ
- Health check endpoint

## Usage

### Building and Deploying

The application can be built and deployed using the provided Nx targets. The following targets are defined:

- **serve**: Runs the Go application.
- **test**: Runs the tests.
- **lint**: Lints the code.
- **tidy**: Tidies the Go modules.
- **godoc**: Generates Go documentation.
- **build**: Builds the Go application for a Linux environment.
- **image**: Builds the Docker image.

Example command to build the Docker image:
```bash
npx nx image services-shared-events-router
```

## Docker Compose Configuration

#### Events Router

- **Image**: fabiocaffarello/events-router:latest
- **Environment Variables**:
  - `DOCDB_DBNAME`: Document database name
  - `CONSUMER_NAME`: Name of the consumer
  - `RABBITMQ_USER`: RabbitMQ username
  - `RABBITMQ_PASSWORD`: RabbitMQ password
  - `RABBITMQ_HOST`: RabbitMQ host
  - `RABBITMQ_PORT`: RabbitMQ port
  - `RABBITMQ_PROTOCOL`: RabbitMQ protocol
  - `RABBITMQ_EXCHANGE_NAME`: RabbitMQ exchange name
  - `RABBITMQ_EXCHANGE_TYPE`: RabbitMQ exchange type
