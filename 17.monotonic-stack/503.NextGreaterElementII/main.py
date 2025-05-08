from typing import List


# idea: dup array size
# https://leetcode.com/problems/next-greater-element-ii/
# TC: O(n), SC: O(n)
class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        n = len(nums)
        nums = nums * 2
        stack = []
        res = [-1] * n
        for i in range(n*2 - 1, -1, -1):
            while stack and stack[-1] <= nums[i]:
                stack.pop()
            if i < n:
                res[i] =  stack[-1] if stack else -1
            stack.append(nums[i])
        return res
    
    def nextGreaterElements2(self, nums: List[int]) -> List[int]:
        n = len(nums)
        stack = []
        res = [-1] * n
        for i in range(n*2 - 1, -1, -1):
            i = i % n
            while stack and stack[-1] <= nums[i]:
                stack.pop()
            if i < n:
                res[i] = stack[-1] if stack else -1
            stack.append(nums[i])
        return res

