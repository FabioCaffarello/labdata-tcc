import os
import base_fixture.fixture as base_fixtures
from service_fixture.service_mapping import ServiceAppNameFixtureMapping

service_name = os.getenv("SERVICE_NAME")


class ServiceTestsFixture(base_fixtures.BaseTestsFixture):
    """
    A fixture class for setting up and tearing down service tests.

    Attributes:
        service_name (str): The name of the service.
        cmd (str): The command to run the service.
    """
    service_name = service_name
    cmd = "python"

    async def asyncSetUp(self):
        """
        Asynchronously sets up the test environment.
        
        Returns:
            None
        """
        return await super().asyncSetUp()

    async def asyncTearDown(self):
        """
        Asynchronously tears down the test environment.
        
        Returns:
            None
        """
        return await super().asyncTearDown()

    def get_app_name(self) -> str:
        """
        Gets the application name based on the service name.

        Returns:
            str: The application name.
        """
        return ServiceAppNameFixtureMapping.get_app_name(self.service_name)

    def _get_service_process_args(self) -> list:
        """
        Generates the service process arguments dynamically.

        Returns:
            list: The service process arguments.
        """
        app_name = self.get_app_name()
        args = [
            self.cmd,
            f"/app/{app_name}/main.py",
        ]
        return args
