# py-log

`py-log` is a Python library that provides functionality to set up logging using JSON format. This library is useful for configuring logging in your applications to output logs in a structured JSON format, making it easier to parse and analyze logs.

## Features

- Set up logging with JSON format.
- Configurable log levels.
- Option to propagate logs to the parent logger.

## Installation

```sh
npx nx run <PROJECT>:add --name python-shared-py-log --local
```

## Usage

### Setting Up Logging

To set up logging with JSON format, use the `setup_logging` function:

```python
from pylog import setup_logging

logger = setup_logging(module_name="my_module")

logger.info("This is an info message")
logger.error("This is an error message")
```

### Configuring Log Level

You can configure the log level by setting the `LOG_LEVEL` environment variable or by passing the `log_level` argument to the `setup_logging` function:

```python
import os
from pylog import setup_logging

os.environ["LOG_LEVEL"] = "DEBUG"
logger = setup_logging(module_name="my_module")

logger.debug("This is a debug message")
```

### Propagating Logs

By default, logs are not propagated to the parent logger. You can enable log propagation by setting the `propagate` argument to `True`:

```python
from pylog import setup_logging

logger = setup_logging(module_name="my_module", propagate=True)

logger.info("This message will be propagated to the parent logger")
```

## Example

Here is a complete example demonstrating how to use the `py-log` library:

```python
import os
from pylog import setup_logging

def main():
    # Set the log level through environment variable
    os.environ["LOG_LEVEL"] = "DEBUG"
    
    # Set up logging for the module
    logger = setup_logging(module_name="example_module")
    
    # Log messages with different log levels
    logger.debug("Debugging information")
    logger.info("Informational message")
    logger.warning("Warning message")
    logger.error("Error message")
    logger.critical("Critical message")

if __name__ == "__main__":
    main()
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-shared-py-log
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-shared-py-log --with dev
```