from dataclasses import dataclass, field
from typing import Dict, Any


@dataclass
class InputMetadataDTO:
    input_id: str = field(metadata={"json": "input_id"})
    schema_version_id: str = field(metadata={"json": "schema_version_id"})
    processing_order_id: str = field(metadata={"json": "processing_order_id"})


@dataclass
class OutputMetadataDTO:
    schema_version_id: str = field(metadata={"json": "schema_version_id"})


@dataclass
class MetadataDTO:
    provider: str = field(metadata={"json": "provider"})
    service: str = field(metadata={"json": "service"})
    source: str = field(metadata={"json": "source"})
    processing_id: str = field(metadata={"json": "processing_id"})
    config_id: str = field(metadata={"json": "config_id"})
    config_version_id: str = field(metadata={"json": "config_version_id"})
    input_metadata: InputMetadataDTO = field(metadata={"json": "input_metadata"})
    output_metadata: OutputMetadataDTO = field(metadata={"json": "output_metadata"})


@dataclass
class StatusDTO:
    code: str = field(metadata={"json": "code"})
    detail: str = field(metadata={"json": "detail"})


@dataclass
class ServiceFeedBackDTO:
    data: Dict[str, Any] = field(metadata={"json": "data"})
    metadata: MetadataDTO = field(metadata={"json": "metadata"})
    status: StatusDTO = field(metadata={"json": "status"})
