from typing import List
from collections import deque


# https://leetcode.com/problems/daily-temperatures/description/
# TC: O(n), SC: O(n)
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = deque()  # stack to keep indices of days
        n = len(temperatures)
        res = [0] * n
        for i in range(n - 1, -1, -1): # traverse from right to left
            while stack and temperatures[stack[-1]] <= temperatures[i]:
                stack.pop()
            if stack:
                res[i] = stack[-1] - i
            stack.append(i)
        return res
