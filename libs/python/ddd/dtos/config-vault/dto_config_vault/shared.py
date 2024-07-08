from dataclasses import dataclass, field


@dataclass
class JobDependenciesDTO:
    service: str = field(metadata={"json": "service"})
    source: str = field(metadata={"json": "source"})
