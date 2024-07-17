# Architecture

## Overview

The architecture of the LabData TCC project is designed to be modular and scalable, consisting of shared services and operation services. Shared services provide common functionality used across the system, while operation services handle specific tasks related to the project's goals.

## Shared Services

### Configuration API (Config Vault)

**Purpose**: Manages and provides configuration settings required by other services to function correctly.  
**Technologies**: Golang  
**Key Features**:
- Centralized configuration management.
- Dynamic configuration updates.

**Example Configuration**:
```json
{
    "active": true,
    "service": "video-downloader",
    "source": "pinkfong",
    "provider": "kids",
    "job_parameters": {
        "parser_module": "youtube_downloader"
    },
    "depends_on": [{
        "service": "file-watcher",
        "source": "pinkfong"
    }]
}
```

### Schema API (Schema Vault)

**Purpose**: Handles the schemas for input and output data, ensuring data consistency and validation.  
**Technologies**: Golang  
**Key Features**:
- Schema management for different jobs.

**Example Input Schema**:
```json
{
    "service": "video-downloader",
    "source": "pinkfong",
    "provider": "kids",
    "schema_type": "input",
    "json_schema": {
        "type": "object",
        "properties": {
            "videoId": {
                "type": "string"
            }
        },
        "required": ["videoId"]
    }
}
```

**Example Output Schema**:
```json
{
    "service": "video-downloader",
    "source": "pinkfong",
    "provider": "kids",
    "schema_type": "output",
    "json_schema": {
        "type": "object",
        "properties": {
            "videoUri": {
                "type": "string"
            },
            "partition": {
                "type": "string"
            }
        },
        "required": ["videoUri", "partition"]
    }
}
```

### Input API

**Purpose**: Manages all input data and triggers jobs by publishing messages to RabbitMQ.  
**Technologies**: Golang
**Key Features**:
- Input data handling.
- Message publishing to RabbitMQ.

**Example Input**:
```json
{
    "service": "video-downloader",
    "source": "pinkfong",
    "provider": "kids",
    "data": {
        "videoId": "XqZsoesa55w"
    }
}
```


### Output API (Output Vault)

**Purpose**: Manages output data, including storage and retrieval.  
**Technologies**: Golang
**Key Features**:
- Centralized output management.
- Scalable storage solutions.

**Example Output**:
```json
{
    "service": "video-downloader",
    "source": "pinkfong",
    "provider": "kids",
    "data": {
        "uri": "http://minio:9000/kids-pinkfong/video-downloader/videos/XqZsoesa55w/video.mp4",
        "partition": "video-downloader/videos/XqZsoesa55w"
    },
    "metadata": {
        "input_id": "da853c5826b798b82320e42024d97837",
        "input": {
            "data": {
                "videoId": "XqZsoesa55w"
            },
            "processing_id": "aae4b46a-855f-449b-aa0a-f595241b6d8d",
            "processing_timestamp": "2024-06-24 08:30:00"
        }
    }
}
```

### Data Lineage (Processing Lineage)
**Purpose**:  Manages and controls job lineage to ensure transparency and traceability of data processing workflows.
**Technologies**: Golang 
**Key Features**:
- Centralized processing management.
- Tracking of job execution history.
- Ensuring traceability and transparency in data processing workflows.

