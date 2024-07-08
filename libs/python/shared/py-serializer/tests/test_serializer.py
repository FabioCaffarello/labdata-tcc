import unittest
from pyserializer import serializer
from dataclasses import dataclass, field
from typing import List


@dataclass
class JobDependenciesDTO:
    service: str
    source: str


@dataclass
class ConfigDTO:
    id: str = field(metadata={"json": "_id"})
    active: bool
    service: str
    source: str
    provider: str
    depends_on: List[JobDependenciesDTO]
    config_version_id: str
    created_at: str = field(repr=False)
    updated_at: str = field(repr=False)


@dataclass
class Address:
    street: str
    city: str


@dataclass
class Person:
    name: str
    age: int
    address: Address


class TestSerializer(unittest.TestCase):

    def setUp(self):
        self.address = Address(street="123 Main St", city="Anytown")
        self.person = Person(name="John Doe", age=30, address=self.address)

        self.job_dependencies = JobDependenciesDTO(service="dep-service", source="dep-source")
        self.config = ConfigDTO(
            id="test_id",
            active=True,
            service="test-service",
            source="test-source",
            provider="provider",
            depends_on=[self.job_dependencies],
            config_version_id="xyz123",
            created_at="2024-02-01 00:00:00",
            updated_at="2024-02-01 00:00:00"
        )

    def test_serialize_person_to_json(self):
        json_str = serializer.serialize_to_json(self.person)
        expected_json = '{"address": {"city": "Anytown", "street": "123 Main St"}, "age": 30, "name": "John Doe"}'
        self.assertEqual(json_str, expected_json)

    def test_serialize_person_to_dict(self):
        dict_obj = serializer.serialize_to_dict(self.person)
        expected_dict = {
            "name": "John Doe",
            "age": 30,
            "address": {
                "street": "123 Main St",
                "city": "Anytown"
            }
        }
        self.assertEqual(dict_obj, expected_dict)

    def test_serialize_person_to_dataclass(self):
        data = {
            "name": "John Doe",
            "age": 30,
            "address": {
                "street": "123 Main St",
                "city": "Anytown"
            }
        }
        person_obj = serializer.serialize_to_dataclass(data, Person)
        self.assertEqual(person_obj, self.person)

    def test_serialize_config_to_json(self):
        json_str = serializer.serialize_to_json(self.config)
        expected_json = (
            '{"_id": "test_id", "active": true, "config_version_id": "xyz123", "created_at": "2024-02-01 00:00:00", '
            '"depends_on": [{"service": "dep-service", "source": "dep-source"}], "provider": "provider", '
            '"service": "test-service", "source": "test-source", "updated_at": "2024-02-01 00:00:00"}'
        )
        self.assertEqual(json_str, expected_json)

    def test_serialize_config_to_dict(self):
        dict_obj = serializer.serialize_to_dict(self.config)
        expected_dict = {
            "_id": "test_id",
            "active": True,
            "config_version_id": "xyz123",
            "created_at": "2024-02-01 00:00:00",
            "depends_on": [{"service": "dep-service", "source": "dep-source"}],
            "provider": "provider",
            "service": "test-service",
            "source": "test-source",
            "updated_at": "2024-02-01 00:00:00"
        }
        self.assertEqual(dict_obj, expected_dict)

    def test_serialize_config_to_dataclass(self):
        data = {
            "_id": "test_id",
            "active": True,
            "config_version_id": "xyz123",
            "created_at": "2024-02-01 00:00:00",
            "depends_on": [{"service": "dep-service", "source": "dep-source"}],
            "provider": "provider",
            "service": "test-service",
            "source": "test-source",
            "updated_at": "2024-02-01 00:00:00"
        }
        config_obj = serializer.serialize_to_dataclass(data, ConfigDTO)
        self.assertEqual(config_obj, self.config)


if __name__ == "__main__":
    unittest.main()
