from typing import Any, Dict


def get_config(
    config_id="123",
    active=True,
    config_version_id="xyz123",
    service="test-service",
    source="test-source",
    provider="provider",
    dep_service="dep-service",
    dep_source="dep-source"
) -> Dict[str, Any]:
    return {
        "_id": config_id,
        "active": active,
        "service": service,
        "source": source,
        "provider": provider,
        "job_parameters": {
            "parser_module": "test_parser",
        },
        "depends_on": [{
            "service": dep_service,
            "source": dep_source
        }],
        "config_version_id": config_version_id,
        "created_at": "2024-02-01 00:00:00",
        "updated_at": "2024-02-01 00:00:00 ",
    }


def get_configs() -> Dict[str, Any]:
    return [
        get_config(),
        get_config(
            config_id="456",
            active=False,
            dep_service="dep-service-2",
            dep_source="dep-source-2"
        )
    ]
