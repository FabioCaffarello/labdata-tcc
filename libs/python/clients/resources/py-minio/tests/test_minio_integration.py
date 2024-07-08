# import os
# import time
# import unittest
# from minio import Minio, S3Error
# from pyminio.client import MinioClient


# class MinioIntegrationTest(unittest.TestCase):

#     @classmethod
#     def setUpClass(cls):
#         cls.minio_client = Minio(
#             endpoint="localhost:9003",
#             access_key="minio-root-user",
#             secret_key="minio-root-password",
#             secure=False,
#         )
#         cls.client = MinioClient(
#             endpoint="localhost:9003",
#             access_key="minio-root-user",
#             secret_key="minio-root-password"
#         )
#         cls.bucket_name = "test-bucket"
#         cls.object_name = "test_file.txt"
#         cls.file_path = "test_file.txt"
#         cls.download_path = "downloaded_test_file.txt"
#         cls.wait_for_minio(cls.minio_client)

#         # Ensure the bucket exists before running tests
#         if not cls.minio_client.bucket_exists(cls.bucket_name):
#             cls.client.create_bucket(cls.bucket_name)

#     @classmethod
#     def tearDownClass(cls):
#         # Clean up after tests
#         if os.path.exists(cls.file_path):
#             os.remove(cls.file_path)
#         if os.path.exists(cls.download_path):
#             os.remove(cls.download_path)

#         # Remove all objects in the bucket before deleting it
#         try:
#             objects = cls.minio_client.list_objects(cls.bucket_name, recursive=True)
#             for obj in objects:
#                 cls.minio_client.remove_object(cls.bucket_name, obj.object_name)
#         except S3Error as err:
#             print(f"Error removing objects: {err}")

#         try:
#             cls.minio_client.remove_bucket(cls.bucket_name)
#         except S3Error as err:
#             print(f"Error removing bucket: {err}")

#     @staticmethod
#     def wait_for_minio(client, timeout=60):
#         """Wait for Minio service to be up and running."""
#         start = time.time()
#         while time.time() - start < timeout:
#             try:
#                 if client.bucket_exists("test-bucket"):
#                     return
#             except Exception:
#                 time.sleep(2)
#         raise Exception("Minio did not start within the given time")

#     def test_create_bucket(self):
#         bucket_name = "test-bucket"
#         if not self.minio_client.bucket_exists(bucket_name):
#             self.client.create_bucket(bucket_name)
#         buckets = self.client.list_buckets()
#         self.assertIn(bucket_name, buckets)

#     def test_upload_file(self):
#         bucket_name = "test-bucket"
#         file_path = "test_file.txt"
#         with open(file_path, "w") as f:
#             f.write("This is a test file.")
#         object_name = "test_file.txt"
#         self.client.upload_file(bucket_name, object_name, file_path)
#         objects = self.client.list_objects(bucket_name)
#         self.assertIn(object_name, objects)

#     def test_download_file(self):
#         bucket_name = "test-bucket"
#         object_name = "test_file.txt"
#         download_path = "downloaded_test_file.txt"
#         self.client.download_file(bucket_name, object_name, download_path)
#         self.assertTrue(os.path.exists(download_path))
#         with open(download_path, "r") as f:
#             content = f.read()
#         self.assertEqual(content, "This is a test file.")

# # FIXME: This test is failing because the Minio server is not running
# # if __name__ == "__main__":
# #     unittest.main()
