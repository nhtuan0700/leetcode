from typing import List

# https://leetcode.com/problems/sum-of-subarray-minimums/

# brute force
# O(n^2)
class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        n = len(arr)
        res = 0
        for i in range(n):
            minNum = arr[i]
            for j in range(i, n):
                minNum = min(arr[j], minNum)
                res = (res + minNum) % (1e9+7)
        return int(res)

# 3 1 2 4
# 
