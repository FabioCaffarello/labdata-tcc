# config-loader

`config-loader` is a Python library that provides functionality for loading configurations from a configuration handler API client. This library is designed to facilitate the fetching and managing of configuration data for different services and providers in a clean and efficient manner.

## Features

- **ConfigLoader**: Class to load configurations from the config handler API client.
- **fetch_configs**: Function to fetch configurations for a given service and provider.

## Installation

To install the `config-loader` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-ddd-shared-config-loader --local
```

## Usage

### ConfigLoader

The `ConfigLoader` class provides methods to fetch and register configurations for different services and providers.

#### Example

```python
from config_loader.loader import ConfigLoader
import asyncio

async def main():
    config_loader = ConfigLoader()
    configs = await config_loader.fetch_configs_for_service("example_service", "example_provider")
    print(configs)

asyncio.run(main())
```

### Methods

#### `fetch_configs_for_service(service_name: str, provider: str) -> Dict[str, ConfigDTO]`

Fetches configurations for a specific service and provider.

```python
configs = await config_loader.fetch_configs_for_service("example_service", "example_provider")
print(configs)
```

#### `register_config(config_id: str, config: ConfigDTO) -> None`

Registers a configuration in the `mapping_config` dictionary.

```python
config_loader.register_config("config_id", config)
```

### fetch_configs

The `fetch_configs` function provides a simple way to fetch configurations for a given service and provider.

#### Example

```python
from config_loader.loader import fetch_configs
import asyncio

async def main():
    configs = await fetch_configs("example_service", "example_provider")
    print(configs)

asyncio.run(main())
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-ddd-shared-config-loader
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-ddd-shared-config-loader --with dev
```
