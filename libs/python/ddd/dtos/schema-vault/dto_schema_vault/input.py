from dataclasses import dataclass, field
from typing import Dict, Any
from dto_schema_vault.shared import JsonSchemaDTO


@dataclass
class SchemaDTO:
    # Service represents the name of the service for which the configuration is created.
    service: str = field(metadata={"json": "service"})
    # Source indicates the origin or source of the configuration.
    source: str = field(metadata={"json": "source"})
    # Provider specifies the provider of the configuration.
    provider: str = field(metadata={"json": "provider"})
    # SchemaType specifies the type of schema.
    schema_type: str = field(metadata={"json": "schema_type"})
    # JsonSchemaDTO represents the JSON schema of the configuration.
    json_schema: JsonSchemaDTO = field(metadata={"json": "json_schema"})


@dataclass
class SchemaDataDTO:
    # Service represents the name of the service for which the configuration is created.
    service: str = field(metadata={"json": "service"})
    # Source indicates the origin or source of the configuration.
    source: str = field(metadata={"json": "source"})
    # Provider specifies the provider of the configuration.
    provider: str = field(metadata={"json": "provider"})
    # SchemaType specifies the type of schema.
    schema_type: str = field(metadata={"json": "schema_type"})
    # Data represents the data of the respective schema type.
    data: Dict[str, Any] = field(metadata={"json": "data"})
