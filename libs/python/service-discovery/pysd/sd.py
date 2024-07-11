import os
from dataclasses import dataclass
from typing import Dict, Optional
from pylog.log import setup_logging

logger = setup_logging(__name__)


class UnrecoverableError(Exception):
    """Exception raised for errors that are unrecoverable."""
    def __init__(self, message: str) -> None:
        """
        Initializes an UnrecoverableError with a given message.

        Args:
            message (str): The error message.
        """
        super().__init__(message)


class ServiceUnavailableError(Exception):
    """Exception raised for service availability errors."""
    def __init__(self, message: str) -> None:
        """
        Initializes a ServiceUnavailableError with a given message.

        Args:
            message (str): The error message.
        """
        super().__init__(message)


@dataclass(frozen=True)
class ServiceVars:
    """Class to hold service variable constants."""
    RABBITMQ: str = "RABBITMQ"
    MINIO: str = "MINIO"
    MONGODB: str = "MONGODB"
    SERVICES_RABBITMQ_EXCHANGE: str = "services"
    CONFIG_VAULT: str = "CONFIG_VAULT"
    SCHEMA_VAULT: str = "SCHEMA_VAULT"


class ServiceDiscovery:
    """Class to handle service discovery using environment variables."""
    GATEWAY_ENVIRONMENT: str = 'GATEWAY_ENVIRONMENT'

    def __init__(self, envvars: Optional[Dict[str, str]]) -> None:
        """
        Initializes a ServiceDiscovery instance.

        Args:
            envvars (Optional[Dict[str, str]]): A dictionary of environment variables.

        Raises:
            UnrecoverableError: If environment variables are not set.
        """
        if envvars is None:
            raise UnrecoverableError('Environment variables not set')
        self._vars = envvars
        self._service_vars = ServiceVars()

    def _get_endpoint(self, var_name: str, service_name: str, protocol: str = "http") -> str:
        """
        Gets the endpoint for a service.

        Args:
            var_name (str): The name of the environment variable containing the service endpoint.
            service_name (str): The name of the service.
            protocol (str): The protocol to use (default is "http").

        Returns:
            str: The service endpoint.

        Raises:
            ServiceUnavailableError: If the environment variable is not set.
        """
        if var_name not in self._vars:
            raise ServiceUnavailableError(f'Environment variable {var_name} not set')
        tcp_addr = self._vars[var_name]
        gt_host = self._get_gateway_host(service_name)
        return tcp_addr.replace("tcp", protocol).replace("gateway_host", gt_host)

    def _get_gateway_host(self, service_name: str) -> str:
        """
        Gets the gateway host for a service.

        Args:
            service_name (str): The name of the service.

        Returns:
            str: The gateway host.
        """
        return os.getenv(f'{service_name}_GATEWAY_HOST', 'localhost')

    def _modify_localhost_port(self, endpoint: str, original_port: str, new_port: str) -> str:
        """
        Modifies the port for localhost endpoints.

        Args:
            endpoint (str): The original endpoint.
            original_port (str): The original port to replace.
            new_port (str): The new port to set.

        Returns:
            str: The modified endpoint.
        """
        if "localhost" in endpoint:
            return endpoint.replace(original_port, new_port)
        return endpoint

    @property
    def services_config_vault_endpoint(self) -> str:
        logger.info(f"services_config_vault_endpoint: {self._service_vars.CONFIG_VAULT}")
        endpoint = self._get_endpoint("CONFIG_VAULT_PORT_8000_TCP", self._service_vars.CONFIG_VAULT)
        logger.info(f"services_config_vault_endpoint: {endpoint}")
        mod_endpoint = self._modify_localhost_port(endpoint, "8001", "8000")
        logger.info(f"services_config_vault_endpoint: {mod_endpoint}")
        return mod_endpoint

    @property
    def services_schema_vault_endpoint(self) -> str:
        endpoint = self._get_endpoint("SCHEMA_VAULT_PORT_8000_TCP", self._service_vars.SCHEMA_VAULT)
        return self._modify_localhost_port(endpoint, "8002", "8000")

    @property
    def rabbitmq_endpoint(self) -> str:
        """
        Gets the RabbitMQ endpoint.

        Returns:
            str: The RabbitMQ endpoint in 'amqp' protocol.
        """
        return self._get_endpoint("RABBITMQ_PORT_6572_TCP", self._service_vars.RABBITMQ, protocol="amqp")

    @property
    def services_rabbitmq_exchange(self) -> str:
        """
        Gets the services RabbitMQ exchange.

        Returns:
            str: The name of the services RabbitMQ exchange.
        """
        return self._service_vars.SERVICES_RABBITMQ_EXCHANGE

    @property
    def minio_endpoint(self) -> str:
        """
        Gets the Minio endpoint.

        Returns:
            str: The Minio endpoint.
        """
        return self._get_endpoint("MINIO_PORT_9000_TCP", self._service_vars.MINIO)

    @property
    def minio_access_key(self) -> Optional[str]:
        """
        Gets the Minio access key.

        Returns:
            Optional[str]: The Minio access key.
        """
        return os.getenv("MINIO_ACCESS_KEY")

    @property
    def minio_secret_key(self) -> Optional[str]:
        """
        Gets the Minio secret key.

        Returns:
            Optional[str]: The Minio secret key.
        """
        return os.getenv("MINIO_SECRET_KEY")

    @property
    def mongodb_endpoint(self) -> str:
        """
        Gets the MongoDB endpoint.

        Returns:
            str: The MongoDB endpoint.
        """
        return self._get_endpoint("MONGODB_PORT_27017_TCP", self._service_vars.MONGODB)


def new_from_env() -> ServiceDiscovery:
    """
    Creates a ServiceDiscovery instance using environment variables.

    Returns:
        ServiceDiscovery: A new ServiceDiscovery instance.
    """
    return ServiceDiscovery(dict(os.environ))
