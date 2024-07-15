import argparse


def new(description: str) -> argparse.ArgumentParser:
    """
    Creates a new ArgumentParser with specified description and arguments.

    Args:
        description (str): Description for the argument parser.

    Returns:
        argparse.ArgumentParser: Configured ArgumentParser instance.
    """
    parser = argparse.ArgumentParser(description=description)

    parser.add_argument(
        "--enable-debug-storage",
        help="Enable debug module",
        default=False,
        action="store_true"
    )

    parser.add_argument(
        "--debug-storage-dir",
        help="Sets the base directory for debug storage",
        dest="debug_storage_dir",
        default="/app/tests/debug/storage"
    )

    return parser
