class Solution:
    def digitSum(self, s: str, k: int) -> str:
        def sumGroup(s):
            total = 0
            for c in s:
                total += ord(c) - ord('0')
            return str(total)


        while len(s) > k:
            tmp = ""
            i = 0
            n = len(s)
            while i < n:
                tmp += sumGroup(s[i:i+k])
                i += k
            s = tmp

        return s
