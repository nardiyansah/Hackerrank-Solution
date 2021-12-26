class Solution:
    def reorderLogFiles(self, logs: List[str]) -> List[str]:
        mirror_arr = []
        for e in logs:
            mirror_arr.append(e.split())
        return mirror_arr
