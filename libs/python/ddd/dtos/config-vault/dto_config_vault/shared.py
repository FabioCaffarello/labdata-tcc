from dataclasses import dataclass, field


@dataclass
class JobDependenciesDTO:
    service: str = field(metadata={"json": "service"})
    source: str = field(metadata={"json": "source"})


@dataclass
class JobParametersDTO:
    parser_module: str = field(metadata={"json": "parser_module"})
