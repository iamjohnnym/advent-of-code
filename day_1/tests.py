import unittest
import run as day_1


class TestPartOne(unittest.TestCase):
    def test_equal_three(self):
        result = day_1.resulting_frequency([+1, +1, +1])
        self.assertEqual(3, result)

    def test_equal_zero(self):
        result = day_1.resulting_frequency([+1, +1, -2])
        self.assertEqual(0, result)

    def test_equal_negative_six(self):
        result = day_1.resulting_frequency([-1, -2, -3])
        self.assertEqual(-6, result)


class TestPartTwo(unittest.TestCase):

    def test_equal_zero(self):
        result = day_1.first_repeated_frequency([+1, -1])
        self.assertEqual(0, result)

    def test_equal_ten(self):
        result = day_1.first_repeated_frequency([+3, +3, +4, -2, -4])
        self.assertEqual(10, result)

    def test_equal_five(self):
        result = day_1.first_repeated_frequency([-6, +3, +8, +5, -6])
        self.assertEqual(5, result)

    def test_equal_fourteen(self):
        result = day_1.first_repeated_frequency([+7, +7, -2, -7, -4])
        self.assertEqual(14, result)


if __name__ == '__main__':
    unittest.main()
