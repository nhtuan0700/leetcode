from typing import List

# https://leetcode.com/problems/find-the-middle-index-in-array/

# TC: O(N), SC:O(N)
class Solution:
    def findMiddleIndex(self, nums: List[int]) -> int:
        total = sum(nums)
        leftSum = 0
        for i, num in enumerate(nums):
            if leftSum == total - leftSum - num:
                return i
            leftSum += num

        return -1
