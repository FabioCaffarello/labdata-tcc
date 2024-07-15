import unittest
from unittest.mock import AsyncMock, patch
import respx
from httpx import Response
import os
from pylog.log import setup_logging
from config_loader.loader import ConfigLoader, fetch_configs
from tests.reference_test import get_config, get_configs
from pyserializer.serializer import serialize_to_dict

logger = setup_logging(__name__)


class TestConfigLoader(unittest.IsolatedAsyncioTestCase):

    @patch.dict(os.environ, {"CONFIG_VAULT_PORT_8000_TCP": "http://localhost:8001"})
    def setUp(self):
        self.service_name = "test-service-2"
        self.provider = "provider"
        self.config_loader = ConfigLoader()

    @respx.mock
    @patch('cli_config_vault.client.async_py_config_vault_client')
    async def test_fetch_configs_for_service(self, mock_client):
        mock_configs = get_configs()
        mock_client.return_value.list_configs_by_service_and_provider = AsyncMock(return_value=mock_configs)

        # Convert ConfigDTO objects to dictionaries
        mock_configs_dicts = [serialize_to_dict(config) for config in mock_configs]

        print(mock_configs_dicts)

        # Mock the actual HTTP call made by the client
        respx.get(
            f"http://localhost:8000/config/provider/{self.provider}/service/{self.service_name}" # noqa
        ).mock(
            return_value=Response(200, json=mock_configs_dicts)
        )

        result = await self.config_loader.fetch_configs_for_service(self.service_name, self.provider)
        print(result)

        expected = {
            config.config_id: config for config in mock_configs
        }
        self.assertEqual(result, expected)

    def test_register_config(self):
        config = get_config()
        self.config_loader.register_config("test_id", config)

        expected = {
            "test_id": config
        }
        self.assertEqual(self.config_loader.mapping_config, expected)

    def test_register_config_duplicate(self):
        config = get_config()
        self.config_loader.register_config("test_id", config)

        with self.assertRaises(ValueError):
            self.config_loader.register_config("test_id", config)


@patch.dict(os.environ, {"CONFIG_VAULT_PORT_8000_TCP": "http://localhost:8001"})
class TestFetchConfigs(unittest.IsolatedAsyncioTestCase):

    @patch('config_loader.loader.ConfigLoader.fetch_configs_for_service')
    async def test_fetch_configs(self, mock_fetch_configs_for_service):
        mock_config_loader = ConfigLoader()
        mock_config_loader.fetch_configs_for_service = AsyncMock(return_value={})
        mock_fetch_configs_for_service.return_value = {}

        result = await fetch_configs("test_service", "test_provider")

        self.assertEqual(result, {})
        mock_fetch_configs_for_service.assert_called_once_with("test_service", "test_provider")


if __name__ == '__main__':
    unittest.main()
