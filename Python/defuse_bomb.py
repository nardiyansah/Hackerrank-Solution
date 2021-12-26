class Solution:
    def decrypt(self, code: List[int], k: int) -> List[int]:
        decrypt = []
        l = len(code)
        if k == 0:
            return [0] * l
        elif k > 0:
            for i in range(l):
                sum = 0
                for j in range(k):
                    sum += code[(i + j + 1) % l]
                decrypt.append(sum)
            return decrypt
        else:
            for i in range(l):
                sum = 0
                for j in range(-k):
                    sum += code[(i + l - 1 - j) % l]
                decrypt.append(sum)
            return decrypt
