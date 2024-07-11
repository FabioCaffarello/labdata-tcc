import importlib
from typing import Union, Any
from dto_config_vault.output import ConfigDTO
from pylog.log import setup_logging
from pydebug import debug
import warlock
from dto_events_router.input import ServiceFeedBackDTO, MetadataDTO


logger = setup_logging(__name__)


class Importer:
    """Handles the dynamic import of job parser modules.

    Args:
        job_parser (str): The name of the job parser module to import.

    Attributes:
        job_parser (str): The name of the job parser module.
        __module (Any): The imported job parser module.
    """
    def __init__(self, job_parser: str):
        self.job_parser = job_parser
        self.__module = self.__import_module()

    def __get_module_ref(self) -> str:
        """Generates the module reference string.

        Returns:
            str: The module reference string.
        """
        logger.info(f"Getting module reference: {self.job_parser}")
        return f"{self.job_parser}.job"

    def __import_module(self) -> Any:
        """Imports the job parser module.

        Returns:
            Any: The imported job parser module.

        Raises:
            ImportError: If the module cannot be imported.
        """
        module_ref = self.__get_module_ref()
        logger.info(f"Importing job parser: {module_ref}")
        try:
            return importlib.import_module(module_ref)
        except ModuleNotFoundError as e:
            raise ImportError(f"Failed to import job parser: {e}")

    @property
    def module(self) -> Any:
        """Returns the imported job parser module.

        Returns:
            Any: The imported job parser module.
        """
        return self.__module


class JobHandler(Importer):
    """Handles the execution of jobs using the imported job parser.

    Args:
        config (ConfigDTO): Configuration data transfer object.
        metadata (MetadataDTO): Metadata data transfer object.
        dbg (Union[debug.EnabledDebug, debug.DisabledDebug]): Debugging object.
    """
    def __init__(self, config: ConfigDTO, metadata: MetadataDTO, dbg: Union[debug.EnabledDebug, debug.DisabledDebug]):
        self.config = config
        self.dbg = dbg
        self.metadata = metadata
        self.job_parser = self.__get_job_parser()
        super().__init__(self.job_parser)

    @property
    def config_id(self) -> str:
        """str: Returns the configuration ID."""
        return self.config.config_id

    def __get_job_parser(self) -> str:
        """Extracts the job parser module name from the configuration.

        Returns:
            str: The job parser module name.
        """
        logger.info(f"Getting Config {self.config}")
        logger.info(f"Getting job parser: {self.config.job_parameters.parser_module}")
        return self.config.job_parameters.parser_module

    async def execute(self, source_input: type[warlock.model.Model]) -> ServiceFeedBackDTO:
        """Executes the job using the provided input data.

        Args:
            source_input (type[warlock.model.Model]): The input data for the job.

        Returns:
            ServiceFeedBackDTO: The feedback data transfer object containing the result of the job execution.
        """
        logger.info(f"[RUNNING JOB] - Config ID: {self.config_id} - handler: {self.job_parser}")

        job_data = await self.module.Job(self.config, self.metadata, self.dbg).run(source_input)
        return job_data
