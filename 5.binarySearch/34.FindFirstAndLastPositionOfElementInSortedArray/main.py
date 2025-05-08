from typing import List

# https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

# TC: O(logn), SC: O(1)
class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        def findStart():
            nonlocal nums
            l, r = 0, len(nums) - 1
            res = -1
            while l <= r:
                m = l + (r-l) // 2
                if nums[m] >= target:
                    r = m - 1
                else:
                    l = m + 1
                if nums[m] == target:
                    res = m
            return res
        def findEnd():
            nonlocal nums
            l, r = 0, len(nums) - 1
            res = -1
            while l <= r:
                m = l + (r-l) // 2
                if nums[m] <= target:
                    l = m + 1
                else:
                    r = m - 1
                if nums[m] == target:
                    res = m
            return res

        start = findStart()
        end = findEnd()
        if start == -1:
            return [-1, -1]
        return [start, end]
