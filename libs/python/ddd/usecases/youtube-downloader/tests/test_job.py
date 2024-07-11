import asyncio
import unittest
from io import BytesIO
from unittest.mock import MagicMock, patch

from pyserializer.serializer import serialize_to_dict
from tests.reference_test import get_config, get_metadata
from youtube_downloader.job import Job, Output


class TestJob(unittest.TestCase):

    def setUp(self):
        self.config = get_config()
        self.metadata = get_metadata()
        self.debug = MagicMock()
        self.job = Job(config=self.config, metadata=self.metadata, dbg=self.debug)

    @patch('youtube_downloader.job.download_to_buffer')
    @patch('youtube_downloader.job.minio_client')
    def test_run_successful(self, mock_minio_client, mock_download_to_buffer):
        async def run_test():
            # Setup mock return values
            mock_video_data = BytesIO(b'test video data')
            mock_download_to_buffer.return_value = mock_video_data
            mock_minio = mock_minio_client.return_value
            mock_minio.upload_bytes.return_value = "http://minio.bucket/test-service/videos/video123/video.mp4"

            input_data = MagicMock()
            input_data.videoId = "video123"

            expected_output = Output(
                video_uri="http://minio.bucket/test-service/videos/video123/video.mp4",
                partition="test-service/videos/video123"
            )

            # Run the job
            feedback = await self.job.run(input_data)

            # Assert download_to_buffer is called with correct URL
            mock_download_to_buffer.assert_called_with(self.job.base_url.format("video123"))

            # Assert upload_bytes is called with correct parameters
            mock_minio.upload_bytes.assert_called_with(
                self.job.bucket_name,
                f"test-service/videos/video123/{self.job.file_name}",
                mock_video_data
            )

            # Check feedback is correct
            self.assertEqual(feedback.status.code, 200)
            self.assertEqual(feedback.status.detail, "Video uploaded successfully")
            self.assertEqual(feedback.data, serialize_to_dict(expected_output))
            self.assertEqual(feedback.metadata, self.metadata)

        asyncio.run(run_test())

    def test_get_video_id_from_event_order_valid(self):
        input_data = MagicMock()
        input_data.videoId = "video123"

        video_id = self.job.get_video_id(input_data)

        self.assertEqual(video_id, "video123")

    def test_get_video_id_from_event_order_invalid(self):
        input_data = MagicMock()
        input_data.videoId = None

        video_id = self.job.get_video_id(input_data)

        self.assertIsNone(video_id)

    @patch('youtube_downloader.job.logger')
    def test_download_video(self, mock_logger):
        with patch('youtube_downloader.job.download_to_buffer') as mock_download_to_buffer:
            mock_video_data = BytesIO(b'test video data')
            mock_download_to_buffer.return_value = mock_video_data

            self.job.video_id = "video123"
            video = self.job.download_video()

            self.assertEqual(video, mock_video_data)
            mock_download_to_buffer.assert_called_with(self.job.base_url.format("video123"))
            mock_logger.info.assert_called_with(f"endpoint: {self.job.base_url.format('video123')}")

    @patch('youtube_downloader.job.minio_client')
    def test_upload_video(self, mock_minio_client):
        mock_minio = mock_minio_client.return_value
        mock_minio.upload_bytes.return_value = "http://minio.bucket/video.mp4"

        video_data = BytesIO(b'test video data')
        uri = self.job.upload_video(video_data)

        self.assertEqual(uri, "http://minio.bucket/video.mp4")
        mock_minio.upload_bytes.assert_called_with(
            self.job.bucket_name,
            self.job.file_path,
            video_data
        )


if __name__ == '__main__':
    unittest.main()
