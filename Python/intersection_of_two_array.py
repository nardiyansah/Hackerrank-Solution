class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        result = []
        
        if len(nums1) < len(nums2):
            for n in nums1:
                if n in nums2 and n not in result
                    result.append(n)
        else:
            for n in num2:
                if n in nums1 and n not in result
                    result.append(n)

        return result

