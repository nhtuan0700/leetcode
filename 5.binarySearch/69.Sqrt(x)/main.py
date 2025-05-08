# https://leetcode.com/problems/sqrtx/
# TC: O(logn), SC: O(1)
class Solution:
    def mySqrt(self, x: int) -> int:
        l, h = 1, x
        res = 0
        while l <= h:
            m = l + (h - l) // 2
            product = m * m
            # if product == x:
            #     return m
            if product <= x:
                l = m + 1
                res = m
            else:
                h = m - 1
        return res
