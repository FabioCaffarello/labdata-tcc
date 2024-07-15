# python-shared-testlib-fixtures-service-fixture

A Python library providing service fixtures for testing, specifically designed for services with different application names. This library facilitates setting up and tearing down test environments for services by leveraging the `BaseTestsFixture` and mapping service names to application names.

## Features

- Setup and teardown test environments asynchronously.
- Map service names to application names using an enum.
- Generate service process arguments dynamically based on service name.

## Installation

To install the `python-shared-testlib-fixtures-service-fixture` library, use the following command:

```bash
npx nx run <PROJECT>:add --name python-shared-testlib-fixtures-service-fixture --local
```

## Usage

### ServiceTestsFixture

The `ServiceTestsFixture` class provides methods to set up and tear down test environments, and generate service process arguments based on the service name.

#### Example

```python
import asyncio
import unittest
import pytest
from service_fixture.fixture import ServiceTestsFixture
from pylog.log import setup_logging


logger = setup_logging(__name__)


class DefaultHandlerIntegrationTests(ServiceTestsFixture):

    async def asyncSetUp(self):
        return await super().asyncSetUp()

    async def asyncTearDown(self):
        return await super().asyncTearDown()

    @pytest.mark.asyncio
    async def test_should(self):
        ...

if __name__ == "__main__":
    unittest.main()

```

### ServiceAppNameFixtureMapping

The `ServiceAppNameFixtureMapping` enum maps service names to application names.

#### Example

```python
from service_fixture import ServiceAppNameFixtureMapping

service_name = "video-downloader"
app_name = ServiceAppNameFixtureMapping.get_app_name(service_name)
print(app_name)  # Outputs: "downloader"
```

## Running Tests

To run the tests, use `pytest`:

```sh
npx nx test python-shared-testlib-fixtures-service-fixture
```

Make sure you have the development dependencies installed:

```sh
npx nx install python-shared-testlib-fixtures-service-fixture --with dev
```
