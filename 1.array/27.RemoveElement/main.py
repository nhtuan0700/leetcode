from typing import List

# https://leetcode.com/problems/remove-element/

# TC: O(n), SC: O(1)
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        k = 0
        for num in nums:
            if num != val:
                nums[k] = num
                k += 1
        
        return k
