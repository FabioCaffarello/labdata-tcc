import unittest
from unittest.mock import MagicMock, patch

from pyyoutube.client import YoutubeDownloaderError, download_to_buffer


class TestYoutubeDownloader(unittest.TestCase):

    @patch('pyyoutube.client.YoutubeDL')
    def test_download_successful(self, mock_youtubedl):
        # Mock info dict
        mock_info_dict = {'url': 'http://testvideo.com/video'}

        # Mock ydl instance
        mock_ydl_instance = MagicMock()
        mock_ydl_instance.extract_info.return_value = mock_info_dict
        mock_ydl_instance.urlopen.return_value.read.return_value = b'test video data'
        mock_youtubedl.return_value.__enter__.return_value = mock_ydl_instance

        result = download_to_buffer('http://testvideo.com')

        self.assertEqual(result, b'test video data')
        mock_ydl_instance.extract_info.assert_called_once_with('http://testvideo.com', download=False)
        mock_ydl_instance.urlopen.return_value.read.assert_called_once()

    @patch('pyyoutube.client.YoutubeDL')
    def test_no_video_url(self, mock_youtubedl):
        # Mock info dict with no URL
        mock_info_dict = {}

        # Mock ydl instance
        mock_ydl_instance = MagicMock()
        mock_ydl_instance.extract_info.return_value = mock_info_dict
        mock_youtubedl.return_value.__enter__.return_value = mock_ydl_instance

        with self.assertRaises(YoutubeDownloaderError) as context:
            download_to_buffer('http://testvideo.com')

        self.assertEqual(str(context.exception), 'Failed to download video: No video URL found.')
        mock_ydl_instance.extract_info.assert_called_once_with('http://testvideo.com', download=False)

    @patch('pyyoutube.client.YoutubeDL')
    def test_download_failure(self, mock_youtubedl):
        # Mock ydl instance to raise an exception
        mock_ydl_instance = MagicMock()
        mock_ydl_instance.extract_info.side_effect = Exception('test error')
        mock_youtubedl.return_value.__enter__.return_value = mock_ydl_instance

        with self.assertRaises(YoutubeDownloaderError) as context:
            download_to_buffer('http://testvideo.com')

        self.assertTrue('Failed to download video: test error' in str(context.exception))
        mock_ydl_instance.extract_info.assert_called_once_with('http://testvideo.com', download=False)


if __name__ == '__main__':
    unittest.main()
