# events-router/dtos

`events-router/dtos` is a Python library that provides Data Transfer Objects (DTOs) for managing event routing data within a Domain-Driven Design (DDD) context. This library defines structured data formats to facilitate the transfer of event routing data between different layers and services in a clean and efficient manner.

## Features

- **ErrMsgDTO**: Represents error message data, including the error, listener tag, and message.
- **ProcessOrderDTO**: Represents order processing data, including order ID, processing ID, service, source, provider, stage, input ID, and data.
- **InputMetadataDTO**: Represents input metadata information.
- **OutputMetadataDTO**: Represents output metadata information.
- **MetadataDTO**: Represents comprehensive metadata including provider, service, source, processing ID, config ID, input metadata, and output metadata.
- **StatusDTO**: Represents status information including code and detail.
- **ServiceFeedBackDTO**: Represents service feedback including data, metadata, and status.

## Installation

To install the `events-router` DTO library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-ddd-dtos-events-router --local
```

## Usage

### ErrMsgDTO

`ErrMsgDTO` is a data class that represents error message data, including the error, listener tag, and message.

```python
from dto_events_router.input import ErrMsgDTO

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
    order_id="order-id",
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

### InputMetadataDTO

`InputMetadataDTO` is a data class that represents input metadata information.

```python
from dto_events_router.input import InputMetadataDTO

input_metadata = InputMetadataDTO(
    input_id="input-id",
    schema_version_id="schema-version-id",
    processing_order_id="processing-order-id"
)

print(input_metadata)
```

### OutputMetadataDTO

`OutputMetadataDTO` is a data class that represents output metadata information.

```python
from dto_events_router.input import OutputMetadataDTO

output_metadata = OutputMetadataDTO(
    schema_version_id="schema-version-id"
)

print(output_metadata)
```

### MetadataDTO

`MetadataDTO` is a data class that represents comprehensive metadata including provider, service, source, processing ID, config ID, input metadata, and output metadata.

```python
from dto_events_router.input import MetadataDTO, InputMetadataDTO, OutputMetadataDTO

input_metadata = InputMetadataDTO(
    input_id="input-id",
    schema_version_id="schema-version-id",
    processing_order_id="processing-order-id"
)

output_metadata = OutputMetadataDTO(
    schema_version_id="schema-version-id"
)

metadata = MetadataDTO(
    provider="test-provider",
    service="test-service",
    source="test-source",
    processing_id="processing-id",
    config_id="config-id",
    input_metadata=input_metadata,
    output_metadata=output_metadata
)

print(metadata)
```

### StatusDTO

`StatusDTO` is a data class that represents status information including code and detail.

```python
from dto_events_router.input import StatusDTO

status = StatusDTO(
    code="200",
    detail="Success"
)

print(status)
```

### ServiceFeedBackDTO

`ServiceFeedBackDTO` is a data class that represents service feedback including data, metadata, and status.

```python
from dto_events_router.input import ServiceFeedBackDTO, MetadataDTO, InputMetadataDTO, OutputMetadataDTO, StatusDTO

input_metadata = InputMetadataDTO(
    input_id="input-id",
    schema_version_id="schema-version-id",
    processing_order_id="processing-order-id"
)

output_metadata = OutputMetadataDTO(
    schema_version_id="schema-version-id"
)

metadata = MetadataDTO(
    provider="test-provider",
    service="test-service",
    source="test-source",
    processing_id="processing-id",
    config_id="config-id",
    input_metadata=input_metadata,
    output_metadata=output_metadata
)

status = StatusDTO(
    code="200",
    detail="Success"
)

service_feedback = ServiceFeedBackDTO(
    data={"key": "value"},
    metadata=metadata,
    status=status
)

print(service_feedback)
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-ddd-dtos-events-router
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-ddd-dtos-events-router --with dev
```
