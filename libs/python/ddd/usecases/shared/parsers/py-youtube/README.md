# py-youtube

A Python library for downloading YouTube videos to a buffer, providing easy-to-use functions for handling YouTube video downloads.

## Features

- Download YouTube videos to a buffer.
- Custom error handling for YouTube download errors.

## Installation

To install the `py-youtube` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-ddd-usecases-shared-parsers-py-youtube --local
```

## Usage

### Download YouTube Video to Buffer

The `download_to_buffer` function allows you to download a YouTube video and store it in a buffer as bytes.

#### Example

```python
from pyyoutube import download_to_buffer, YoutubeDownloaderError

try:
    video_bytes = download_to_buffer("https://www.youtube.com/watch?v=example_video_id")
    # Do something with video_bytes
except YoutubeDownloaderError as e:
    print(f"Error downloading video: {e}")
```

### Classes

#### `YoutubeDownloaderError`

Custom exception class for YouTube downloader errors.

```python
from pyyoutube import YoutubeDownloaderError

# Raise a custom error
raise YoutubeDownloaderError("An error occurred while downloading the video")
```

### Functions

#### `download_to_buffer(url: str) -> bytes`

Downloads a video from the given URL to a buffer and returns its bytes.

```python
from pyyoutube import download_to_buffer

video_bytes = download_to_buffer("https://www.youtube.com/watch?v=example_video_id")
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-ddd-usecases-shared-parsers-py-youtube
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-ddd-usecases-shared-parsers-py-youtube --with dev
```
