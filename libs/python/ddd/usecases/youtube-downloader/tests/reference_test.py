from typing import List

from dto_config_vault.output import ConfigDTO
from dto_config_vault.shared import JobDependenciesDTO, JobParametersDTO
from dto_events_router.input import (InputMetadataDTO, MetadataDTO,
                                     OutputMetadataDTO)


def get_config(
    config_id="test_id",
    active=True,
    config_version_id="xyz123",
    service="test-service",
    source="test-source",
    provider="provider",
    dep_service="dep-service",
    dep_source="dep-source"
) -> ConfigDTO:
    return ConfigDTO(
        config_id=config_id,
        active=active,
        service=service,
        source=source,
        provider=provider,
        job_parameters=JobParametersDTO(
            parser_module="test_parser",
        ),
        depends_on=[
            JobDependenciesDTO(
                service=dep_service,
                source=dep_source
            )],
        config_version_id=config_version_id,
        created_at="2024-02-01 00:00:00",
        updated_at="2024-02-01 00:00:00 ",
    )


def get_metadata(
    config_id="test_id",
    service="test-service",
    source="test-source",
    provider="provider",
    processing_id="123"
) -> MetadataDTO:
    return MetadataDTO(
        config_id=config_id,
        service=service,
        source=source,
        provider=provider,
        processing_id=processing_id,
        input_metadata=InputMetadataDTO(
            input_id="input_id",
            schema_version_id="schema_version_id",
            processing_order_id="processing_order_id",
        ),
        config_version_id="config_version_id",
        output_metadata=OutputMetadataDTO(
            schema_version_id="schema_version_id_2",
        )
    )


def get_configs() -> List[ConfigDTO]:
    return [
        get_config(config_id="456", service="test-service-2", source="test-source-2"),
        get_config(config_id="457", service="test-service-2", source="test-source-3"),
    ]
