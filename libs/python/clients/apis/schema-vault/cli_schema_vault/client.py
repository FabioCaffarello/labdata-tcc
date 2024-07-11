from typing import List, Dict, Any
from pyrequest.async_factory import RateLimitedAsyncHttpClient
from pysd.sd import new_from_env
from dto_schema_vault.output import SchemaDTO
from dto_schema_vault.input import SchemaDataDTO
from pyserializer.serializer import serialize_to_dataclass, serialize_to_dict


class AsyncPySchemaVaultClient:
    """
    A client for handling asynchronous interactions with a schema service.

    Args:
        base_url (str): The base URL of the schema service.

    Attributes:
        max_calls (int): The maximum number of API calls allowed in a period.
        period (int): The time period (in seconds) in which API calls are rate-limited.
        client (RateLimitedAsyncHttpClient): An instance of the rate-limited HTTP client.
        schemas_endpoint (str): The endpoint for schema operations.
    """

    def __init__(self, base_url: str):
        self.max_calls = 100
        self.period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.max_calls, self.period)
        self.schemas_endpoint = "/schema"

    async def create(self, data: Dict[str, Any]) -> SchemaDTO:
        """
        Create a new schema using the provided data.

        Args:
            data (Dict[str, Any]): The schema data to be created.

        Returns:
            SchemaDTO: The created schema in the form of a data class.
        """
        schema_data = await self.client.make_request("POST", self.schemas_endpoint, data=data)
        return serialize_to_dataclass(schema_data, SchemaDTO)

    async def update_schema(self, data: Dict[str, Any]) -> SchemaDTO:
        """
        Update an existing schema using the provided data.

        Args:
            data (Dict[str, Any]): The schema data to be updated.

        Returns:
            SchemaDTO: The updated schema in the form of a data class.
        """
        schema_data = await self.client.make_request("PUT", self.schemas_endpoint, data=data)
        return serialize_to_dataclass(schema_data, SchemaDTO)

    async def list_all_schemas(self) -> List[SchemaDTO]:
        """
        Retrieve a list of all schemas.

        Returns:
            List[SchemaDTO]: A list of all schemas in the form of data classes.
        """
        schemas_data = await self.client.make_request("GET", self.schemas_endpoint)
        return [serialize_to_dataclass(schema, SchemaDTO) for schema in schemas_data]

    async def get_schema_by_id(self, schema_id: str) -> SchemaDTO:
        """
        Retrieve a specific schema by its ID.

        Args:
            schema_id (str): The unique identifier of the schema.

        Returns:
            SchemaDTO: The requested schema in the form of a data class.
        """
        endpoint = f"{self.schemas_endpoint}/{schema_id}"
        schema_data = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(schema_data, SchemaDTO)

    async def delete_schema(self, schema_id: str) -> None:
        """
        Delete a specific schema by its ID.

        Args:
            schema_id (str): The unique identifier of the schema.

        Returns:
            None
        """
        endpoint = f"{self.schemas_endpoint}/{schema_id}"
        await self.client.make_request("DELETE", endpoint)

    async def list_schemas_by_service_and_provider(self, provider: str, service: str) -> List[SchemaDTO]:
        """
        Retrieve a list of schemas associated with a specific service and provider.

        Args:
            provider (str): The name of the provider.
            service (str): The name of the service.

        Returns:
            List[SchemaDTO]: A list of schemas for the specified service and provider in the form of data classes.
        """
        endpoint = f"{self.schemas_endpoint}/provider/{provider}/service/{service}"
        schemas_data = await self.client.make_request("GET", endpoint)
        return [serialize_to_dataclass(schema, SchemaDTO) for schema in schemas_data]

    async def list_schemas_by_source_and_provider(self, provider: str, source: str) -> List[SchemaDTO]:
        """
        Retrieve a list of schemas associated with a specific source and provider.

        Args:
            provider (str): The name of the provider.
            source (str): The name of the source.

        Returns:
            List[SchemaDTO]: A list of schemas for the specified source and provider in the form of
            data classes.
        """
        endpoint = f"{self.schemas_endpoint}/provider/{provider}/source/{source}"
        schemas_data = await self.client.make_request("GET", endpoint)
        return [serialize_to_dataclass(schema, SchemaDTO) for schema in schemas_data]

    async def list_schemas_by_service_source_and_provider(
        self,
        provider: str,
        service: str,
        source: str
    ) -> List[SchemaDTO]:
        """
        Retrieve a list of schemas associated with a specific service, source, and provider.

        Args:
            provider (str): The name of the provider.
            service (str): The name of the service.
            source (str): The name of the source.

        Returns:
            List[SchemaDTO]: A list of schemas for the specified service, source, and provider in the
            form of data classes.
        """
        endpoint = f"{self.schemas_endpoint}/provider/{provider}/service/{service}/source/{source}"
        schemas_data = await self.client.make_request("GET", endpoint)
        return [serialize_to_dataclass(schema, SchemaDTO) for schema in schemas_data]

    async def list_schemas_by_service_source_provider_and_schema_type(
        self,
        provider: str,
        service: str,
        source: str,
        schema_type: str
    ) -> SchemaDTO:
        """
        Retrieve a list of schemas associated with a specific service, source, provider, and schema type.

        Args:
            provider (str): The name of the provider.
            service (str): The name of the service.
            source (str): The name of the source.
            schema_type (str): The type of the schema.

        Returns:
            SchemaDTO: A schemas for the specified service, source, provider, and schema type in the
            form of data classes.
        """
        endpoint = (
            f"{self.schemas_endpoint}"
            f"/provider/{provider}"
            f"/service/{service}"
            f"/source/{source}"
            f"/schema-type/{schema_type}"
        )
        schema_data = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(schema_data, SchemaDTO)

    async def validate_schema(self, data: SchemaDataDTO) -> bool:
        """
        Validate a schema using the provided data.

        Args:
            data (Dict[str, Any]): The schema data to be validated.

        Returns:
            bool: True if the schema is valid, False otherwise.
        """
        endpoint = f"{self.schemas_endpoint}/validate"
        form_data = serialize_to_dict(data)
        validation_response = await self.client.make_request("POST", endpoint, data=form_data)
        return validation_response.get("valid", False)


def async_py_schema_vault_client() -> AsyncPySchemaVaultClient:
    """
    Create an instance of the AsyncPySchemaVaultClient using service discovery information.

    Returns:
        AsyncPySchemaVaultClient: An instance of the schema vault client.
    """
    sd = new_from_env()
    return AsyncPySchemaVaultClient(sd.services_schema_vault_endpoint)
