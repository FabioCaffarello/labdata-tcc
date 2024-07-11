# job-handler

A Python library for dynamically importing and executing job handlers based on configuration data.

## Features

- Dynamically imports job parser modules.
- Executes jobs using the imported job parsers.
- Handles configuration and metadata for job execution.

## Installation

To install the `job-handler` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-ddd-usecases-shared-job-handler --local
```

## Usage

### Importer

The `Importer` class handles the dynamic import of job parser modules.

#### Example

```python
from job_handler.importer import Importer

# Initialize the Importer
importer = Importer(job_parser="example_job_parser")

# Access the imported module
module = importer.module
print(module)
```

### JobHandler

The `JobHandler` class handles the execution of jobs using the imported job parser.

#### Example

```python
from job_handler.handler import JobHandler
from dto_config_vault.output import ConfigDTO
from dto_events_router.input import MetadataDTO
from pydebug import debug

# Initialize configuration and metadata
config = ConfigDTO(
    ...
)
metadata = MetadataDTO(
    ...
)
dbg = debug.EnabledDebug()

# Initialize the JobHandler
job_handler = JobHandler(config=config, metadata=metadata, dbg=dbg)

# Execute the job
source_input = ...  # Provide the input data for the job
result = await job_handler.execute(source_input)
print(result)
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-ddd-usecases-shared-job-handler
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-ddd-usecases-shared-job-handler --with dev
```
