import asyncio
import os
from pylog.log import setup_logging

logger = setup_logging(__name__)


class SubprocessDAsync:
    """
    A class to manage asynchronous subprocess execution with logging.

    Attributes:
        subprocess_args (list): Arguments for the subprocess.
        log_file_path (str): Path to the directory where logs will be stored.
    """

    def __init__(self, subprocess_args, log_file_path):
        """
        Initializes the SubprocessDAsync instance with subprocess arguments and log file path.

        Args:
            subprocess_args (list): Arguments for the subprocess.
            log_file_path (str): Path to the directory where logs will be stored.
        """
        self.subprocess_args = subprocess_args
        self.log_file_path = log_file_path
        self._subprocess = None

    async def start(self):
        """
        Starts the subprocess and logs the output to a file.
        """
        logger.debug("Starting subprocess with args: %s", self.subprocess_args)
        if not os.path.isdir(self.log_file_path):
            os.makedirs(self.log_file_path)
        self._logs = open(os.path.join(self.log_file_path, "logs.txt"), "w")
        self._subprocess = await asyncio.create_subprocess_exec(
            *self.subprocess_args,
            stdout=self._logs,
            stderr=self._logs
        )
        logger.debug("Subprocess started.")

    async def stop(self):
        """
        Stops the subprocess and closes the log file.
        """
        logger.debug("Stopping subprocess.")
        if self._subprocess is None:
            logger.debug("No subprocess to stop.")
            return
        self._subprocess.kill()
        logger.debug("Subprocess killed.")
        await self._subprocess.wait()
        logger.debug("Subprocess wait completed.")
        self._subprocess = None
        self._logs.close()
        logger.debug("Logs closed.")
