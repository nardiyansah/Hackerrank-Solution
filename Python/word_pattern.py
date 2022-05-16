class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        list_s = s.split()
        dict_pattern_word = dict()

        if len(pattern) != len(list_s):
            return False

        for idx, c in enumerate(pattern):
            if dict_pattern_word.get(c):
                if list_s[idx] != dict_pattern_word.get(c):
                    return False
            elif list_s[idx] in dict_pattern_word.values():
                return False
            else:
                dict_pattern_word[c] = list_s[idx]

        return True
