Sure, here is the README file for your debug library following the pattern provided:

# py-debug

`py-debug` is a Python library that provides functionality to enable or disable debug storage for saving responses. This library is useful for creating and managing debug storage, enabling you to save response data for debugging purposes.

## Features

- Enable or disable debug storage.
- Save response data to files in a specified directory.
- Automatically manage and create directories for storing debug responses.

## Installation

To install the `py-debug` library, run the following command:

```sh
npx nx run <PROJECT>:add --name python-shared-py-debug --local
```

## Usage

### Creating Debug Storage

To create a new debug storage instance, use the `new` function:

```python
from pydebug.debug import new

# Create an enabled debug storage instance
debug_storage = new(debug_enabled=True, debug_dir="/path/to/debug/dir")

# Create a disabled debug storage instance
debug_storage = new(debug_enabled=False, debug_dir="/path/to/debug/dir")
```

### Saving Response Data

To save response data using the debug storage, use the `save_response` method of the `EnabledDebug` class:

```python
from pydebug.debug import EnabledDebug

# Initialize enabled debug storage
debug_storage = EnabledDebug(debug_dir="/path/to/debug/dir")

# Save response data
debug_storage.save_response(file_name="response.txt", response_body=b"response content")
```

## Example

Here is a complete example demonstrating how to use the `py-debug` library:

```python
from pydebug.debug import new

def main():
    # Create an enabled debug storage instance
    debug_storage = new(debug_enabled=True, debug_dir="/app/tests/debug/storage")
    
    # Save response data
    debug_storage.save_response(file_name="response.txt", response_body=b"response content")

if __name__ == "__main__":
    main()
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-shared-py-debug
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-shared-py-debug --with dev
```
