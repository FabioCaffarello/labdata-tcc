from unittest.mock import AsyncMock, MagicMock, patch

import pytest
import warlock
from dto_config_vault.output import ConfigDTO, JobParametersDTO
from dto_events_router.input import MetadataDTO, ServiceFeedBackDTO
from job_handler.core import Importer, JobHandler
from pydebug import debug


class TestImporter:

    def setup_method(self):
        self.job_parser = "test_parser"

    @patch('importlib.import_module')
    @patch('job_handler.core.logger')
    def test_import_successful(self, mock_logger, mock_import_module):
        mock_module = MagicMock()
        mock_import_module.return_value = mock_module

        importer = Importer(self.job_parser)

        assert importer.module == mock_module
        mock_import_module.assert_called_once_with(f"{self.job_parser}.job")
        mock_logger.info.assert_any_call(f"Getting module reference: {self.job_parser}")
        mock_logger.info.assert_any_call(f"Importing job parser: {self.job_parser}.job")

    @patch('importlib.import_module')
    @patch('job_handler.core.logger')
    def test_import_failure(self, mock_logger, mock_import_module):
        mock_import_module.side_effect = ModuleNotFoundError("Module not found")

        with pytest.raises(ImportError) as context:
            Importer(self.job_parser)

        assert str(context.value) == "Failed to import job parser: Module not found"
        mock_import_module.assert_called_once_with(f"{self.job_parser}.job")
        mock_logger.info.assert_any_call(f"Getting module reference: {self.job_parser}")
        mock_logger.info.assert_any_call(f"Importing job parser: {self.job_parser}.job")


class TestJobHandler:

    def setup_method(self):
        self.config = MagicMock(spec=ConfigDTO)
        self.config.config_id = "test_config_id"  # Ensure config_id is set
        self.job_parameters = MagicMock(spec=JobParametersDTO)
        self.config.job_parameters = self.job_parameters
        self.config.job_parameters.parser_module = "test_parser"
        self.metadata = MagicMock(spec=MetadataDTO)
        self.dbg = MagicMock(spec=debug.EnabledDebug)

    @patch('importlib.import_module')
    @patch('job_handler.core.logger')
    def test_job_handler_init(self, mock_logger, mock_import_module):
        mock_module = MagicMock()
        mock_import_module.return_value = mock_module

        job_handler = JobHandler(self.config, self.metadata, self.dbg)

        assert job_handler.config == self.config
        assert job_handler.metadata == self.metadata
        assert job_handler.dbg == self.dbg
        assert job_handler.job_parser == "test_parser"
        mock_import_module.assert_called_once_with("test_parser.job")
        mock_logger.info.assert_any_call(f"Getting Config {self.config}")
        mock_logger.info.assert_any_call(f"Getting job parser: {self.config.job_parameters.parser_module}")

    @pytest.mark.asyncio
    @patch('importlib.import_module')
    @patch('job_handler.core.logger')
    async def test_execute(self, mock_logger, mock_import_module):
        mock_module = MagicMock()
        mock_job_instance = AsyncMock()
        mock_job_instance.run = AsyncMock(return_value=MagicMock(spec=ServiceFeedBackDTO))
        mock_module.Job.return_value = mock_job_instance
        mock_import_module.return_value = mock_module

        job_handler = JobHandler(self.config, self.metadata, self.dbg)
        source_input = MagicMock(spec=warlock.model.Model)

        job_data = await job_handler.execute(source_input)

        assert job_data == mock_job_instance.run.return_value
        mock_module.Job.assert_called_once_with(self.config, self.metadata, self.dbg)
        mock_job_instance.run.assert_awaited_once_with(source_input)
        mock_logger.info.assert_any_call(
            f"[RUNNING JOB] - Config ID: {job_handler.config_id} - handler: {job_handler.job_parser}"
        )


if __name__ == '__main__':
    pytest.main()
