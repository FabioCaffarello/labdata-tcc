# webserver

`webserver` is a Go library that provides a simple and flexible HTTP server using the `chi` router. This library allows you to create and configure an HTTP server with default middlewares, register routes, and manage route groups efficiently.

## Features

- Create and configure an HTTP server with default middlewares.
- Register individual routes with different HTTP methods.
- Group routes under common prefixes.
- Easy-to-use interface for starting the server.

## Usage

### Creating and Configuring the Web Server

The `NewWebServer` function creates a new `Server` instance with the specified address.

```go
package main

import (
    "fmt"
    "net/http"
    "libs/golang/server/http/chi-webserver/server"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}

func main() {
    server := webserver.NewWebServer(":8080")
    server.ConfigureDefaults()
    server.RegisterRoute("GET", "/hello", helloHandler)

    if err := server.Start(); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
```

### Adding Default Middlewares

The `ConfigureDefaults` method sets up default middlewares for the server, including request ID, real IP, logger, recoverer, and a timeout of 60 seconds.

```go
func main() {
    server := webserver.NewWebServer(":8080")
    server.ConfigureDefaults()
    // Other configuration and route registration
}
```

### Registering Routes

The `RegisterRoute` method adds a new route with the specified HTTP method, URL pattern, and handler function.

```go
func main() {
    server := webserver.NewWebServer(":8080")
    server.ConfigureDefaults()

    server.RegisterRoute("GET", "/hello", helloHandler)
    server.RegisterRoute("POST", "/submit", submitHandler)

    if err := server.Start(); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
```

### Grouping Routes

The `RegisterRouteGroup` method allows you to group routes under a common prefix.

```go
func main() {
    server := webserver.NewWebServer(":8080")
    server.ConfigureDefaults()

    server.RegisterRouteGroup("/api", func(r chi.Router) {
        r.Get("/hello", helloHandler)
        r.Post("/submit", submitHandler)
    })

    if err := server.Start(); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
```

### Starting the Server

The `Start` method runs the web server on the specified address.

```go
func main() {
    server := webserver.NewWebServer(":8080")
    server.ConfigureDefaults()

    server.RegisterRoute("GET", "/hello", helloHandler)

    if err := server.Start(); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
```

## API

### Server

#### `NewWebServer(addr string) *Server`

Creates and returns a new `Server` instance with the specified address.

#### `ConfigureDefaults()`

Sets up default middlewares for the server, including request ID, real IP, logger, recoverer, and a timeout of 60 seconds.

#### `RegisterMiddlewares(middlewares ...func(http.Handler) http.Handler)`

Adds multiple middlewares to the server.

#### `RegisterRoute(method, pattern string, handler http.HandlerFunc, group ...string)`

Adds a new route with an HTTP method, pattern, and handler function. Optionally, a group name can be specified.

#### `RegisterRouteGroup(prefix string, routes func(r chi.Router))`

Registers a group of routes under a common prefix.

#### `Start() error`

Runs the web server on the specified address.

## Example

```go
package main

import (
    "fmt"
    "net/http"
    "libs/golang/server/http/chi-webserver/server"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}

func main() {
    server := webserver.NewWebServer(":8080")
    server.ConfigureDefaults()

    server.RegisterRoute("GET", "/hello", helloHandler)

    if err := server.Start(); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
```

## Testing

To run the tests for the `webserver` package, use the following command:

```sh
npx nx test libs-golang-server-http-chi-webserver
```
