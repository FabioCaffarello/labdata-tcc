from unittest.mock import AsyncMock, patch
import unittest
import respx
import httpx
from cli_config_vault.client import async_py_config_vault_client, AsyncPyConfigVaultClient
from dto_config_vault.output import ConfigDTO
from typing import Dict
from tests.reference_test import get_config, get_configs


class TestAsyncPyConfigHandlerClient(unittest.IsolatedAsyncioTestCase):

    def setUp(self):
        patcher = patch('pysd.sd.ServiceDiscovery.services_config_vault_endpoint', new_callable=AsyncMock)
        self.mock_endpoint = patcher.start()
        self.mock_endpoint.return_value = "http://localhost:8001"
        self.addCleanup(patcher.stop)

        self.client = AsyncPyConfigVaultClient(base_url=self.mock_endpoint.return_value)
        self.mock_response = lambda data: httpx.Response(200, json=data)
        self.client.configs_endpoint = "/configs"

    def mock_response(self, data: Dict):
        return httpx.Response(200, json=data)

    @respx.mock
    async def test_create_config(self):
        data = get_config()
        endpoint = f"{self.client.client.base_url}{self.client.configs_endpoint}"
        respx.post(endpoint).mock(return_value=self.mock_response(data))

        result = await self.client.create_config(data)
        self.assertIsInstance(result, ConfigDTO)
        self.assertEqual(result.id, data["_id"])

    @respx.mock
    async def test_update_config(self):
        data = get_config()
        endpoint = f"{self.client.client.base_url}{self.client.configs_endpoint}"
        respx.put(endpoint).mock(return_value=self.mock_response(data))

        result = await self.client.update_config(data)
        self.assertIsInstance(result, ConfigDTO)
        self.assertEqual(result.id, data["_id"])

    @respx.mock
    async def test_list_all_configs(self):
        configs_data = get_configs()
        endpoint = f"{self.client.client.base_url}{self.client.configs_endpoint}"
        respx.get(endpoint).mock(return_value=self.mock_response(configs_data))

        result = await self.client.list_all_configs()
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(config, ConfigDTO) for config in result))

    @respx.mock
    async def test_get_config_by_id(self):
        config_id = "123"
        config_data = get_config(config_id=config_id)
        endpoint = f"{self.client.client.base_url}{self.client.configs_endpoint}/{config_id}"
        respx.get(endpoint).mock(return_value=self.mock_response(config_data))

        result = await self.client.get_config_by_id(config_id)
        self.assertIsInstance(result, ConfigDTO)
        self.assertEqual(result.id, config_data["_id"])

    @respx.mock
    async def test_delete_config(self):
        config_id = "123"
        endpoint = f"{self.client.client.base_url}{self.client.configs_endpoint}/{config_id}"
        respx.delete(endpoint).mock(return_value=httpx.Response(204))

        await self.client.delete_config(config_id)

    @respx.mock
    async def test_list_configs_by_service_and_provider(self):
        provider = "provider"
        service = "test-service"
        configs_data = get_configs()
        endpoint = f"{self.client.client.base_url}{self.client.configs_endpoint}/provider/{provider}/service/{service}"
        respx.get(endpoint).mock(return_value=self.mock_response(configs_data))

        result = await self.client.list_configs_by_service_and_provider(provider, service)
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(config, ConfigDTO) for config in result))

    @respx.mock
    async def test_list_configs_by_source_and_provider(self):
        provider = "provider"
        source = "test-source"
        configs_data = get_configs()
        endpoint = f"{self.client.client.base_url}{self.client.configs_endpoint}/provider/{provider}/source/{source}"
        respx.get(endpoint).mock(return_value=self.mock_response(configs_data))

        result = await self.client.list_configs_by_source_and_provider(provider, source)
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(config, ConfigDTO) for config in result))

    @respx.mock
    async def test_list_configs_by_service_provider_and_active(self):
        provider = "provider"
        service = "test-service"
        active = True
        configs_data = get_configs()
        endpoint = (
            f"{self.client.client.base_url}{self.client.configs_endpoint}"
            f"/provider/{provider}"
            f"/service/{service}"
            f"/active/{active}"
        )
        respx.get(endpoint).mock(return_value=self.mock_response(configs_data))

        result = await self.client.list_configs_by_service_provider_and_active(provider, service, active)
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(config, ConfigDTO) for config in result))

    @respx.mock
    async def test_list_configs_by_service_source_and_provider(self):
        provider = "provider"
        service = "test-service"
        source = "test-source"
        configs_data = get_configs()
        endpoint = (
            f"{self.client.client.base_url}{self.client.configs_endpoint}"
            f"/provider/{provider}"
            f"/service/{service}"
            f"/source/{source}"
        )
        respx.get(endpoint).mock(return_value=self.mock_response(configs_data))

        result = await self.client.list_configs_by_service_source_and_provider(provider, service, source)
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(config, ConfigDTO) for config in result))

    @respx.mock
    async def test_list_configs_by_provider_and_dependencies(self):
        provider = "provider"
        service = "test-service"
        source = "test-source"
        configs_data = get_configs()
        endpoint = (
            f"{self.client.client.base_url}{self.client.configs_endpoint}"
            f"/provider/{provider}"
            "/dependencies"
            f"/service/{service}"
            f"/source/{source}"
        )
        respx.get(endpoint).mock(return_value=self.mock_response(configs_data))

        result = await self.client.list_configs_by_provider_and_dependencies(provider, service, source)
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(config, ConfigDTO) for config in result))

    @respx.mock
    async def test_async_py_config_vault_client(self):
        client = async_py_config_vault_client()
        self.assertIsInstance(client, AsyncPyConfigVaultClient)


if __name__ == "__main__":
    unittest.main()
