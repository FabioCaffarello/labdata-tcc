import unittest
from unittest.mock import patch, MagicMock
from pyminio.client import MinioClient
from minio.error import S3Error
from pylog import log

logger = log.setup_logging(__name__)


class TestMinioClient(unittest.TestCase):

    @patch('pyminio.client.Minio')
    @patch('pyminio.client.new_from_env')
    def setUp(self, MockNewFromEnv, MockMinio):
        self.mock_minio = MockMinio.return_value
        mock_sd = MockNewFromEnv.return_value
        mock_sd.minio_endpoint.return_value = "minio.example.com"
        mock_sd.minio_access_key.return_value = "your_access_key"
        mock_sd.minio_secret_key.return_value = "your_secret_key"
        self.client = MinioClient(
            endpoint="http://minio.example.com",
            access_key="your_access_key",
            secret_key="your_secret_key"
        )

    def test_create_bucket(self):
        self.client.create_bucket("test_bucket")
        self.mock_minio.make_bucket.assert_called_with("test_bucket")

    def test_create_bucket_exception(self):
        self.mock_minio.make_bucket.side_effect = S3Error(
            "Error", "BucketAlreadyOwnedByYou", "resource", "request_id", "host_id", "response"
        )
        with self.assertRaises(Exception) as context:
            self.client.create_bucket("test_bucket")
        self.assertIn("BucketAlreadyOwnedByYou", str(context.exception))

    def test_list_buckets(self):
        mock_bucket1 = MagicMock()
        mock_bucket1.name = "bucket1"
        mock_bucket2 = MagicMock()
        mock_bucket2.name = "bucket2"
        self.mock_minio.list_buckets.return_value = [mock_bucket1, mock_bucket2]
        buckets = self.client.list_buckets()
        self.assertEqual(buckets, ["bucket1", "bucket2"])

    def test_list_buckets_exception(self):
        self.mock_minio.list_buckets.side_effect = S3Error(
            "Error", "AccessDenied", "resource", "request_id", "host_id", "response"
        )
        with self.assertRaises(Exception) as context:
            self.client.list_buckets()
        self.assertIn("AccessDenied", str(context.exception))

    def test_upload_file(self):
        self.client.upload_file("test_bucket", "test_object", "test_path")
        self.mock_minio.fput_object.assert_called_with("test_bucket", "test_object", "test_path")

    def test_upload_file_exception(self):
        self.mock_minio.fput_object.side_effect = S3Error(
            "Error", "NoSuchBucket", "resource", "request_id", "host_id", "response"
        )
        with self.assertRaises(Exception) as context:
            self.client.upload_file("test_bucket", "test_object", "test_path")
        self.assertIn("NoSuchBucket", str(context.exception))

    def test_upload_bytes(self):
        self.client.upload_bytes("test_bucket", "test_object", b"test_data")
        self.mock_minio.put_object.assert_called()

    def test_upload_bytes_exception(self):
        self.mock_minio.put_object.side_effect = S3Error(
            "Error", "NoSuchBucket", "resource", "request_id", "host_id", "response"
        )
        with self.assertRaises(Exception) as context:
            self.client.upload_bytes("test_bucket", "test_object", b"test_data")
        self.assertIn("NoSuchBucket", str(context.exception))

    def test_download_file(self):
        self.client.download_file("test_bucket", "test_object", "test_path")
        self.mock_minio.fget_object.assert_called_with("test_bucket", "test_object", "test_path")

    def test_download_file_exception(self):
        self.mock_minio.fget_object.side_effect = S3Error(
            "Error", "NoSuchKey", "resource", "request_id", "host_id", "response"
        )
        with self.assertRaises(Exception) as context:
            self.client.download_file("test_bucket", "test_object", "test_path")
        self.assertIn("NoSuchKey", str(context.exception))

    def test_download_file_as_bytes(self):
        mock_response = MagicMock()
        mock_response.read.return_value = b"test_data"
        self.mock_minio.get_object.return_value = mock_response
        data = self.client.download_file_as_bytes("test_bucket", "test_object")
        self.assertEqual(data, b"test_data")

    def test_download_file_as_bytes_exception(self):
        self.mock_minio.get_object.side_effect = S3Error(
            "Error", "NoSuchKey", "resource", "request_id", "host_id", "response"
        )
        with self.assertRaises(Exception) as context:
            self.client.download_file_as_bytes("test_bucket", "test_object")
        self.assertIn("NoSuchKey", str(context.exception))

    def test_list_objects(self):
        self.mock_minio.list_objects.return_value = [
            MagicMock(object_name="object1"), MagicMock(object_name="object2")
        ]
        objects = self.client.list_objects("test_bucket")
        self.assertEqual(objects, ["object1", "object2"])

    def test_list_objects_exception(self):
        self.mock_minio.list_objects.side_effect = S3Error(
            "Error", "NoSuchBucket", "resource", "request_id", "host_id", "response"
        )
        with self.assertRaises(Exception) as context:
            self.client.list_objects("test_bucket")
        self.assertIn("NoSuchBucket", str(context.exception))


if __name__ == "__main__":
    unittest.main()
