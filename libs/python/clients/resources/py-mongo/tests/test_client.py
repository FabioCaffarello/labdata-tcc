import unittest
from unittest.mock import MagicMock, patch

from pymongodb.client import drop_database, get_mongo_client


class TestMongoClient(unittest.TestCase):

    @patch('pymongodb.client.MongoClient')
    @patch('pymongodb.client.new_from_env')
    def test_get_mongo_client(self, mock_new_from_env, mock_mongo_client):
        mock_env_instance = MagicMock()
        mock_env_instance.mongodb_endpoint = 'mongodb://test_url:27017'
        mock_new_from_env.return_value = mock_env_instance

        client = get_mongo_client()

        mock_new_from_env.assert_called_once()
        mock_mongo_client.assert_called_once_with('mongodb://test_url:27017')
        self.assertEqual(client, mock_mongo_client.return_value)

    @patch('pymongodb.client.MongoClient')
    @patch('pymongodb.client.new_from_env')
    def test_drop_database(self, mock_new_from_env, mock_mongo_client):
        mock_env_instance = MagicMock()
        mock_env_instance.mongodb_endpoint = 'mongodb://test_url:27017'
        mock_new_from_env.return_value = mock_env_instance

        mock_client_instance = MagicMock()
        mock_mongo_client.return_value = mock_client_instance

        drop_database('test_db')

        mock_new_from_env.assert_called_once()
        mock_mongo_client.assert_called_once_with('mongodb://test_url:27017')
        mock_client_instance.drop_database.assert_called_once_with('test_db')


if __name__ == '__main__':
    unittest.main()
