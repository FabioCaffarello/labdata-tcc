from dataclasses import dataclass, field
from dto_config_vault.shared import JobDependenciesDTO, JobParametersDTO
from typing import List


@dataclass
class ConfigDTO:
    active: bool = field(metadata={"json": "active"})
    service: str = field(metadata={"json": "service"})
    source: str = field(metadata={"json": "source"})
    provider: str = field(metadata={"json": "provider"})
    job_parameters: JobParametersDTO = field(metadata={"json": "job_parameters"})
    depends_on: List[JobDependenciesDTO] = field(metadata={"json": "depends_on"})
