import json
from typing import Dict, Type, Any
from dataclasses import dataclass, fields, is_dataclass


def _get_serialized_object(obj: dataclass) -> Dict[str, any]:
    """
    Serializes a dataclass object to a dictionary.

    Args:
        obj (dataclass): A dataclass object to be serialized.

    Returns:
        Dict[str, any]: A dictionary containing the serialized data from the dataclass object.
    """
    data = {}
    for field_obj in fields(obj):
        field_name = field_obj.name
        json_name = field_obj.metadata.get("json", field_name)
        field_value = getattr(obj, field_name)

        if field_value is not None:
            if is_dataclass(field_value):
                data[json_name] = _get_serialized_object(field_value)
            elif isinstance(field_value, list) and field_value and is_dataclass(field_value[0]):
                data[json_name] = [_get_serialized_object(item) for item in field_value]
            else:
                data[json_name] = field_value
    return data


def serialize_to_json(obj: dataclass) -> str:
    """
    Serializes a dataclass object to a JSON string.

    Args:
        obj (dataclass): A dataclass object to be serialized.

    Returns:
        str: A JSON string representing the serialized data from the dataclass object.
    """
    data = _get_serialized_object(obj)
    return json.dumps(data, sort_keys=True)


def serialize_to_dict(obj: dataclass) -> Dict[str, any]:
    """
    Serializes a dataclass object to a dictionary.

    Args:
        obj (dataclass): A dataclass object to be serialized.

    Returns:
        Dict[str, any]: A dictionary containing the serialized data from the dataclass object.
    """
    return _get_serialized_object(obj)


def serialize_to_dataclass(data: Dict[str, Any], cls: Type) -> Any:
    """
    Deserializes a dictionary to a dataclass object.

    Args:
        data (Dict[str, Any]): A dictionary containing the data to be deserialized.
        cls (Type): The dataclass type to which the dictionary will be deserialized.

    Returns:
        Any: A dataclass object containing the deserialized data.
    """
    if not is_dataclass(cls):
        raise TypeError('must be called with a dataclass type or instance')

    args = {}
    for field_obj in fields(cls):
        field_name = field_obj.name
        json_name = field_obj.metadata.get("json", field_name)

        if json_name in data:
            field_type = field_obj.type
            field_value = data[json_name]

            if is_dataclass(field_type):
                args[field_name] = serialize_to_dataclass(field_value, field_type)
            elif isinstance(field_value, list) and hasattr(field_type, '__origin__') and field_type.__origin__ == list:
                args[field_name] = [
                    serialize_to_dataclass(item, field_type.__args__[0])
                    if is_dataclass(field_type.__args__[0]) else item
                    for item in field_value
                ]
            else:
                args[field_name] = field_value

    return cls(**args)