**Example Processing Lineage**:
```json
{
    "parent_processing_id": "1591d02f-3bd6-4c94-94e1-4646210e2c89",
    "tasks": [
        {
            "source": "pinkfong",
            "service": "file-watcher",
            "provider": "kids",
            "processing_id": "1591d02f-3bd6-4c94-94e1-4646210e2c89",
            "parent_processing_id": "1591d02f-3bd6-4c94-94e1-4646210e2c89",
            "status_code": 200,
            "configs": {
                "config_id": "ffabbe29f917c95197cd41fadc9c1635",
                "config_version_id": "be188d85-e856-5e8c-be10-9f6a38cf2eae",
                "schemas": [
                    {
                        "version": "f6992dbd-6c40-4130-8c1d-b8084b044c67",
                        "schema_type": "input",
                        "schemaID": "9a77d331df7c3bc78b34be01e3c11739"
                    },{
                        "version": "722f9b83-2e7d-47fa-a444-025e8486c542",
                        "schema_type": "output",
                        "schemaID": "d627223805943b9336298f6c9edb7716"
                    }
                ]
            },
            "input_id": "fca1ced9e6efb2252ca056988d2f4495",
            "output_id": "3ef702427d2e223f401c55c43db4f717",
            "processing_timestamp": "2024-01-10T01:53:12Z"
        },
        {
            "source": "pinkfong",
            "service": "video-watcher",
            "provider": "kids",
            "processing_id": "5064fb9a-67ed-44d9-9a21-85d77cbada0e",
            "parent_processing_id": "1591d02f-3bd6-4c94-94e1-4646210e2c89",
            "status_code": 200,
            "configs": {
                "config_id": "8ebfb00547d519fc36a2f4bb07188b36",
                "config_version_id": "9ac394fd-a50c-431c-999b-3bd63d548f51",
                "schemas": [
                    {
                        "version": "aa26e22a-e61e-48cd-95ae-32fd35006a16",
                        "schema_type": "input",
                        "schemaID": "3b99a6260cf0762e9a0de53bfb5b7bc5"
                    },{
                        "version": "664894cd-a6f0-40c0-be51-876508264c5b",
                        "schema_type": "output",
                        "schemaID": "178b6e9c84e665439c80544ca34e490e"
                    }
                ]
            },
            "input_id": "0d4918689ee5225da72d60fbbc7d2f06",
            "output_id": "ad43547e5896519be80e175c0210da20",
            "processing_timestamp": "2024-07-10T01:53:12Z"
        },
        {
            "source": "pinkfong",
            "service": "video-watcher",
            "provider": "kids",
            "processing_id": "60d55a89-efcf-4a7f-bc7a-bab28ad3a474",
            "parent_processing_id": "1591d02f-3bd6-4c94-94e1-4646210e2c89",
            "status_code": 200,
            "configs": {
                "config_id": "8ebfb00547d519fc36a2f4bb07188b36",
                "config_version_id": "9ac394fd-a50c-431c-999b-3bd63d548f51",
                "schemas": [
                    {
                        "version": "aa26e22a-e61e-48cd-95ae-32fd35006a16",
                        "schema_type": "input",
                        "schemaID": "3b99a6260cf0762e9a0de53bfb5b7bc5"
                    },{
                        "version": "664894cd-a6f0-40c0-be51-876508264c5b",
                        "schema_type": "output",
                        "schemaID": "178b6e9c84e665439c80544ca34e490e"
                    }
                ]
            },
            "input_id": "1195c9bbf8d73025a0c33298bc59f5be",
            "output_id": "bc22dbca359dacb500948e0573771c10",
            "processing_timestamp": "2024-07-10T01:55:57Z"
        },
    ]
}
```

### Orchestrator Service (Events Router)

**Purpose**: Listens to messages from RabbitMQ and orchestrates the sequence of job executions by triggering appropriate operation services.  
**Technologies**: Golang  
**Key Features**:
- Job orchestration
    - Pre processing Input:
        - Validate Input with the properly schema.
        - Create an unit operation job at the processing lineage API.
        - Dispatch an Input to the properly service.
    - Processing:
        - Update Input status
    - Service Feedback
        - Save the output
        - Update Input Status
        - Triggered jobs that depends on the jobs that finished (if success status)
        - Creates the data processing lineage.

## Operation Services

### Video Downloader Service (Video downloader)

**Purpose**: Downloads videos from YouTube and saves MP4 files in a storage bucket.  
**Technologies**: Python 
**Key Features**:
- Video downloading
- Storage management

### Audio Converter Service (Audio Provider)

**Purpose**: Converts MP4 files to MP3 format.  
**Technologies**: Python  
**Key Features**:
- Audio conversion
- Format handling

### Text Extractor Service  (Speech Recognizer)

**Purpose**: Extracts text from MP3 files using LLM models.  
**Technologies**: Python
**Key Features**:
- Text extraction
- Audio processing

### RAG Service (Rag Interpreter)

**Purpose**: Applies Retrieval-Augmented Generation (RAG) to preprocess text data before embedding creation.  
**Technologies**: Python
**Key Features**:
- Text preprocessing
- Enhancement of data quality for embeddings


### Knowledge Graph Creation Service (Graph Forge)

**Purpose**: Utilizes RAG-processed text to create and manage a knowledge graph in Neo4J. This service is essential for establishing relationships between different entities derived from processed data and enabling complex querying capabilities.
**Technologies**: Python  
**Key Features**:
- Knowledge graph creation: Constructs the knowledge graph using text data processed by the RAG service.
- Relationship mapping:Establishes connections between entities based on semantic relationships identified in the RAG-processed text.

### Front-end
**Purpose**: Designed to provide an intuitive and interactive user interface for users to interact with the knowledge graph semantically.
**Technologies**: React.js 
**Key Features**:
- Prompt to execute semantic queries