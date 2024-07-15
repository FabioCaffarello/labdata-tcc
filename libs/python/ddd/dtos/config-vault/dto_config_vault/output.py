from dataclasses import dataclass, field
from dto_config_vault.shared import JobDependenciesDTO, JobParametersDTO
from typing import List


@dataclass
class ConfigDTO:
    config_id: str = field(metadata={"json": "_id"})
    active: bool = field(metadata={"json": "active"})
    service: str = field(metadata={"json": "service"})
    source: str = field(metadata={"json": "source"})
    provider: str = field(metadata={"json": "provider"})
    depends_on: List[JobDependenciesDTO] = field(metadata={"json": "depends_on"})
    job_parameters: JobParametersDTO = field(metadata={"json": "job_parameters"})
    config_version_id: str = field(metadata={"json": "config_version_id"})
    created_at: str = field(metadata={"json": "created_at"}, repr=False)
    updated_at: str = field(metadata={"json": "updated_at"}, repr=False)
