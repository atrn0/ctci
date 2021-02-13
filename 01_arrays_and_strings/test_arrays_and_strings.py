import unittest
from arrays_and_strings import *


class Testarrays_and_strings(unittest.TestCase):
    """test class of arrays_and_strings.py
    """

    def test_IsUnique(self):
        self.assertEqual(IsUnique("asdfg"), True)
        self.assertEqual(IsUnique("asafg"), False)
        self.assertEqual(IsUnique("aaaa"), False)
        self.assertEqual(IsUnique(""), True)

    def test_CheckPermutation(self):
        self.assertEqual(CheckPermutation("asdfg","sdfag"), True)
        self.assertEqual(CheckPermutation("asdfg","sdfsg"), False)
        self.assertEqual(CheckPermutation("asdfg","sdag"), False)

    def test_URLify(self):
        self.assertEqual(URLify("Mr John Smith    ", 13), "Mr%20John%20Smith")
        self.assertEqual(URLify("I am a student  ", 14), "I%20am%20a%20student")

    def test_PalindromePermutation(self):
        self.assertEqual(PalindromePermutation("Tact Coa"), True)
        self.assertEqual(PalindromePermutation("Adcs sd"), False)


if __name__ == "__main__":
    unittest.main()