from typing import List

# https://leetcode.com/problems/find-lucky-integer-in-an-array/

# TC: O(n), SC: O(1)
class Solution:
    def findLucky(self, arr: List[int]) -> int:
        freqs = [0] * 501
        for num in arr:
            freqs[num] += 1

        res = -1
        for num in arr:
            if num == freqs[num]:
                res = max(res, num)
        return res
