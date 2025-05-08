from typing import List

# https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

class Solution:
    def findMin(self, nums: List[int]) -> int:
        if nums[0] <= nums[-1]:
            return nums[0]
        n = len(nums)
        l, h = 0, n - 1
        while l < h:
            m = l + (h - l) // 2
            if m < n - 1 and nums[m+1] < nums[m]:
                return nums[m+1]
            if m > 0 and nums[m - 1] > nums[m]:
                return nums[m]
            if nums[m] >= nums[l]: # left sort
                if nums[l] > nums[h]:
                    l = m + 1
                else:
                    h = m - 1
            else:
                if nums[l] > nums[h]:
                    h = m - 1
                else:
                    l = m + 1
        return -1

class Solution:
    def findMin(self, nums: List[int]) -> int:
        if nums[0] <= nums[-1]:
            return nums[0]
        n = len(nums)
        l, h = 0, n - 1
        while l < h:
            m = l + (h - l) // 2
            if nums[m] > nums[h]:
                l = m + 1
            else:
                h = m
        return nums[l]

