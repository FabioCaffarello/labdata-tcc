from dataclasses import dataclass, field
from typing import Dict, Any,  List


@dataclass
class JsonSchemaDTO:
    # Required lists the required fields in the JSON schema.
    required: List[str] = field(metadata={"json": "required"})
    # Properties lists the properties in the JSON schema.
    properties: Dict[str, Any] = field(metadata={"json": "properties"})
    # JsonType specifies the type of JSON schema.
    json_type: str = field(metadata={"json": "type"})
