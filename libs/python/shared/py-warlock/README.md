# py-warlock

`py-warlock` is a Python library that provides functionality to deserialize data from dictionaries into dataclass objects using schemas defined with the Warlock library. This library is useful for converting dictionaries into dataclass instances for easy manipulation and usage within Python applications.

## Features

- Deserialize dictionaries to dataclass objects.
- Supports schema definitions using Warlock.
- Easy integration with existing Python applications.

## Installation

To install the library, run:

```sh
npx nx run <PROJECT>:add --name python-shared-py-warlock --local
```

## Usage

### Defining Schemas and Deserializing Data

First, define your schema using Warlock. For example:

```python
import warlock

schema = {
    "name": "Person",
    "properties": {
        "name": {"type": "string"},
        "age": {"type": "integer"},
        "address": {
            "type": "object",
            "properties": {
                "street": {"type": "string"},
                "city": {"type": "string"}
            }
        }
    }
}

Person = warlock.model_factory(schema)
```

### Deserializing from Dictionary

To deserialize a dictionary to a dataclass object, use the `serialize_to_dataclass` function:

```python
from pywarlock import serialize_to_dataclass

data = {
    "name": "John Doe",
    "age": 30,
    "address": {
        "street": "123 Main St",
        "city": "Anytown"
    }
}

person_obj = serialize_to_dataclass(schema, data)
print(person_obj)
```

## Example

Here is a complete example demonstrating how to use the `py-warlock` library:

```python
import warlock
from typing import Dict
from pywarlock import serialize_to_dataclass

# Define the schema
schema = {
    "name": "Person",
    "properties": {
        "name": {"type": "string"},
        "age": {"type": "integer"},
        "address": {
            "type": "object",
            "properties": {
                "street": {"type": "string"},
                "city": {"type": "string"}
            }
        }
    }
}

# Create the dataclass model using Warlock
Person = warlock.model_factory(schema)

# Sample input data
data = {
    "name": "John Doe",
    "age": 30,
    "address": {
        "street": "123 Main St",
        "city": "Anytown"
    }
}

# Deserialize the data into a dataclass object
person_obj = serialize_to_dataclass(schema, data)
print("Deserialized to dataclass:", person_obj)
```

## Running Tests

To run the tests, use pytest:

```sh
npx nx test python-shared-py-warlock
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-shared-py-warlock --with dev
```
