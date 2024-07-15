import warlock
from typing import Dict


def serialize_to_dataclass(
    schema_parser_class: Dict[str, any],
    input_data: Dict[str, any]
) -> type[warlock.model.Model]:
    """
    Deserializes data from a dictionary into a dataclass object using a schema.

    Args:
        schema_parser_class (Dict[str, any]): A dictionary containing the schema.
        input_data (Dict[str, any]): A dictionary containing data to be deserialized.

    Returns:
        type[warlock.model.Model]: An instance of the dataclass type with data deserialized from the input dictionary.
    """
    Input_dataclass = warlock.model_factory(schema_parser_class)
    return Input_dataclass(**input_data)
