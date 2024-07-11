# schema-vault/client

## Overview

`schema-vault/client` is a Python client library for handling asynchronous interactions with a schema service. It utilizes rate-limited HTTP requests to manage schemas efficiently, ensuring compliance with API rate limits. The library provides a variety of methods to create, update, retrieve, and delete schemas, as well as query schemas based on specific criteria.

## Features

- **Create Schemas**: Create new schemas asynchronously.
- **Update Schemas**: Update existing schemas asynchronously.
- **Retrieve Schemas**: Retrieve schemas by ID or list all schemas.
- **Delete Schemas**: Delete schemas by ID.
- **Query Schemas**: Retrieve schemas based on service, provider, source, and schema type.
- **Validate Schemas**: Validate schemas using provided data.

## Installation

To install the `schema-vault` client library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-clients-apis-schema-vault --local
```

## Usage

### Initialization

To create an instance of the `AsyncPySchemaVaultClient`, you need to provide the base URL of the schema service:

```python
from your_module import async_py_schema_vault_client

client = async_py_schema_vault_client()
```

### Creating a Schema

To create a new schema:

```python
import asyncio

data = {
    "service": "test-service",
    "source": "test-source",
    "provider": "provider",
    "schema_type": "type",
    "json_schema": {
        "type": "object",
        "properties": {
            "field": {"type": "string"}
        },
        "required": ["field"]
    }
}

async def create_schema():
    schema = await client.create(data)
    print(schema)

asyncio.run(create_schema())
```

### Updating a Schema

To update an existing schema:

```python
data = {
    "service": "test-service-updated",
    "source": "test-source",
    "provider": "provider",
    "schema_type": "type",
    "json_schema": {
        "type": "object",
        "properties": {
            "field": {"type": "string"}
        },
        "required": ["field"]
    }
}

async def update_schema():
    schema = await client.update_schema(data)
    print(schema)

asyncio.run(update_schema())
```

### Listing All Schemas

To retrieve a list of all schemas:

```python
async def list_all_schemas():
    schemas = await client.list_all_schemas()
    for schema in schemas:
        print(schema)

asyncio.run(list_all_schemas())
```

### Retrieving a Schema by ID

To retrieve a specific schema by its ID:

```python
schema_id = "schema-id"

async def get_schema_by_id():
    schema = await client.get_schema_by_id(schema_id)
    print(schema)

asyncio.run(get_schema_by_id())
```

### Deleting a Schema

To delete a specific schema by its ID:

```python
schema_id = "schema-id"

async def delete_schema():
    await client.delete_schema(schema_id)

asyncio.run(delete_schema())
```

### Querying Schemas

To retrieve schemas based on service and provider:

```python
provider = "provider"
service = "test-service"

async def list_schemas_by_service_and_provider():
    schemas = await client.list_schemas_by_service_and_provider(provider, service)
    for schema in schemas:
        print(schema)

asyncio.run(list_schemas_by_service_and_provider())
```

To retrieve schemas based on source and provider:

```python
provider = "provider"
source = "test-source"

async def list_schemas_by_source_and_provider():
    schemas = await client.list_schemas_by_source_and_provider(provider, source)
    for schema in schemas:
        print(schema)

asyncio.run(list_schemas_by_source_and_provider())
```

To retrieve schemas based on service, source, and provider:

```python
provider = "provider"
service = "test-service"
source = "test-source"

async def list_schemas_by_service_source_and_provider():
    schemas = await client.list_schemas_by_service_source_and_provider(provider, service, source)
    for schema in schemas:
        print(schema)

asyncio.run(list_schemas_by_service_source_and_provider())
```

To retrieve schemas based on service, source, provider, and schema type:

```python
provider = "provider"
service = "test-service"
source = "test-source"
schema_type = "type"

async def list_schemas_by_service_source_provider_and_schema_type():
    schemas = await client.list_schemas_by_service_source_provider_and_schema_type(provider, service, source, schema_type)
    for schema in schemas:
        print(schema)

asyncio.run(list_schemas_by_service_source_provider_and_schema_type())
```

### Validating a Schema

To validate a schema using provided data:

```python
data = {
    "service": "test-service",
    "source": "test-source",
    "provider": "provider",
    "schema_type": "type",
    "json_schema": {
        "type": "object",
        "properties": {
            "field": {
                "type": "string"
            }
        },
        "required": ["field"]
    }
}

async def validate_schema():
    is_valid = await client.validate_schema(data)
    print("Schema is valid:", is_valid)

asyncio.run(validate_schema())
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-clients-apis-schema-vault
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-clients-apis-schema-vault --with dev
```
