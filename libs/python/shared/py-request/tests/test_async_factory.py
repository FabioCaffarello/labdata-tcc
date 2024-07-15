import unittest
import asyncio
import httpx
from pyrequest.async_factory import RateLimitedAsyncHttpClient
import respx
from httpx import Response


class TestRateLimitedAsyncHttpClient(unittest.IsolatedAsyncioTestCase):

    def setUp(self):
        self.base_url = "https://example.com"
        self.client = RateLimitedAsyncHttpClient(self.base_url, max_calls=2, period=1)

    @respx.mock
    async def test_make_get_request(self):
        endpoint = "/test"
        url = self.base_url + endpoint
        respx.get(url).mock(return_value=Response(200, json={"message": "success"}))

        response = await self.client.make_request("GET", endpoint)
        self.assertEqual(response, {"message": "success"})

    @respx.mock
    async def test_make_post_request(self):
        endpoint = "/post-test"
        request_data = {"key": "value"}
        url = self.base_url + endpoint
        respx.post(url).mock(return_value=Response(201, json={"message": "created"}))

        response = await self.client.make_request("POST", endpoint, data=request_data)
        self.assertEqual(response, {"message": "created"})

    @respx.mock
    async def test_rate_limiting(self):
        endpoint = "/rate-limit-test"
        url = self.base_url + endpoint
        respx.get(url).mock(return_value=Response(200, json={"message": "success"}))

        tasks = [self.client.make_request("GET", endpoint) for _ in range(5)]
        responses = await asyncio.gather(*tasks, return_exceptions=True)

        success_responses = [response for response in responses if isinstance(response, dict)]
        self.assertEqual(len(success_responses), 5)

    @respx.mock
    async def test_http_error_handling(self):
        endpoint = "/error"
        url = self.base_url + endpoint
        respx.get(url).mock(return_value=Response(404))

        with self.assertRaises(httpx.HTTPStatusError):
            await self.client.make_request("GET", endpoint)


if __name__ == "__main__":
    unittest.main()
