# healthz

The `healthz` package is a Go library that provides a handler for HTTP health checks. This library allows you to monitor the health status of your server by checking its uptime and readiness. It supports customizable time providers to facilitate testing and real-time monitoring.

## Features

- HTTP handler for health checks
- Customizable time providers for flexibility and testability
- Responds with appropriate HTTP status codes based on server uptime

## Usage

### Creating a Health Check Handler

The `WebHealthzHandler` struct is used to handle health check requests. You can create a new instance of this handler using the `NewWebHealthzHandler` function, which accepts a `TimeProvider` to manage time-related functions.

#### Example with RealTimeProvider

```go
package main

import (
    "net/http"
    "libs/golang/ddd/adapters/http/handlers/health-check/healthz"
)

func main() {
    timeProvider := &healthz.RealTimeProvider{}
    handler := healthz.NewWebHealthzHandler(timeProvider)

    http.HandleFunc("/healthz", handler.Healthz)
    http.ListenAndServe(":8080", nil)
}
```

### Implementing Custom Time Providers

You can implement the `TimeProvider` interface to create your own custom time providers. This is useful for testing or for integrating with other time-based systems.

#### Example

```go
package main

import (
    "time"
)

type CustomTimeProvider struct{}

func (c *CustomTimeProvider) Now() time.Time {
    // Custom implementation
}

func (c *CustomTimeProvider) Since(t time.Time) time.Duration {
    // Custom implementation
}
```

## Testing

To run the tests for the `healthz` package, use the following command:

```sh
npx nx test libs-golang-ddd-adapters-http-handlers-health-check
```
