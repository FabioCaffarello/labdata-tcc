# py-mongo

A Python library providing helper functions to interact with MongoDB, including connecting to a MongoDB instance and dropping databases.

## Features

- Connect to a MongoDB instance using environment configurations.
- Drop specified databases.

## Installation

To install the `py-mongo` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-clients-resources-py-mongo --local
```

## Usage

### get_mongo_client

The `get_mongo_client` function retrieves a MongoDB client connected to the MongoDB endpoint defined in the environment.

#### Example

```python
from pymongo_helper import get_mongo_client

# Get the MongoDB client
client = get_mongo_client()
print(client.list_database_names())
```

### drop_database

The `drop_database` function drops the specified database from the MongoDB server.

#### Example

```python
from pymongo_helper import drop_database

# Drop the database named "test_db"
drop_database("test_db")
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-clients-resources-py-mongo
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-clients-resources-py-mongo --with dev
```