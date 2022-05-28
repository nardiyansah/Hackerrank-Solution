class Solution:
    def count1(self, n: int):
        count = 0
        while n:
            count += n & 1
            n >>= 1
        return count

    def countBits(self, n: int) -> List[int]:
        result = [self.count1(i) for i in range(n + 1)]
        return result