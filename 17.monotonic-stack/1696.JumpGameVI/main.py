from typing import List
import math
from collections import deque

# https://leetcode.com/problems/jump-game-vi/

# Approach 1: Naive Dynamic Programming
# O(n*k)
class Solution:
    def maxResult(self, nums: List[int], k: int) -> int:
        for i in range(1, len(nums)): # O(n)
            bestScore = -math.inf
            for j in range(max(0, i-k), i): # O(k)
                bestScore = max(bestScore, nums[j] + nums[i])
            nums[i] = bestScore
        return nums[-1]

# Approach 2: Dynamic programing + decreasing queue

# dp[i] is the maximum score we can get when ending at index i
# dp[i] = max(dp[i-k], dp[i-k+1], ..., dp[i-1])
class Solution:
    def maxResult(self, nums: List[int], k: int) -> int:
        n = len(nums)
        dq = deque([0])
        for i in range(1, n):
            nums[i] = nums[dq[0]] + nums[i]
            while dq and nums[dq[-1]] <= nums[i]:
                dq.pop()
            dq.append(i)
            if dq[0] <= i - k:
                dq.popleft()
                
        return nums[-1]
