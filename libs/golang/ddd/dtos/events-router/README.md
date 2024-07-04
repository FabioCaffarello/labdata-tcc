# input-broker/dto

`input-broker/dto` is a Go library that provides data transfer objects (DTOs) for managing input data within a system. This library includes structures for various DTOs, facilitating the transfer and validation of input data.

## Features

- Define DTOs for input entities.
- Facilitate data transfer between different components of the system.
- Ensure consistency and validation of input data.

## Usage

### Defining Input DTOs

The `InputDTO` struct represents an input data transfer object with attributes such as ID, processing ID, service, source, provider, stage, and data.

#### ErrMsgDTO

The `ErrMsgDTO` struct is used to represent error messages with attributes such as error, listener tag, and message.

```go
package main

import (
    "fmt"
    outputdto "libs/golang/ddd/dtos/events-router/output"
)

func main() {
    errMsg := outputdto.ErrMsgDTO{
        Err:         fmt.Errorf("an error occurred"),
        ListenerTag: "listener_1",
        Msg:         []byte("error message"),
    }

    fmt.Printf("ErrMsgDTO: %+v\n", errMsg)
}
```

#### ProcessOrderDTO

The `ProcessOrderDTO` struct is used to represent order processing data with attributes such as ID, processing ID, service, source, provider, stage, and data.

```go
package main

import (
    "fmt"
    outputdto "libs/golang/ddd/dtos/events-router/output"
)

func main() {
    order := outputdto.ProcessOrderDTO{
        ID:           "order_123",
        ProcessingID: "proc_456",
        Service:      "exampleService",
        Source:       "exampleSource",
        Provider:     "exampleProvider",
        Stage:        "initial",
        Data:         map[string]interface{}{"key1": "value1", "key2": "value2"},
    }

    fmt.Printf("ProcessOrderDTO: %+v\n", order)
}
```

## Testing

To run the tests for the `dto` package, use the following command:

```sh
npx nx test libs-golang-ddd-dtos-events-router
```

## Error Handling

The DTOs include attributes and methods for managing error scenarios, such as:

- Representing error messages with detailed information.
- Ensuring data consistency and validation during order processing.

These DTOs help in responding to various errors with appropriate messages and data structures.
