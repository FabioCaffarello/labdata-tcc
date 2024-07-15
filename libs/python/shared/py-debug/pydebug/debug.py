import os
from pylog.log import setup_logging
import shutil

logger = setup_logging(__name__, log_level="DEBUG")


def new(debug_enabled: bool, debug_dir: str):
    """
    Create a new debug storage instance based on the debug flag.

    Args:
        debug_enabled (bool): Flag indicating if debug is enabled.
        debug_dir (str): Directory for storing debug responses.

    Returns:
        EnabledDebug or DisabledDebug: Debug storage instance.
    """
    if debug_enabled:
        logger.info(f"Creating debug enabled storage at: {debug_dir}")
        return EnabledDebug(debug_dir)
    else:
        logger.info("Debug disabled, creating stub debug storage")
        return DisabledDebug()


class EnabledDebug:
    """
    Enabled debug storage for saving responses.

    Attributes:
        _debug_dir (str): Directory for storing debug responses.
        _save_responses (dict): Dictionary to keep track of saved responses.
    """

    def __init__(self, debug_dir: str):
        """
        Initialize EnabledDebug with the specified directory.

        Args:
            debug_dir (str): Directory for storing debug responses.
        """
        self._debug_dir = debug_dir
        self._save_responses = {}
        EnabledDebug._create_dir(self._get_response_dir())

    def save_response(self, file_name: str, response_body: str):
        """
        Save the response body to a file in the debug directory.

        Args:
            file_name (str): Name of the response file.
            response_body (str): Content of the response to save.
        """
        filename = EnabledDebug._get_filename(file_name, self._save_responses)
        EnabledDebug._write_file(self._get_response_dir(), filename, response_body)

    def _get_response_dir(self):
        """
        Get the directory for storing responses.

        Returns:
            str: Directory path for storing responses.
        """
        return f"{self._debug_dir}/responses/"

    @staticmethod
    def _create_dir(path: str):
        """
        Create the specified directory, removing it first if it exists.

        Args:
            path (str): Directory path to create.
        """
        shutil.rmtree(path, ignore_errors=True)
        os.makedirs(path)

    @staticmethod
    def _get_filename(file_name: str, saved_files: dict):
        """
        Generate a unique filename for the response file.

        Args:
            file_name (str): Base name of the file.
            saved_files (dict): Dictionary to keep track of saved files.

        Returns:
            str: Unique filename.
        """
        if file_name in saved_files:
            saved_files[file_name] += 1
        else:
            saved_files[file_name] = 1
        count = saved_files[file_name]
        return f"{count}-{file_name}"

    @staticmethod
    def _write_file(dirname, file_name, file_to_write):
        """
        Write the response content to a file.

        Args:
            dirname (str): Directory name where the file will be saved.
            file_name (str): Name of the file.
            file_to_write (str): Content to write to the file.
        """
        logger.info(f"Writing file {file_name} to {dirname}")
        with open(dirname + file_name, "wb") as writer:
            writer.write(file_to_write)


class DisabledDebug:
    """
    Disabled debug storage that does not save responses.
    """

    def save_response(self, file_name: str, response_body: str):
        """
        Dummy method to match interface, does nothing.

        Args:
            file_name (str): Name of the response file.
            response_body (str): Content of the response to save.
        """
        pass
