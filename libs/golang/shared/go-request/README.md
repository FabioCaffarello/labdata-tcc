# requests

The `go-request` is a Go library designed to simplify the creation, handling, and sending of HTTP requests. This library provides functions to build URLs, set headers, marshal request bodies, and send requests with a specified timeout.

## Features

- Parse and validate base URLs.
- Construct full URLs with path and query parameters.
- Marshal request bodies into JSON, XML, or URL-encoded forms.
- Set request headers.
- Create and send HTTP requests with context and timeout.

## Usage

### Parse Base URL

The `parseBaseURL` function parses a base URL string and returns a parsed `*url.URL` or an error if the URL is invalid.

```go
package main

import (
	"fmt"
	"libs/golang/shared/go-request/requests"
)

func main() {
	baseURL := "https://example.com"
	parsedURL, err := requests.parseBaseURL(baseURL)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed URL:", parsedURL)
	}
}
```

### Build URL

The `buildURL` function constructs a full URL with the given base URL, path parameters, and query parameters.

```go
package main

import (
	"fmt"
	"libs/golang/shared/go-request/requests"
)

func main() {
	baseURL := "https://example.com"
	pathParams := []string{"api", "v1", "resource"}
	queryParams := map[string]string{"query1": "value1", "query2": "value2"}
	fullURL, err := requests.buildURL(baseURL, pathParams, queryParams)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Full URL:", fullURL)
	}
}
```

### Create and Send Request

The `CreateRequest` function creates an HTTP request with the specified parameters, and the `SendRequest` function sends the request using the provided client, handling the response within the specified timeout.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"libs/golang/shared/go-request/requests"
)

func main() {
	ctx := context.Background()
	baseURL := "https://example.com"
	pathParams := []string{"api", "v1", "resource"}
	queryParams := map[string]string{"query1": "value1", "query2": "value2"}
	body := map[string]string{"key": "value"}
	headers := map[string]string{"Content-Type": "application/json"}
	method := http.MethodGet

	req, err := requests.CreateRequest(ctx, baseURL, pathParams, queryParams, body, headers, method)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	var result map[string]interface{}
	timeout := 5 * time.Second

	err = requests.SendRequest(ctx, req, client, &result, timeout)
	if err != nil {
		fmt.Println("Error sending request:", err)
	} else {
		fmt.Println("Response:", result)
	}
}
```

## Testing

To run the tests for the `requests` package, use the following command:

```sh
npx nx test libs-golang-shared-go-request
```
