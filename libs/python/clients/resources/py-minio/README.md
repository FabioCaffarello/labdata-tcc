# py-minio

A Python library for interacting with Minio, providing easy-to-use client class methods for managing buckets and objects on a Minio server.

## Features

- Create and delete buckets.
- Upload and download files and bytes.
- List buckets and objects.
- Generate URIs for accessing objects.

## Installation

To install the `py-minio` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-clients-resources-py-minio --local
```

## Usage

### MinioClient

The `MinioClient` class provides methods to interact with a Minio server, including creating buckets, uploading files, and downloading files.

#### Example

```python
from pyminio.client import MinioClient

# Initialize the Minio client
minio = MinioClient(endpoint="http://localhost:9000", access_key="your_access_key", secret_key="your_secret_key")

# Create a new bucket
minio.create_bucket("my_bucket")

# Upload a file
minio.upload_file("my_bucket", "example.txt", "path/to/local/file.txt")

# List objects in the bucket
objects = minio.list_objects("my_bucket")
print(objects)

# Download a file
minio.download_file("my_bucket", "example.txt", "path/to/download/file.txt")
```

### Methods

#### `create_bucket(bucket_name: str) -> None`

Creates a new bucket on the Minio server.

```python
minio.create_bucket("my_bucket")
```

#### `list_buckets() -> List[str]`

Lists all buckets available on the Minio server.

```python
buckets = minio.list_buckets()
print(buckets)
```

#### `upload_file(bucket_name: str, object_name: str, file_path: str) -> str`

Uploads a file to a specified bucket on the Minio server and returns the URI of the uploaded file.

```python
uri = minio.upload_file("my_bucket", "example.txt", "path/to/local/file.txt")
print(uri)
```

#### `upload_bytes(bucket_name: str, object_name: str, bytes_data: bytes) -> str`

Uploads bytes data to a specified bucket on the Minio server and returns the URI of the uploaded data.

```python
data = b"Hello, World!"
uri = minio.upload_bytes("my_bucket", "example_bytes", data)
print(uri)
```

#### `download_file(bucket_name: str, object_name: str, file_path: str) -> None`

Downloads a file from a specified bucket on the Minio server and saves it locally.

```python
minio.download_file("my_bucket", "example.txt", "path/to/download/file.txt")
```

#### `download_file_as_bytes(bucket_name: str, object_name: str) -> bytes`

Downloads a file from a specified bucket on the Minio server and returns it as bytes.

```python
data = minio.download_file_as_bytes("my_bucket", "example.txt")
print(data.decode())
```

#### `list_objects(bucket_name: str) -> List[str]`

Lists objects in a specified bucket on the Minio server.

```python
objects = minio.list_objects("my_bucket")
print(objects)
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-clients-resources-py-minio
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-clients-resources-py-minio --with dev
```