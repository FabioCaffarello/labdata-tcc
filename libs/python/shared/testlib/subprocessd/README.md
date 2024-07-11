# subprocessd

A Python library for managing asynchronous subprocess execution with logging capabilities.

## Features

- Start and stop subprocesses asynchronously.
- Log subprocess output to a specified directory.

## Installation

To install the `subprocessd` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-shared-testlib-subprocessd --local
```

## Usage

### SubprocessDAsync

The `SubprocessDAsync` class provides methods to manage asynchronous subprocess execution and logging.

#### Example

```python
import asyncio
from subprocessd.subprocessd import SubprocessDAsync

async def main():
    subprocess_args = ["ls", "-l"]
    log_file_path = "/path/to/logs"

    subprocess_manager = SubprocessDAsync(subprocess_args, log_file_path)

    # Start the subprocess
    await subprocess_manager.start()

    # Do other async tasks
    await asyncio.sleep(10)

    # Stop the subprocess
    await subprocess_manager.stop()

asyncio.run(main())
```


## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-shared-testlib-subprocessd
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-shared-testlib-subprocessd --with dev
```
