from typing import List

# https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

# TC: O(n), SC: O(1)
class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        n = len(nums) + 1
        for i, num in enumerate(nums):
            idx = abs(num) - 1
            if nums[idx] > 0:
                nums[idx] = -nums[idx]
        
        res = []
        for i in range(1, n):
            if nums[i-1] > 0:
                res.append(i)
        return res

# Cyclic sort: swap nums[i], nums[nums[i]-1] until nums[i] = nums[nums[i]-1]
# TC: O(n), SC: O(n)
class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        n = len(nums)
        for i in range(n):
            while nums[i] != nums[nums[i]-1]:
                # wrong:  nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
                nums[nums[i]-1], nums[i]  = nums[i], nums[nums[i]-1]

        return [i+1 for i in range(n) if nums[i] != i+1]
