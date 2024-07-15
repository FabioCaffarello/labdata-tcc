# config-vault/dtos

## Overview

`config-vault/dtos` is a Python library that provides Data Transfer Objects (DTOs) for managing configuration data within a Domain-Driven Design (DDD) context. This library defines structured data formats to facilitate the transfer of configuration data between different layers and services in a clean and efficient manner.

## Features

- **ConfigDTO**: Represents configuration data, including active status, service, source, provider, dependencies, version ID, and timestamps.
- **JobDependenciesDTO**: Represents job dependencies, specifying the service and source.

## Installation

To install the `config-vault` dto library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-ddd-dtos-config-vault --local
```

## Usage

### ConfigDTO

`ConfigDTO` is a data class that represents configuration data. It includes fields for active status, service, source, provider, dependencies, version ID, and timestamps.

```python
from dto_config_vault.output import ConfigDTO
from dto_config_vault.shared import JobDependenciesDTO, JobParametersDTO

dependencies = [
    JobDependenciesDTO(service="dep-service", source="dep-source")
]

job_parameters = JobParametersDTO(
    parser_module="paser"
)

config = ConfigDTO(
    config_id="123",
    active=True,
    service="test-service",
    source="test-source",
    provider="provider",
    depends_on=dependencies,
    job_parameters=job_parameters,
    config_version_id="xyz123",
    created_at="2024-02-01 00:00:00",
    updated_at="2024-02-01 00:00:00"
)

print(config)
```

### JobDependenciesDTO

`JobDependenciesDTO` is a data class that represents job dependencies, specifying the service and source.

```python
from dto_config_vault.output import JobDependenciesDTO

dependency = JobDependenciesDTO(
    service="dep-service",
    source="dep-source"
)

print(dependency)
```

### 

```python
from dto_config_vault.shared import JobParametersDTO

job_parameters = JobParametersDTO(
    parser_module="paser"
)

print(job_parameters)
```
