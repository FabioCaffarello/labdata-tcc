import unittest
from pywarlock.serializer import serialize_to_dataclass


class TestSerializer(unittest.TestCase):

    def setUp(self):
        self.schema = {
            "name": "Person",
            "properties": {
                "name": {"type": "string"},
                "age": {"type": "integer"},
                "address": {
                    "type": "object",
                    "properties": {
                        "street": {"type": "string"},
                        "city": {"type": "string"}
                    }
                }
            }
        }
        self.data = {
            "name": "John Doe",
            "age": 30,
            "address": {
                "street": "123 Main St",
                "city": "Anytown"
            }
        }

    def test_serialize_to_dataclass(self):
        Person = serialize_to_dataclass(self.schema, self.data)
        self.assertEqual(Person.name, self.data["name"])
        self.assertEqual(Person.age, self.data["age"])
        self.assertEqual(Person.address["street"], self.data["address"]["street"])
        self.assertEqual(Person.address["city"], self.data["address"]["city"])

    def test_missing_property(self):
        incomplete_data = {
            "name": "John Doe",
            "address": {
                "street": "123 Main St"
                # Missing 'city'
            }
        }
        Person = serialize_to_dataclass(self.schema, incomplete_data)
        self.assertEqual(Person.name, incomplete_data["name"])
        self.assertEqual(Person.address["street"], incomplete_data["address"]["street"])
        self.assertIsNone(Person.address.get("city"))  # 'city' should be None


if __name__ == "__main__":
    unittest.main()
