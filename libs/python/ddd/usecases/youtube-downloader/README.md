# youtube-downloader

A Python library for downloading YouTube videos, processing them, and uploading them to MinIO, providing an easy-to-use interface for handling video download and upload workflows.

## Features

- Download YouTube videos.
- Upload videos to MinIO.
- Structured handling of metadata and status.

## Installation

To install the `python-ddd-usecases-youtube-downloader` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-ddd-usecases-youtube-downloader --local
```

## Usage

### Job

The `Job` class provides methods to download YouTube videos, upload them to MinIO, and manage metadata and status.

#### Example

```python
from youtube_downloader.job import Job
from dto_config_vault.output import ConfigDTO
from dto_events_router.input import MetadataDTO
from pydebug import debug

# Initialize the job
config = ConfigDTO(...)  # Provide necessary configuration
metadata = MetadataDTO(...)  # Provide necessary metadata
dbg = debug.EnabledDebug(...)  # Initialize debug

job = Job(config, metadata, dbg)

# Run the job
input_data = ...  # Provide necessary input data
result = await job.run(input_data)
print(result)
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-ddd-usecases-youtube-downloader
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-ddd-usecases-youtube-downloader --with dev
```
