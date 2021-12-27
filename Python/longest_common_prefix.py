class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if len(strs) == 1:
            return strs[0]
        
        longest_common_prefix = ""
        start_ch = 0
        end_ch = 1
        while end_ch <= len(strs[0]):
            temp_str = strs[0][0:end_ch]
            is_common_prefix = True
            for s in strs:
                if temp_str not in s[0:end_ch]:
                    is_common_prefix = False
                    break
            if is_common_prefix and len(temp_str) > len(longest_common_prefix):
                longest_common_prefix = temp_str
                end_ch += 1
            else:
                break
        return longest_common_prefix
