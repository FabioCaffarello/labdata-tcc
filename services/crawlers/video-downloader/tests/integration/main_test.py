import asyncio # noqa
import unittest

import pytest
from pylog.log import setup_logging
from service_fixture.fixture import ServiceTestsFixture

logger = setup_logging(__name__)


class DefaultHandlerIntegrationTests(ServiceTestsFixture):

    async def asyncSetUp(self):
        return await super().asyncSetUp()

    async def asyncTearDown(self):
        return await super().asyncTearDown()

    @pytest.mark.asyncio
    async def test_should_download_youtube_video(self):
        valid_input = {
            "videoId": "XqZsoesa55w",
        }
        source = "pinkfong"
        provider = "kids"
        await self.push_job(valid_input, provider, source)
        logger.info("Job pushed to queue")
        await self.pop_job(provider, source, 60)
        self.assertFalse(self.queue.empty())
        while not self.queue.empty():
            result = await self.queue.get()
            logger.info(f"Job result: {result}")
            self.assertEqual(result["status"]["code"], 200)
            self.assertIsNotNone(result)


if __name__ == "__main__":
    unittest.main()
