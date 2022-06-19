class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]):
        result = []

        if len(nums) < len(nums2):
            for n in nums1:
                if n in nums2:
                    result.append(n)
                    nums2.remove(n)
        else:
            for n in nums2:
                if n in nums1:
                    result.append(n)
                    nums1.remove(n)

        return result
            
