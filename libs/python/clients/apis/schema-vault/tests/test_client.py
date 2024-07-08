import unittest
from unittest.mock import patch
import respx
import httpx
from typing import Dict
from dto_schema_vault.output import SchemaDTO
from cli_schema_vault.client import AsyncPySchemaVaultClient
from tests.reference_test import get_schema, get_schemas


class TestAsyncPySchemaVaultClient(unittest.IsolatedAsyncioTestCase):

    def setUp(self):
        patcher = patch('pysd.sd.ServiceDiscovery.services_schema_vault_endpoint', return_value="http://localhost:8002")
        self.mock_endpoint = patcher.start()
        self.mock_endpoint.return_value = "http://localhost:8002"
        self.addCleanup(patcher.stop)

        self.client = AsyncPySchemaVaultClient(base_url=self.mock_endpoint.return_value)
        self.mock_response = lambda data: httpx.Response(200, json=data)
        self.client.configs_endpoint = "/schemas"

    def mock_response(self, data: Dict):
        return httpx.Response(httpx.codes.OK, json=data)

    @respx.mock
    async def test_create_schema(self):
        data = get_schema()
        endpoint = f"{self.client.client.base_url}{self.client.schemas_endpoint}"
        respx.post(endpoint).mock(return_value=self.mock_response(data))

        result = await self.client.create_schema(data)
        self.assertIsInstance(result, SchemaDTO)
        self.assertEqual(result.id, data["_id"])

    @respx.mock
    async def test_update_schema(self):
        data = get_schema()
        endpoint = f"{self.client.client.base_url}{self.client.schemas_endpoint}"
        respx.put(endpoint).mock(return_value=self.mock_response(data))

        result = await self.client.update_schema(data)
        self.assertIsInstance(result, SchemaDTO)
        self.assertEqual(result.id, data["_id"])

    @respx.mock
    async def test_list_all_schemas(self):
        schemas_data = get_schemas()
        endpoint = f"{self.client.client.base_url}{self.client.schemas_endpoint}"
        respx.get(endpoint).mock(return_value=self.mock_response(schemas_data))

        result = await self.client.list_all_schemas()
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(schema, SchemaDTO) for schema in result))

    @respx.mock
    async def test_get_schema_by_id(self):
        schema_id = "schema-id"
        schema_data = get_schema(schema_id=schema_id)
        endpoint = f"{self.client.client.base_url}{self.client.schemas_endpoint}/{schema_id}"
        respx.get(endpoint).mock(return_value=self.mock_response(schema_data))

        result = await self.client.get_schema_by_id(schema_id)
        self.assertIsInstance(result, SchemaDTO)
        self.assertEqual(result.id, schema_data["_id"])

    @respx.mock
    async def test_delete_schema(self):
        schema_id = "schema-id"
        endpoint = f"{self.client.client.base_url}{self.client.schemas_endpoint}/{schema_id}"
        respx.delete(endpoint).mock(return_value=httpx.Response(204))

        await self.client.delete_schema(schema_id)

    @respx.mock
    async def test_list_schemas_by_service_and_provider(self):
        provider = "provider"
        service = "test-service"
        schemas_data = get_schemas()
        endpoint = f"{self.client.client.base_url}{self.client.schemas_endpoint}/provider/{provider}/service/{service}"
        respx.get(endpoint).mock(return_value=self.mock_response(schemas_data))

        result = await self.client.list_schemas_by_service_and_provider(provider, service)
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(schema, SchemaDTO) for schema in result))

    @respx.mock
    async def test_list_schemas_by_source_and_provider(self):
        provider = "provider"
        source = "test-source"
        schemas_data = get_schemas()
        endpoint = f"{self.client.client.base_url}{self.client.schemas_endpoint}/provider/{provider}/source/{source}"
        respx.get(endpoint).mock(return_value=self.mock_response(schemas_data))

        result = await self.client.list_schemas_by_source_and_provider(provider, source)
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(schema, SchemaDTO) for schema in result))

    @respx.mock
    async def test_list_schemas_by_service_source_and_provider(self):
        provider = "provider"
        service = "test-service"
        source = "test-source"
        schemas_data = get_schemas()
        endpoint = (
            f"{self.client.client.base_url}{self.client.schemas_endpoint}"
            f"/provider/{provider}"
            f"/service/{service}"
            f"/source/{source}"
        )
        respx.get(endpoint).mock(return_value=self.mock_response(schemas_data))

        result = await self.client.list_schemas_by_service_source_and_provider(provider, service, source)
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(schema, SchemaDTO) for schema in result))

    @respx.mock
    async def test_list_schemas_by_service_source_provider_and_schema_type(self):
        provider = "provider"
        service = "test-service"
        source = "test-source"
        schema_type = "type"
        schemas_data = get_schemas()
        endpoint = (
            f"{self.client.client.base_url}{self.client.schemas_endpoint}"
            f"/provider/{provider}"
            f"/service/{service}"
            f"/source/{source}"
            f"/schema-type/{schema_type}"
        )

        respx.get(endpoint).mock(return_value=self.mock_response(schemas_data))

        result = await self.client.list_schemas_by_service_source_provider_and_schema_type(
            provider,
            service,
            source,
            schema_type
        )
        self.assertIsInstance(result, list)
        self.assertTrue(all(isinstance(schema, SchemaDTO) for schema in result))

    @respx.mock
    async def test_validate_schema(self):
        data = get_schema()
        endpoint = f"{self.client.client.base_url}{self.client.schemas_endpoint}/validate"
        respx.post(endpoint).mock(return_value=self.mock_response({"valid": True}))

        result = await self.client.validate_schema(data)
        self.assertTrue(result)


if __name__ == "__main__":
    unittest.main()
