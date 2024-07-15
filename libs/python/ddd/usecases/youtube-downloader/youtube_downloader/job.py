from dataclasses import dataclass, field
from typing import Union

import warlock
from dto_config_vault.output import ConfigDTO
from dto_events_router.input import MetadataDTO, ServiceFeedBackDTO, StatusDTO
from pydebug import debug
from pylog.log import setup_logging
from pyminio.client import minio_client
from pyserializer.serializer import serialize_to_dict
from pyyoutube.client import download_to_buffer

logger = setup_logging(__name__)


@dataclass
class Output:
    """Dataclass representing the output of the video upload.

    Attributes:
        video_uri (str): The URI of the uploaded video.
        partition (str): The partition where the video is stored.
    """
    video_uri: str = field(metadata={"json": "videoUri"})
    partition: str = field(metadata={"json": "partition"})


class BaseJob:
    """Base class for job execution with debugging support.

    Args:
        dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debugging object.
    """
    def __init__(self, dbg: Union[debug.EnabledDebug, debug.DisabledDebug]) -> None:
        self.dbg = dbg

    def debug_response(self, file_name: str, response_body: bytes) -> None:
        """Saves the response for debugging purposes.

        Args:
            file_name (str): The name of the file to save the response.
            response_body (bytes): The response body to save.
        """
        self.dbg.save_response(file_name, response_body)


class Job(BaseJob):
    """Job class to handle video download and upload.

    Args:
        config (ConfigDTO): Configuration data transfer object.
        metadata (MetadataDTO): Metadata data transfer object.
        dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debugging object.
    """
    def __init__(
        self,
        config: ConfigDTO,
        metadata: MetadataDTO,
        dbg: Union[debug.EnabledDebug, debug.DisabledDebug]
    ) -> None:
        self.config = config
        self.metadata = metadata
        self.target_object = "video"
        self.file_extension = "mp4"
        self.base_url = "https://www.youtube.com/watch?v={}"
        self.video_id = None
        super().__init__(dbg)

    @property
    def config_id(self) -> str:
        """str: Returns the configuration ID."""
        return self.config.config_id

    @property
    def provider(self) -> str:
        """str: Returns the provider name."""
        return self.config.provider

    @property
    def service(self) -> str:
        """str: Returns the service name."""
        return self.config.service

    @property
    def source(self) -> str:
        """str: Returns the source name."""
        return self.config.source

    def __get_partition(self) -> str:
        """Generates the partition path.

        Returns:
            str: The partition path.
        """
        return f"{self.service}/videos/{self.video_id}"

    def __get_bucket_name(self) -> str:
        """Generates the bucket name.

        Returns:
            str: The bucket name.
        """
        return f"{self.provider}-{self.source}"

    def __valid_event_order(self, input_data) -> bool:
        """Validates the event order.

        Args:
            input_data: The input data to validate.

        Returns:
            bool: True if the event order is valid, False otherwise.
        """
        return input_data.videoId is not None

    def ___get_video_id_from_event_order(self, input_data) -> Union[str, None]:
        """Extracts the video ID from the event order.

        Args:
            input_data: The input data containing the event order.

        Returns:
            Union[str, None]: The video ID if valid, None otherwise.
        """
        if not self.__valid_event_order(input_data):
            return None
        return input_data.videoId

    @property
    def file_name(self) -> str:
        """str: Returns the file name with extension."""
        return f"{self.target_object}.{self.file_extension}"

    @property
    def partition(self) -> str:
        """str: Returns the partition path."""
        return self.__get_partition()

    @property
    def bucket_name(self) -> str:
        """str: Returns the bucket name."""
        return self.__get_bucket_name()

    @property
    def file_path(self) -> str:
        """str: Returns the file path."""
        return f"{self.partition}/{self.file_name}"

    def target_endpoint(self) -> str:
        """Generates the target endpoint URL for the video.

        Returns:
            str: The target endpoint URL.
        """
        return self.base_url.format(self.video_id)

    def get_status(self, code: str, detail: str) -> StatusDTO:
        """Generates a status object.

        Args:
            code (str): The status code.
            detail (str): The status detail.

        Returns:
            StatusDTO: The status data transfer object.
        """
        return StatusDTO(code=code, detail=detail)

    def download_video(self) -> bytes:
        """Downloads the video from the target endpoint.

        Returns:
            bytes: The downloaded video data.
        """
        logger.info(f"endpoint: {self.target_endpoint()}")
        return download_to_buffer(self.target_endpoint())

    def upload_video(self, video: bytes) -> str:
        """Uploads the video to MinIO and returns the URI.

        Args:
            video (bytes): The video data to upload.

        Returns:
            str: The URI of the uploaded video.
        """
        minio = minio_client()
        return minio.upload_bytes(self.bucket_name, self.file_path, video)

    def get_video_id(self, input_data: type[warlock.model.Model]) -> str:
        """Extracts the video ID from the input data.

        Args:
            input_data (type[warlock.model.Model]): The input data containing the event order.

        Returns:
            str: The extracted video ID.
        """
        return self.___get_video_id_from_event_order(input_data)

    async def run(self, input_data: type[warlock.model.Model]) -> ServiceFeedBackDTO:
        """Runs the job to download and upload the video.

        Args:
            input_data (type[warlock.model.Model]): The input data containing the event order.

        Returns:
            ServiceFeedBackDTO: The feedback data transfer object containing the result.
        """
        logger.info(f"Job triggered with input: {input_data}")
        logger.info(f"Job triggered with input type: {type(input_data)}")
        self.video_id = self.get_video_id(input_data)

        video = self.download_video()
        self.debug_response("video.mp4", video)
        uri = self.upload_video(video)
        status = self.get_status(200, "Video uploaded successfully")
        out = Output(video_uri=uri, partition=self.partition)
        return ServiceFeedBackDTO(
            data=serialize_to_dict(out),
            metadata=self.metadata,
            status=status
        )
