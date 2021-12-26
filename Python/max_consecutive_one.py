class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        max = 0
        temp = 0
        for i in nums:
            if i == 0:
                temp = 0
            else:
                temp += 1
            if temp > max:
                max = temp
        return max
