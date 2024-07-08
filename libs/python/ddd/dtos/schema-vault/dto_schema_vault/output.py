from dataclasses import dataclass, field
from dto_schema_vault.shared import JsonSchemaDTO


@dataclass
class SchemaDTO:
    # ID is the unique identifier of the Schema entity.
    id: str = field(metadata={"json": "_id"})
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
    # SchemaVersionID is the unique identifier of the schema version.
    schema_version_id: str = field(metadata={"json": "schema_version_id"})
    # CreatedAt is the timestamp when the Schema entity was created.
    created_at: str = field(metadata={"json": "created_at"}, repr=False)
    # UpdatedAt is the timestamp when the Schema entity was last updated.
    updated_at: str = field(metadata={"json": "updated_at"}, repr=False)
