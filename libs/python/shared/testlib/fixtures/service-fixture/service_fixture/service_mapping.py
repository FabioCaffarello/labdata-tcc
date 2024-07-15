from enum import Enum


def get_service_name(service_name: str) -> str:
    """
    Converts the service name to a format suitable for the enum.

    Args:
        service_name (str): The original service name.

    Returns:
        str: The formatted service name.
    """
    return service_name.replace("-", "_")


class ServiceAppNameFixtureMapping(Enum):
    """
    Enum for mapping service names to application names.
    """
    video_downloader = "downloader"

    @classmethod
    def get_app_name(cls, service_name: str) -> str:
        """
        Gets the application name based on the service name.

        Args:
            service_name (str): The service name.

        Returns:
            str: The application name.

        Raises:
            ValueError: If no mapping is found for the service name.
        """
        app_name_key = get_service_name(service_name)

        if app_name_key in cls.__members__:
            return cls.__members__[app_name_key].value
        else:
            raise ValueError(f"No mapping found for service name: {service_name}")
