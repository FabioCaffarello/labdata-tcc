# config-vault/client

## Overview

`config-vault/client` is a Python client library for handling asynchronous interactions with a configuration service. It utilizes rate-limited HTTP requests to manage configurations efficiently, ensuring compliance with API rate limits. The library provides a variety of methods to create, update, retrieve, and delete configurations, as well as query configurations based on specific criteria.

## Features

- **Create Configurations**: Create new configurations asynchronously.
- **Update Configurations**: Update existing configurations asynchronously.
- **Retrieve Configurations**: Retrieve configurations by ID or list all configurations.
- **Delete Configurations**: Delete configurations by ID.
- **Query Configurations**: Retrieve configurations based on service, provider, source, and active status.

## Installation

To install the `config-vault` client library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-clients-apis-config-vault --local
```

## Usage

### Initialization

To create an instance of the `AsyncPyConfigVaultClient`, you need to provide the base URL of the configuration service:

```python
from your_module import async_py_config_vault_client

client = async_py_config_vault_client()
```

### Creating a Configuration

To create a new configuration:

```python
import asyncio

data = {
    "active": True,
    "service": "test-service",
    "source": "test-source",
    "provider": "provider",
    "depends_on": [{
        "service": "dep-service",
        "source": "dep-source"
    }]
}

async def create_config():
    config = await client.create_config(data)
    print(config)

asyncio.run(create_config())
```

### Updating a Configuration

To update an existing configuration:

```python
data = {
    "active": False,
    "service": "test-service-updated",
    "source": "test-source",
    "provider": "provider",
    "depends_on": [{
        "service": "dep-service",
        "source": "dep-source"
    }]
}

async def update_config():
    config = await client.update_config(data)
    print(config)

asyncio.run(update_config())
```

### Listing All Configurations

To retrieve a list of all configurations:

```python
async def list_all_configs():
    configs = await client.list_all_configs()
    for config in configs:
        print(config)

asyncio.run(list_all_configs())
```

### Retrieving a Configuration by ID

To retrieve a specific configuration by its ID:

```python
config_id = "123"

async def get_config_by_id():
    config = await client.get_config_by_id(config_id)
    print(config)

asyncio.run(get_config_by_id())
```

### Deleting a Configuration

To delete a specific configuration by its ID:

```python
config_id = "123"

async def delete_config():
    await client.delete_config(config_id)

asyncio.run(delete_config())
```

### Querying Configurations

To retrieve configurations based on service and provider:

```python
provider = "provider"
service = "test-service"

async def list_configs_by_service_and_provider():
    configs = await client.list_configs_by_service_and_provider(provider, service)
    for config in configs:
        print(config)

asyncio.run(list_configs_by_service_and_provider())
```

To retrieve configurations based on source and provider:

```python
provider = "provider"
source = "test-source"

async def list_configs_by_source_and_provider():
    configs = await client.list_configs_by_source_and_provider(provider, source)
    for config in configs:
        print(config)

asyncio.run(list_configs_by_source_and_provider())
```

To retrieve configurations based on service, provider, and active status:

```python
provider = "provider"
service = "test-service"
active = True

async def list_configs_by_service_provider_and_active():
    configs = await client.list_configs_by_service_provider_and_active(provider, service, active)
    for config in configs:
        print(config)

asyncio.run(list_configs_by_service_provider_and_active())
```

To retrieve configurations based on service, source, and provider:

```python
provider = "provider"
service = "test-service"
source = "test-source"

async def list_configs_by_service_source_and_provider():
    configs = await client.list_configs_by_service_source_and_provider(provider, service, source)
    for config in configs:
        print(config)

asyncio.run(list_configs_by_service_source_and_provider())
```

To retrieve configurations based on provider, service, and dependencies:

```python
provider = "provider"
service = "test-service"
source = "test-source"

async def list_configs_by_provider_and_dependencies():
    configs = await client.list_configs_by_provider_and_dependencies(provider, service, source)
    for config in configs:
        print(config)

asyncio.run(list_configs_by_provider_and_dependencies())
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-clients-apis-config-vault
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-clients-apis-schema-vault --with dev
```