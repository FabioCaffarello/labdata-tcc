# schema-vault/dtos

## Overview

`schema-vault/dtos` is a Python library that provides Data Transfer Objects (DTOs) for managing schema data within a Domain-Driven Design (DDD) context. This library defines structured data formats to facilitate the transfer of schema data between different layers and services in a clean and efficient manner.

## Features

- **SchemaDTO**: Represents schema configuration data, including service, source, provider, schema type, JSON schema, version ID, and timestamps.
- **SchemaDataDTO**: Represents schema data, specifying the service, source, provider, schema type, and actual data.
- **JsonSchemaDTO**: Represents the JSON schema, including required fields, properties, and type.

## Installation

To install the `schema-vault` DTO library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-ddd-dtos-schema-vault --local
```

## Usage

### SchemaDTO

`SchemaDTO` is a data class that represents schema configuration data. It includes fields for service, source, provider, schema type, JSON schema, version ID, and timestamps.

```python
from dto_schema_vault.output import SchemaDTO, JsonSchemaDTO

json_schema = JsonSchemaDTO(
    required=["field"],
    properties={
        "field": {"type": "string"}
    },
    json_type="object"
)

schema = SchemaDTO(
    schema_id="schema-id",
    service="test-service",
    source="test-source",
    provider="provider",
    schema_type="type",
    json_schema=json_schema,
    schema_version_id="schema-version-id",
    created_at="2024-02-01 00:00:00",
    updated_at="2024-02-01 00:00:00"
)

print(schema)
```

### SchemaDataDTO

`SchemaDataDTO` is a data class that represents schema data, specifying the service, source, provider, schema type, and actual data.

```python
from dto_schema_vault.output import SchemaDataDTO

schema_data = SchemaDataDTO(
    service="test-service",
    source="test-source",
    provider="provider",
    schema_type="type",
    data={
        "field": "value"
    }
)

print(schema_data)
```

### JsonSchemaDTO

`JsonSchemaDTO` is a data class that represents the JSON schema, including required fields, properties, and type.

```python
from dto_schema_vault.output import JsonSchemaDTO

json_schema = JsonSchemaDTO(
    required=["field"],
    properties={
        "field": {"type": "string"}
    },
    json_type="object"
)

print(json_schema)
```
