from typing import Dict
import httpx
import asyncio
import time
from pylog.log import setup_logging

logger = setup_logging(__name__, log_level="DEBUG")


class RateLimitedAsyncHttpClient:
    """
    A rate-limited asynchronous HTTP client.

    Attributes:
        base_url (str): The base URL for the HTTP client.
        max_calls (int): The maximum number of calls allowed in the specified period.
        period (int): The period in seconds for rate limiting.
        semaphore (asyncio.Semaphore): A semaphore to limit concurrent requests.
        lock (asyncio.Lock): A lock to synchronize access to the rate limiting mechanism.
        last_reset (float): The time when the rate limit was last reset.
        call_count (int): The current count of calls in the current period.
        timeout (int): The timeout for HTTP requests.
        retries (int): The number of retries for failed requests.
        logger: A logger instance for logging debug information.
    """

    def __init__(self, base_url: str, max_calls: int, period: int) -> None:
        """
        Initialize the RateLimitedAsyncHttpClient.

        Args:
            base_url (str): The base URL for the HTTP client.
            max_calls (int): The maximum number of calls allowed in the specified period.
            period (int): The period in seconds for rate limiting.
        """
        self.base_url = base_url
        self.max_calls = max_calls
        self.period = period
        self.semaphore = asyncio.Semaphore(max_calls)
        self.lock = asyncio.Lock()
        self.last_reset = time.time()
        self.call_count = 0
        self.timeout = 10  # Increase timeout
        self.retries = 5  # Increase retries
        self.logger = logger

    async def _acquire(self) -> None:
        """
        Acquire a permit for making an HTTP request, respecting the rate limit.

        This method blocks if the rate limit has been reached and waits until a new period starts.
        """
        async with self.lock:
            current_time = time.time()
            elapsed = current_time - self.last_reset
            if elapsed >= self.period:
                self.last_reset = current_time
                self.call_count = 0
            self.call_count += 1
            if self.call_count > self.max_calls:
                await asyncio.sleep(self.period - elapsed)
                self.last_reset = time.time()
                self.call_count = 1

    async def make_request(
        self, method: str, endpoint: str, data: Dict[str, any] = None, params: Dict[str, any] = None
    ) -> Dict[str, any]:
        """
        Make an HTTP request to the specified endpoint with rate limiting.

        Args:
            method (str): The HTTP method to use for the request (e.g., 'GET', 'POST').
            endpoint (str): The endpoint to send the request to.
            data (Dict[str, any], optional): The JSON data to send in the request body.
            params (Dict[str, any], optional): The query parameters to include in the request.

        Returns:
            Dict[str, any]: The JSON response from the server.

        Raises:
            httpx.RequestError: If the request fails after the specified number of retries.
            httpx.HTTPStatusError: If the server returns an error response.
        """
        url = self.base_url + endpoint
        await self._acquire()
        for attempt in range(self.retries):
            try:
                self.logger.info(f"Making request to {url}")
                async with httpx.AsyncClient(timeout=self.timeout) as client:
                    response = await client.request(method, url, json=data, params=params)
                    response.raise_for_status()
                    if response.status_code == httpx.codes.NO_CONTENT:
                        return {}
                    return response.json()
            except (httpx.HTTPStatusError, httpx.RequestError) as e:
                self.logger.error(f"Attempt {attempt + 1}/{self.retries} - Request failed: {e}")
                if attempt + 1 == self.retries:
                    raise
                await asyncio.sleep(2 ** attempt)  # Exponential backoff
