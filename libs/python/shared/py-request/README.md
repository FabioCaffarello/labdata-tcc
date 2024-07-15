# py-request

`py-request` is a Python library providing an asynchronous HTTP client with built-in rate limiting. This library helps you make HTTP requests while ensuring that you do not exceed a maximum number of requests within a specified time period.

## Features

- Asynchronous HTTP requests using `httpx`
- Configurable rate limiting to prevent exceeding a specified number of requests in a given time period
- Supports various HTTP methods (GET, POST, etc.)

## Installation

```sh
npx nx run <PROJECT>:add --name python-shared-py-request --local
```

## Usage

### Creating a RateLimitedAsyncHttpClient

```python
from pyrequest.async_factory import RateLimitedAsyncHttpClient

# Initialize the client with a base URL, maximum number of calls, and time period (in seconds)
client = RateLimitedAsyncHttpClient(base_url="https://example.com", max_calls=5, period=60)
```

### Making Requests

You can make asynchronous HTTP requests using the `make_request` method. The method supports different HTTP methods such as GET and POST.

#### GET Request

```python
import asyncio

async def main():
    response = await client.make_request("GET", "/api/resource")
    print(response)

asyncio.run(main())
```

#### POST Request

```python
import asyncio

async def main():
    data = {"key": "value"}
    response = await client.make_request("POST", "/api/resource", data=data)
    print(response)

asyncio.run(main())
```

### Handling Rate Limiting

The `RateLimitedAsyncHttpClient` ensures that you do not exceed the maximum number of requests within the specified time period. If the limit is reached, it waits for the period to elapse before allowing more requests.

## Example

Here is a complete example demonstrating how to use the `RateLimitedAsyncHttpClient`:

```python
import asyncio
from pyrequest.async_factory import RateLimitedAsyncHttpClient

async def main():
    client = RateLimitedAsyncHttpClient(base_url="https://example.com", max_calls=5, period=60)
    
    # Make a GET request
    response = await client.make_request("GET", "/api/resource")
    print("GET response:", response)
    
    # Make a POST request
    data = {"key": "value"}
    response = await client.make_request("POST", "/api/resource", data=data)
    print("POST response:", response)
    
    # Simulate rate limiting by making multiple requests
    tasks = [client.make_request("GET", "/api/resource") for _ in range(10)]
    responses = await asyncio.gather(*tasks, return_exceptions=True)
    for i, res in enumerate(responses):
        print(f"Response {i+1}:", res)

asyncio.run(main())
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-shared-py-request
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-shared-py-request --with dev
```
