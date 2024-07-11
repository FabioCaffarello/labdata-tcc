from typing import Any, Dict, List
from dto_schema_vault.input import SchemaDataDTO


def get_schema(
    schema_id="schema-id",
    service="test-service",
    source="test-source",
    provider="provider",
    schema_type="type",
    schema_version_id="schema-version-id",
) -> Dict[str, Any]:
    json_schema = {
        "type": "object",
        "properties": {
            "field": {
                "type": "string"
            }
        },
        "required": ["field"]
    }
    return {
        "_id": schema_id,
        "service": service,
        "source": source,
        "provider": provider,
        "schema_type": schema_type,
        "json_schema": json_schema,
        "schema_version_id": schema_version_id,
        "created_at": "2024-02-01 00:00:00",
        "updated_at": "2024-02-01 00:00:00"
    }


def get_schemas() -> List[Dict[str, Any]]:
    return [
        get_schema(),
        get_schema(
            schema_id="schema-id-2",
            service="test-service-2",
            source="test-source-2",
            provider="provider-2",
            schema_type="type-2",
            schema_version_id="schema-version-id-2"
        )
    ]


def get_schema_data_dto() -> SchemaDataDTO:
    return SchemaDataDTO(
        service="test-service",
        source="test-source",
        provider="provider",
        schema_type="type",
        data={"field": "value"}
    )
