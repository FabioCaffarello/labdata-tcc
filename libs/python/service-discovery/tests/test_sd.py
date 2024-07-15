import unittest
from unittest.mock import patch
import os
from pysd.sd import ServiceDiscovery, ServiceUnavailableError, UnrecoverableError


class TestServiceDiscovery(unittest.TestCase):

    def setUp(self):
        self.envvars = {
            "RABBITMQ_PORT_6572_TCP": "tcp://rabbitmq:5672",
            "MINIO_PORT_9000_TCP": "tcp://minio:9000"
        }
        self.service_discovery = ServiceDiscovery(self.envvars)

    def test_initialization_with_valid_envvars(self):
        self.assertEqual(self.service_discovery._vars, self.envvars)

    def test_initialization_with_none_envvars(self):
        with self.assertRaises(UnrecoverableError):
            ServiceDiscovery(None)

    def test_rabbitmq_endpoint(self):
        expected_endpoint = "amqp://rabbitmq:5672"
        self.assertEqual(self.service_discovery.rabbitmq_endpoint, expected_endpoint)

    def test_minio_endpoint(self):
        expected_endpoint = "http://minio:9000"
        self.assertEqual(self.service_discovery.minio_endpoint, expected_endpoint)

    @patch.dict(os.environ, {"RABBITMQ_GATEWAY_HOST": "gateway_host"})
    def test_get_gateway_host_with_env_var(self):
        self.assertEqual(self.service_discovery._get_gateway_host("RABBITMQ"), "gateway_host")

    @patch.dict(os.environ, {}, clear=True)
    def test_get_gateway_host_without_env_var(self):
        self.assertEqual(self.service_discovery._get_gateway_host("RABBITMQ"), "localhost")

    @patch.dict(os.environ, {"MINIO_ACCESS_KEY": "minio_key", "MINIO_SECRET_KEY": "minio_secret"})
    def test_minio_access_key(self):
        self.assertEqual(self.service_discovery.minio_access_key, "minio_key")

    @patch.dict(os.environ, {"MINIO_ACCESS_KEY": "minio_key", "MINIO_SECRET_KEY": "minio_secret"})
    def test_minio_secret_key(self):
        self.assertEqual(self.service_discovery.minio_secret_key, "minio_secret")

    def test_service_unavailable_error(self):
        with self.assertRaises(ServiceUnavailableError):
            self.service_discovery._get_endpoint("NON_EXISTENT_VAR", "NON_EXISTENT_SERVICE")

    def test_services_rabbitmq_exchange(self):
        self.assertEqual(self.service_discovery.services_rabbitmq_exchange, "services")

    def test_modify_localhost_port(self):
        endpoint = "http://localhost:8000"
        modified_endpoint = self.service_discovery._modify_localhost_port(endpoint, "8000", "8001")
        self.assertEqual(modified_endpoint, "http://localhost:8001")


if __name__ == "__main__":
    unittest.main()
