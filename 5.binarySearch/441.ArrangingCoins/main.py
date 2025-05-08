# https://leetcode.com/problems/arranging-coins/

# 1 - 1
# 2 - 3
# 3 - 6
# 4 - 10
# => 1 + 2 + 3 + 4 = 10
# sum = n * (n + 1) / 2
# TC: O(logn), SC: O(1)
class Solution:
    def arrangeCoins(self, n: int) -> int:
        l, r = 1, n
        res = 0
        while l <= r:
            m = l + (r-l)//2
            s = m * (m + 1) // 2
            if n == s:
                return m
            if s < n:
                l = m + 1
                res = m
            else:
                r = m - 1
        return res
