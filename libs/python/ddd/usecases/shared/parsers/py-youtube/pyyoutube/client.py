import io

from pylog.log import setup_logging
from yt_dlp import YoutubeDL

logger = setup_logging(__name__, log_level="DEBUG")


class YoutubeDownloaderError(Exception):
    """
    Custom exception class for YoutubeDownloader errors.

    Attributes:
        message (str): The error message.
    """

    def __init__(self, message: str):
        """
        Initialize the YoutubeDownloaderError.

        Args:
            message (str): The error message.
        """
        self.message = message
        super().__init__(self.message)


def download_to_buffer(url: str) -> bytes:
    """
    Download a video from the given URL to a buffer and return its bytes.

    Args:
        url (str): The URL of the video to download.

    Returns:
        bytes: The downloaded video as bytes.

    Raises:
        YoutubeDownloaderError: If the video URL is not found or if the download fails.
    """
    logger.info(f"Input URL: {url}")
    buffer = io.BytesIO()

    ydl_opts = {
        'format': 'best',
        'outtmpl': '-',
        'noplaylist': True,
        'quiet': True
    }

    try:
        with YoutubeDL(ydl_opts) as ydl:
            info_dict = ydl.extract_info(url, download=False)
            video_url = info_dict.get('url', None)

            if not video_url:
                raise YoutubeDownloaderError("No video URL found.")

            video_stream = ydl.urlopen(video_url)
            buffer.write(video_stream.read())

        logger.info("Video downloaded successfully.")
        bytes_data = buffer.getvalue()
    except Exception as e:
        logger.error(f"Unexpected error: {e}")
        raise YoutubeDownloaderError(f"Failed to download video: {e}")
    finally:
        buffer.close()

    return bytes_data
