from typing import List

# https://leetcode.com/problems/valid-perfect-square/

# TC: O(log(num)), SC: O(1)
class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        l, r = 1, num
        while l <= r:
            m = l + (r - l)//2
            square = m*m
            if square == num:
                return True
            if square > num:
                r = m - 1
            else:
                l = m + 1
        return False
