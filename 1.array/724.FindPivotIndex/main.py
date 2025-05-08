from typing import List

class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        total = sum(nums)
        leftSum = 0
        for i, num in enumerate(nums):
            total -= num
            if total == leftSum:
                return i
            leftSum += num
        return -1
