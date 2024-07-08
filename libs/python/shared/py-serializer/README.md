# py-serializer

`py-serializer` is a Python library that provides functionality to serialize and deserialize dataclass objects to and from JSON and dictionaries. This library is useful for converting dataclass instances to JSON strings or dictionaries for easy storage or transmission, and vice versa.

## Features

- Serialize dataclass objects to JSON strings.
- Serialize dataclass objects to dictionaries.
- Deserialize dictionaries to dataclass objects.
- Supports nested dataclass objects.

## Installation

```sh
npx nx run <PROJECT>:add --name python-shared-py-serializer --local
```

## Usage

### Creating Dataclass Objects

First, define your dataclass objects. For example:

```python
from dataclasses import dataclass

@dataclass
class Address:
    street: str
    city: str

@dataclass
class Person:
    name: str
    age: int
    address: Address
```

### Serializing to JSON

To serialize a dataclass object to a JSON string, use the `serialize_to_json` function:

```python
from pyserializer.serializer import serialize_to_json

address = Address(street="123 Main St", city="Anytown")
person = Person(name="John Doe", age=30, address=address)

json_str = serialize_to_json(person)
print(json_str)
```

### Serializing to Dictionary

To serialize a dataclass object to a dictionary, use the `serialize_to_dict` function:

```python
from pyserializer.serializer import serialize_to_dict

dict_obj = serialize_to_dict(person)
print(dict_obj)
```

### Deserializing from Dictionary

To deserialize a dictionary to a dataclass object, use the `serialize_to_dataclass` function:

```python
from pyserializer.serializer import serialize_to_dataclass

data = {
    "name": "John Doe",
    "age": 30,
    "address": {
        "street": "123 Main St",
        "city": "Anytown"
    }
}

person_obj = serialize_to_dataclass(data, Person)
print(person_obj)
```

## Example

Here is a complete example demonstrating how to use the `py-serializer` library:

```python
from dataclasses import dataclass
from dataclass_serializer import serialize_to_json, serialize_to_dict, serialize_to_dataclass

@dataclass
class Address:
    street: str
    city: str

@dataclass
class Person:
    name: str
    age: int
    address: Address

def main():
    address = Address(street="123 Main St", city="Anytown")
    person = Person(name="John Doe", age=30, address=address)
    
    # Serialize to JSON
    json_str = serialize_to_json(person)
    print("Serialized to JSON:", json_str)
    
    # Serialize to dictionary
    dict_obj = serialize_to_dict(person)
    print("Serialized to dict:", dict_obj)
    
    # Deserialize from dictionary
    data = {
        "name": "John Doe",
        "age": 30,
        "address": {
            "street": "123 Main St",
            "city": "Anytown"
        }
    }
    person_obj = serialize_to_dataclass(data, Person)
    print("Deserialized to dataclass:", person_obj)

if __name__ == "__main__":
    main()
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-shared-py-serializer
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-shared-py-serializer --with dev
```
