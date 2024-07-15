import unittest
from unittest.mock import patch, MagicMock, AsyncMock
from subprocessd.async_subprocessd import SubprocessDAsync


class TestSubprocessDAsync(unittest.IsolatedAsyncioTestCase):

    @patch("subprocessd.async_subprocessd.asyncio.create_subprocess_exec", new_callable=AsyncMock)
    @patch("subprocessd.async_subprocessd.os.makedirs")
    @patch("subprocessd.async_subprocessd.os.path.isdir")
    @patch("subprocessd.async_subprocessd.open", new_callable=unittest.mock.mock_open)
    async def test_start_creates_log_dir(self, mock_open, mock_isdir, mock_makedirs, mock_create_subprocess_exec):
        mock_isdir.return_value = False
        mock_subprocess = AsyncMock()
        mock_create_subprocess_exec.return_value = mock_subprocess

        subprocess_args = ['ls', '-l']
        log_file_path = '/fake/path'
        subprocess = SubprocessDAsync(subprocess_args, log_file_path)
        await subprocess.start()

        mock_isdir.assert_called_once_with(log_file_path)
        mock_makedirs.assert_called_once_with(log_file_path)
        mock_create_subprocess_exec.assert_called_once_with(
            *subprocess_args,
            stdout=unittest.mock.ANY,
            stderr=unittest.mock.ANY
        )
        mock_open.assert_called_once_with('/fake/path/logs.txt', 'w')

    @patch("subprocessd.async_subprocessd.asyncio.create_subprocess_exec", new_callable=AsyncMock)
    @patch("subprocessd.async_subprocessd.os.makedirs")
    @patch("subprocessd.async_subprocessd.os.path.isdir")
    @patch("subprocessd.async_subprocessd.open", new_callable=unittest.mock.mock_open)
    async def test_start_with_existing_log_dir(self, mock_open, mock_isdir, mock_makedirs, mock_create_subprocess_exec):
        mock_isdir.return_value = True
        mock_subprocess = AsyncMock()
        mock_create_subprocess_exec.return_value = mock_subprocess

        subprocess_args = ['ls', '-l']
        log_file_path = '/fake/path'
        subprocess = SubprocessDAsync(subprocess_args, log_file_path)
        await subprocess.start()

        mock_isdir.assert_called_once_with(log_file_path)
        mock_makedirs.assert_not_called()
        mock_create_subprocess_exec.assert_called_once_with(
            *subprocess_args,
            stdout=unittest.mock.ANY,
            stderr=unittest.mock.ANY
        )
        mock_open.assert_called_once_with('/fake/path/logs.txt', 'w')

    @patch("subprocessd.async_subprocessd.asyncio.create_subprocess_exec", new_callable=AsyncMock)
    @patch("subprocessd.async_subprocessd.os.makedirs")
    @patch("subprocessd.async_subprocessd.os.path.isdir")
    @patch("subprocessd.async_subprocessd.open", new_callable=unittest.mock.mock_open)
    async def test_stop(self, mock_open, mock_isdir, mock_makedirs, mock_create_subprocess_exec):
        mock_isdir.return_value = True
        mock_subprocess = AsyncMock()
        mock_subprocess.kill = MagicMock()
        mock_subprocess.wait = AsyncMock()
        mock_create_subprocess_exec.return_value = mock_subprocess

        subprocess_args = ['ls', '-l']
        log_file_path = '/fake/path'
        subprocess = SubprocessDAsync(subprocess_args, log_file_path)
        await subprocess.start()
        await subprocess.stop()

        mock_create_subprocess_exec.assert_called_once_with(
            *subprocess_args,
            stdout=unittest.mock.ANY,
            stderr=unittest.mock.ANY
        )
        mock_subprocess.kill.assert_called_once()
        mock_subprocess.wait.assert_awaited_once()
        mock_open.assert_called_once_with('/fake/path/logs.txt', 'w')
        self.assertTrue(mock_open().close.called)


if __name__ == '__main__':
    unittest.main()
