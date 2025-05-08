from typing import List
# https://leetcode.com/problems/contiguous-array/

# TC: O(n), SC: O(n)
class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        total = 0
        prefix_dict = {0: -1}
        res = 0
        n = len(nums)
        for i, num in enumerate(nums):
            total += 1 if num == 1 else -1
            if total in prefix_dict:
                res = max(res, i - prefix_dict[total])
            else:
                prefix_dict[total] = i

        return res
