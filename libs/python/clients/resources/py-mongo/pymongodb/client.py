from pymongo import MongoClient
from pysd.sd import new_from_env


def get_mongo_client() -> MongoClient:
    """
    Retrieves a MongoDB client connected to the MongoDB endpoint defined in the environment.

    Returns:
        MongoClient: An instance of MongoClient connected to the MongoDB endpoint.
    """
    mongo_url = new_from_env().mongodb_endpoint
    return MongoClient(mongo_url)


def drop_database(database_name: str) -> None:
    """
    Drops the specified database from the MongoDB server.

    Args:
        database_name (str): The name of the database to drop.

    Returns:
        None
    """
    mongo_client = get_mongo_client()
    mongo_client.drop_database(database_name)
