class Solution:
    def reverseVowels(self, s: str) -> str:
        vowels = 'aiueoAIUEO'
        vow_stack = []
        new_s = ''

        for c in s:
            if c in vowels:
                vow_stack.append(c)

        for i in range(len(s)):
            if s[i] in vowels:
                new_s += vow_stack.pop()
            else:
                new_s += s[i]

        return new_s

