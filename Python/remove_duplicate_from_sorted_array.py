class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        idx_to_change = 0
        uniq_num = -101
        k = 0
        for i in range(len(nums)):
            if nums[i] > uniq_num:
                nums[idx_to_change] = nums[i]
                idx_to_change += 1
                uniq_num = nums[i]
                k += 1
        return k
