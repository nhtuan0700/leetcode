# https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/

class Solution:
    def findKthBit(self, n: int, k: int) -> str:
        def helper(n, k, isInvert):
            if n == 1:
                return "1" if isInvert else "0"
            
            # (2^n - 1)
            l = (1 << n) - 1
            m = l // 2 + 1
            if k == m:
                return "0" if isInvert else "1"
            if k < m:
                return helper(n-1, k, isInvert)
            else:
                k = l - k + 1
                return helper(n-1, k, not isInvert)

        return helper(n, k, False)
