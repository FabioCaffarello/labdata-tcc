# base-fixture

A Python library providing base fixtures for testing services. This library is designed to set up and tear down test environments asynchronously, push jobs to queues, and verify job results.

## Features

- Asynchronous setup and teardown of test environments.
- Pushing jobs to queues and retrieving results.
- Integration with unittest for testing.

## Installation

To install the `python-shared-testlib-fixtures-base-fixture` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-shared-testlib-fixtures-base-fixture --local
```

## Usage

### BaseTestsFixture

The `BaseTestsFixture` class provides base functionality for setting up and tearing down test environments, pushing jobs to queues, and verifying job results.

#### Example

```python
import asyncio
import unittest
from service_fixture.fixture import BaseTestsFixture
from pylog.log import setup_logging

logger = setup_logging(__name__)

class ExampleTests(BaseTestsFixture):

    async def asyncSetUp(self):
        return await super().asyncSetUp()

    async def asyncTearDown(self):
        return await super().asyncTearDown()

    async def test_example(self):
        valid_input = {
            "exampleKey": "exampleValue",
        }
        provider = "exampleProvider"
        source = "exampleSource"
        await self.push_job(valid_input, provider, source)
        logger.info("Job pushed to queue")
        await self.pop_job(provider, source, 2*60)
        self.assertFalse(self.queue.empty())
        while not self.queue.empty():
            result = await self.queue.get()
            logger.info(f"Job result: {result}")
            self.assertEqual(result["status"]["code"], 200)
            self.assertIsNotNone(result)

if __name__ == "__main__":
    unittest.main()
```
