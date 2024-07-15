import unittest
from pyargparse.argsparser import new


class TestArgumentParser(unittest.TestCase):
    """
    Test suite for the new function creating an argparse.ArgumentParser instance.
    """

    def test_parser_defaults(self):
        """
        Test that the parser sets the default values correctly.
        """
        parser = new("Test description")
        args = parser.parse_args([])

        self.assertFalse(args.enable_debug_storage)
        self.assertEqual(args.debug_storage_dir, "/app/tests/debug/storage")

    def test_parser_enable_debug_storage(self):
        """
        Test that the parser sets enable_debug_storage correctly.
        """
        parser = new("Test description")
        args = parser.parse_args(["--enable-debug-storage"])

        self.assertTrue(args.enable_debug_storage)

    def test_parser_debug_storage_dir(self):
        """
        Test that the parser sets debug_storage_dir correctly.
        """
        parser = new("Test description")
        test_dir = "/custom/dir"
        args = parser.parse_args(["--debug-storage-dir", test_dir])

        self.assertEqual(args.debug_storage_dir, test_dir)

    def test_parser_description(self):
        """
        Test that the parser sets the description correctly.
        """
        description = "Custom parser description"
        parser = new(description)

        self.assertEqual(parser.description, description)


if __name__ == '__main__':
    unittest.main()
