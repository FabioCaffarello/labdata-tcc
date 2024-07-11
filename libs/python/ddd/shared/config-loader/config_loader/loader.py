from typing import Dict
from pylog.log import setup_logging
from cli_config_vault.client import async_py_config_vault_client
from dto_config_vault.output import ConfigDTO

logger = setup_logging(__name__)


class ConfigLoader:
    """
    ConfigLoader class to load configurations from the config handler API client.
    """
    def __init__(self) -> None:
        """
        Initialize the ConfigLoader.

        Attributes:
            mapping_config (Dict[str, Dict[str, ConfigDTO]]): Dictionary to store configurations.
        """
        self.mapping_config: Dict[str, ConfigDTO] = {}
        self.__config_vault_api_client = async_py_config_vault_client()

    async def fetch_configs_for_service(self, service_name: str, provider: str) -> Dict[str, ConfigDTO]:
        """
        Fetch configurations for a specific service and context environment.

        Args:
            service_name (str): The name of the service for which configurations are to be fetched.
            provider (str): The provider for which configurations are to be fetched.

        Returns:
            Dict[str, Dict[str, ConfigDTO]]: A dictionary containing configurations organized by context and ID.
        """
        logger.info(f"Fetching configurations for service: {service_name}, provider: {provider}")
        configs = await self.__config_vault_api_client.list_configs_by_service_and_provider(
            provider=provider,
            service=service_name
        )
        for config in configs:
            self.register_config(
                config.config_id,
                config
            )
            logger.info(f"Registered config: {config.config_id} for provider: {config.provider}")
        return self.mapping_config

    def register_config(self, config_id: str, config: ConfigDTO) -> None:
        """
        Register a configuration in the mapping_config dictionary.

        Args:
            context (str): The context for which the configuration is being registered.
            config_id (str): The ID of the configuration.
            config (ConfigDTO): The configuration object to be registered.

        Raises:
            ValueError: If a configuration with the same ID already exists in the context.
        """
        if config_id in self.mapping_config:
            raise ValueError(f"Duplicate config ID '{config_id}'. Overwriting existing config.")

        self.mapping_config[config_id] = config


async def fetch_configs(service: str, provider: str) -> Dict[str, ConfigDTO]:
    """
    Fetch configurations for a given service and context environment.

    Args:
        service (str): The name of the service for which configurations are to be fetched.
        context_env (str): The context environment for which configurations are to be fetched.

    Returns:
        Dict[str, Dict[str, ConfigDTO]]: A dictionary containing configurations organized by context and ID.
    """
    config_loader = ConfigLoader()
    return await config_loader.fetch_configs_for_service(service, provider)
