import os
import shutil
import unittest
from unittest.mock import patch, mock_open
from pydebug.debug import new, EnabledDebug, DisabledDebug
from pylog.log import setup_logging


class TestDebugStorage(unittest.TestCase):

    def setUp(self):
        self.debug_dir = "test_debug"
        self.response_body = b"response content"
        self.file_name = "response.txt"
        # Reset logger handlers for each test
        logger = setup_logging("pydebug.debug", log_level="DEBUG")
        logger.handlers = []

    def tearDown(self):
        shutil.rmtree(self.debug_dir, ignore_errors=True)

    @patch("pydebug.debug.logger")
    def test_new_debug_enabled(self, mock_logger):
        instance = new(True, self.debug_dir)
        mock_logger.info.assert_any_call(f"Creating debug enabled storage at: {self.debug_dir}")
        self.assertIsInstance(instance, EnabledDebug)

    @patch("pydebug.debug.logger")
    def test_new_debug_disabled(self, mock_logger):
        instance = new(False, self.debug_dir)
        mock_logger.info.assert_any_call("Debug disabled, creating stub debug storage")
        self.assertIsInstance(instance, DisabledDebug)

    @patch("builtins.open", new_callable=mock_open)
    @patch("pydebug.debug.os.makedirs")
    @patch("pydebug.debug.shutil.rmtree")
    @patch("pydebug.debug.logger")
    def test_enabled_debug_save_response(self, mock_logger, mock_rmtree, mock_makedirs, mock_open):
        enabled_debug = EnabledDebug(self.debug_dir)
        enabled_debug.save_response(self.file_name, self.response_body)
        response_dir = enabled_debug._get_response_dir()
        mock_rmtree.assert_called_with(response_dir, ignore_errors=True)
        mock_makedirs.assert_called_with(response_dir)
        mock_open.assert_called_with(os.path.join(response_dir, "1-response.txt"), "wb")
        mock_open().write.assert_called_with(self.response_body)
        mock_logger.info.assert_any_call(f"Writing file 1-response.txt to {response_dir}")

    @patch("pydebug.debug.logger")
    def test_disabled_debug_save_response(self, mock_logger):
        disabled_debug = DisabledDebug()
        disabled_debug.save_response(self.file_name, self.response_body)
        mock_logger.info.assert_not_called()

    def test_enabled_debug_get_filename(self):
        enabled_debug = EnabledDebug(self.debug_dir)
        saved_files = {}
        filename = enabled_debug._get_filename("file", saved_files)
        self.assertEqual(filename, "1-file")
        filename = enabled_debug._get_filename("file", saved_files)
        self.assertEqual(filename, "2-file")

    def test_enabled_debug_create_dir(self):
        path = "test_dir"
        with patch("pydebug.debug.os.makedirs") as mock_makedirs, patch("pydebug.debug.shutil.rmtree") as mock_rmtree:
            EnabledDebug._create_dir(path)
            mock_rmtree.assert_called_with(path, ignore_errors=True)
            mock_makedirs.assert_called_with(path)

    def test_enabled_debug_write_file(self):
        dirname = "test_dir/"
        filename = "test_file.txt"
        content = b"test content"
        with patch("builtins.open", mock_open()) as mocked_open:
            EnabledDebug._write_file(dirname, filename, content)
            mocked_open.assert_called_with(os.path.join(dirname, filename), "wb")
            mocked_open().write.assert_called_with(content)


if __name__ == '__main__':
    unittest.main()
