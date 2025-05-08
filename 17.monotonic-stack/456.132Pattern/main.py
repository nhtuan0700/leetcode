from typing import List

# https://leetcode.com/problems/132-pattern/

# for each nums[k] from left to right, try to:
# - Find j < k such that nums[j] > nums[k] 
#   and i < j such that nums[i] < nums[k] 
# => use monotonic decreasing stack to store (nums[j], min_before_j)
#    where min_before_j keep track the minimum number that appears before nums[j]
# TC: O(n), SC: O(n)
class Solution:
    def find132pattern(self, nums: List[int]) -> bool:
        stack = [] # (num, min)
        curMin = nums[0]

        for n in nums[1:]:
            while stack and stack[-1][0] <= n:
                stack.pop()
            # now stack is empty or top num of stack is greater than n
            if stack and stack[-1][1] < n:
                return True
            
            stack.append((n,curMin))
            curMin = min(curMin, stack[-1][0])
        return False
