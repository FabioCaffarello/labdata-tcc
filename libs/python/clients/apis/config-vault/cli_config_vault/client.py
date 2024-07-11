from typing import Any, Dict, List

from dto_config_vault.output import ConfigDTO
from pyrequest.async_factory import RateLimitedAsyncHttpClient
from pysd.sd import new_from_env
from pyserializer.serializer import serialize_to_dataclass
from pylog.log import setup_logging


logger = setup_logging(__name__, log_level="DEBUG")


class AsyncPyConfigVaultClient:
    """
    A client for handling asynchronous interactions with a configuration service.

    Args:
        base_url (str): The base URL of the configuration service.

    Attributes:
        max_calls (int): The maximum number of API calls allowed in a period.
        period (int): The time period (in seconds) in which API calls are rate-limited.
        client (RateLimitedAsyncHttpClient): An instance of the rate-limited HTTP client.
        configs_endpoint (str): The endpoint for configuration operations.
    """
    def __init__(self, base_url: str):
        self.max_calls = 100
        self.period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.max_calls, self.period)
        self.configs_endpoint = "/config"

    async def create(self, data: Dict[str, Any]) -> ConfigDTO:
        """
        Create a new configuration using the provided data.

        Args:
            data (Dict[str, Any]): The configuration data to be created.

        Returns:
            ConfigDTO: The created configuration in the form of a data class.
        """
        config_data = await self.client.make_request("POST", self.configs_endpoint, data=data)
        return serialize_to_dataclass(config_data, ConfigDTO)

    async def update_config(self, data: Dict[str, Any]) -> ConfigDTO:
        """
        Update an existing configuration using the provided data.

        Args:
            data (Dict[str, Any]): The configuration data to be updated.

        Returns:
            ConfigDTO: The updated configuration in the form of a data class.
        """
        config_data = await self.client.make_request("PUT", self.configs_endpoint, data=data)
        return serialize_to_dataclass(config_data, ConfigDTO)

    async def list_all_configs(self) -> List[ConfigDTO]:
        """
        Retrieve a list of all configurations.

        Returns:
            List[ConfigDTO]: A list of all configurations in the form of data classes.
        """
        configs_data = await self.client.make_request("GET", self.configs_endpoint)
        return [serialize_to_dataclass(config, ConfigDTO) for config in configs_data]

    async def get_config_by_id(self, config_id: str) -> ConfigDTO:
        """
        Retrieve a specific configuration by its ID.

        Args:
            config_id (str): The unique identifier of the configuration.

        Returns:
            ConfigDTO: The requested configuration in the form of a data class.
        """
        endpoint = f"{self.configs_endpoint}/{config_id}"
        config_data = await self.client.make_request("GET", endpoint)
        return serialize_to_dataclass(config_data, ConfigDTO)

    async def delete_config(self, config_id: str) -> None:
        """
        Delete a specific configuration by its ID.

        Args:
            config_id (str): The unique identifier of the configuration.

        Returns:
            None
        """
        endpoint = f"{self.configs_endpoint}/{config_id}"
        await self.client.make_request("DELETE", endpoint)

    async def list_configs_by_service_and_provider(self, provider: str, service: str) -> List[ConfigDTO]:
        """
        Retrieve a list of configurations associated with a specific service and provider.

        Args:
            provider (str): The name of the provider.
            service (str): The name of the service.

        Returns:
            List[ConfigDTO]: A list of configurations for the specified service and provider in the form
            of data classes.
        """
        endpoint = f"{self.configs_endpoint}/provider/{provider}/service/{service}"
        configs_data = await self.client.make_request("GET", endpoint)
        return [serialize_to_dataclass(config, ConfigDTO) for config in configs_data]

    async def list_configs_by_source_and_provider(self, provider: str, source: str) -> List[ConfigDTO]:
        """
        Retrieve a list of configurations associated with a specific source and provider.

        Args:
            provider (str): The name of the provider.
            source (str): The name of the source.

        Returns:
            List[ConfigDTO]: A list of configurations for the specified source and provider in the form
            of data classes.
        """
        endpoint = f"{self.configs_endpoint}/provider/{provider}/source/{source}"
        configs_data = await self.client.make_request("GET", endpoint)
        return [serialize_to_dataclass(config, ConfigDTO) for config in configs_data]

    async def list_configs_by_service_provider_and_active(
        self,
        provider: str,
        service: str,
        active: bool
    ) -> List[ConfigDTO]:
        """
        Retrieve a list of configurations associated with a specific service, provider, and active status.

        Args:
            provider (str): The name of the provider.
            service (str): The name of the service.
            active (bool): The active status of the configuration.

        Returns:
            List[ConfigDTO]: A list of configurations for the specified service, provider, and active status in the
            form of data classes.
        """
        endpoint = f"{self.configs_endpoint}/provider/{provider}/service/{service}/active/{active}"
        configs_data = await self.client.make_request("GET", endpoint)
        return [serialize_to_dataclass(config, ConfigDTO) for config in configs_data]

    async def list_configs_by_service_source_and_provider(
        self,
        provider: str,
        service: str,
        source: str
    ) -> List[ConfigDTO]:
        """
        Retrieve a list of configurations associated with a specific service, source, and provider.

        Args:
            provider (str): The name of the provider.
            service (str): The name of the service.
            source (str): The name of the source.

        Returns:
            List[ConfigDTO]: A list of configurations for the specified service, source, and provider in the form of
            data classes.
        """
        endpoint = f"{self.configs_endpoint}/provider/{provider}/service/{service}/source/{source}"
        configs_data = await self.client.make_request("GET", endpoint)
        return [serialize_to_dataclass(config, ConfigDTO) for config in configs_data]

    async def list_configs_by_provider_and_dependencies(
        self,
        provider: str,
        service: str,
        source: str
    ) -> List[ConfigDTO]:
        """
        Retrieve a list of configurations associated with a specific provider, service, and dependencies.

        Args:
            provider (str): The name of the provider.
            service (str): The name of the service.
            source (str): The name of the source.

        Returns:
            List[ConfigDTO]: A list of configurations for the specified provider, service, and dependencies in the
            form of data classes.
        """
        endpoint = f"{self.configs_endpoint}/provider/{provider}/dependencies/service/{service}/source/{source}"
        configs_data = await self.client.make_request("GET", endpoint)
        return [serialize_to_dataclass(config, ConfigDTO) for config in configs_data]


def async_py_config_vault_client() -> AsyncPyConfigVaultClient:
    """
    Create an instance of the AsyncPyConfigHandlerClient using service discovery information.

    Returns:
        AsyncPyConfigHandlerClient: An instance of the configuration handler client.
    """
    sd = new_from_env()
    return AsyncPyConfigVaultClient(sd.services_config_vault_endpoint)
