class Solution:
    def romanToInt(self, s: str) -> int:
        roman = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
        integer_number = 0
        list_char = list(s)
        for i in range(len(list_char)):
            if (
                i != (len(list_char) - 1)
                and roman[list_char[i]] >= roman[list_char[i + 1]]
            ):
                integer_number += roman[list_char[i]]
            elif i != (len(list_char) - 1):
                integer_number -= roman[list_char[i]]
            else:
                integer_number += roman[list_char[i]]
        return integer_number
