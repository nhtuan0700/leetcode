from typing import List

# https://leetcode.com/problems/search-insert-position/

# TC: O(logn), SC: O(1)
class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        l, h = 0, len(nums) - 1
        res = len(nums)
        while l <= h:
            m = l + (h - l) // 2
            if nums[m] >= target:
                h = m - 1
                res = m
            else:
                l = m + 1
        return res
