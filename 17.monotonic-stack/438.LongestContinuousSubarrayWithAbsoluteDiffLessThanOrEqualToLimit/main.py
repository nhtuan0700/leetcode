from typing import List
from collections import deque

# https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
# TC: O(n), SC: O(n)
class Solution:
    def longestSubarray(self, nums: List[int], limit: int) -> int:
        n = len(nums)
        res = 0
        j = -1
        mind = deque()
        maxd = deque()
        for i in range(n):
            while j + 1 < n:
                if mind and nums[j + 1] - nums[mind[0]] > limit:
                    break
                if maxd and nums[maxd[0]] - nums[j + 1] > limit:
                    break

                j += 1
                while mind and nums[mind[-1]] >= nums[j]:
                    mind.pop()
                mind.append(j)
                while maxd and nums[maxd[-1]] <= nums[j]:
                    maxd.pop()
                maxd.append(j)

            res = max(res, j - i + 1)
            if mind[0] == i:
                mind.popleft()
            if maxd[0] == i:
                maxd.popleft()

        return res

# 1. brute force
# i: 0 -> n; j: i -> n
# O(n^2)
class Solution:
    def longestSubarray(self, nums: List[int], limit: int) -> int:
        res = 0
        for i in range(len(nums)):
            maxd = nums[i]
            mind = nums[i]
            for j in range(i, len(nums)):
                mind = min(mind, nums[j])
                maxd = max(maxd, nums[j])
                if maxd - mind > limit:
                    mind = nums[j]
                    maxd = nums[j]
                    break
        
                res = max(res, j - i + 1)
        return res
