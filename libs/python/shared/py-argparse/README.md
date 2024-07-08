# py-argparse

`py-argparse` is a Python library that provides a convenient way to create an argument parser with predefined arguments for enabling debug storage and setting the debug storage directory. This library is useful for configuring command-line arguments in your applications, making it easier to enable debug storage and specify the storage directory.

## Features

- Create an argument parser with a description.
- Add arguments for enabling debug storage.
- Add arguments for setting the debug storage directory.

## Installation

To install the `py-argparse` library, run the following command:

```sh
npx nx run <PROJECT>:add --name python-shared-py-argparse --local
```

## Usage

### Creating an Argument Parser

To create a new argument parser with predefined arguments, use the `new` function:

```python
from pyargparse import new

# Create a new argument parser
parser = new(description="This is a sample argument parser")

# Parse the arguments
args = parser.parse_args()

# Access the arguments
print(f"Enable Debug Storage: {args.enable_debug_storage}")
print(f"Debug Storage Directory: {args.debug_storage_dir}")
```

## Example

Here is a complete example demonstrating how to use the `py-argparse` library:

```python
from pyargparse import new

def main():
    # Create a new argument parser with a description
    parser = new(description="This is a sample argument parser for debug storage")

    # Parse the command-line arguments
    args = parser.parse_args()

    # Access and print the parsed arguments
    print(f"Enable Debug Storage: {args.enable_debug_storage}")
    print(f"Debug Storage Directory: {args.debug_storage_dir}")

if __name__ == "__main__":
    main()
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-shared-py-argparse
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-shared-py-argparse --with dev
```
