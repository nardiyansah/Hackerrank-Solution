class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        for i in range(1, num+1):
          square = i * i
          if square > num:
            return False
          if square == num:
            return True
