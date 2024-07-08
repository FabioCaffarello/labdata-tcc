# events-router/dtos

## Overview

`events-router/dtos` is a Python library that provides Data Transfer Objects (DTOs) for managing event routing data within a Domain-Driven Design (DDD) context. This library defines structured data formats to facilitate the transfer of event routing data between different layers and services in a clean and efficient manner.

## Features

- **ErrMsgDTO**: Represents error message data, including the error, listener tag, and message.
- **ProcessOrderDTO**: Represents order processing data, including order ID, processing ID, service, source, provider, stage, input ID, and data.

## Installation

To install the `events-router` DTO library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-ddd-dtos-events-router --local
```

## Usage

### ErrMsgDTO

`ErrMsgDTO` is a data class that represents error message data, including the error, listener tag, and message.

```python
from dto_events_router.output import ErrMsgDTO

error_message = ErrMsgDTO(
    err=Exception("An error occurred"),
    listener_tag="listener-1",
    msg=b"Error message"
)

print(error_message)
```

### ProcessOrderDTO

`ProcessOrderDTO` is a data class that represents order processing data. It includes fields for order ID, processing ID, service, source, provider, stage, input ID, and data.

```python
from dto_events_router.output import ProcessOrderDTO

order_data = {
    "key": "value"
}

process_order = ProcessOrderDTO(
    id="order-id",
    processing_id="processing-id",
    service="test-service",
    source="test-source",
    provider="provider",
    stage="stage-1",
    input_id="input-id",
    data=order_data
)

print(process_order)
```
