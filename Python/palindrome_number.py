class Solution:
    def isPalindrome(self, x: int) -> bool:
        str_x = list(str(x))
        first_idx = 0
        last_idx = len(str_x) - 1
        while first_idx < last_idx:
            if str_x[first_idx] != str_x[last_idx]:
                return False
            first_idx += 1
            last_idx -= 1
        return True
