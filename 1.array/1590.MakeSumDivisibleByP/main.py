from typing import List

# [total - sum(i,j)] % p = 0
# [total - sum(j) + sum(i-1)] % p = 0
# [sum(j) - total] % p = sum(i-1) % p
# [sum(j) - total + p] % p = sum(i - 1) % p
# (sum(j)%p - total%p + p) % p = sum(i-1) % p

class Solution:
    def minSubarray(self, nums: List[int], p: int) -> int:
        total = sum(nums) % p
        if total == 0:
            return 0
        modMap = {0:-1}
        res = len(nums)
        prefixSum = 0
        for i, num in enumerate(nums):
            prefixSum = (prefixSum + num) % p
            mod = (prefixSum - total + p) % p
            if mod in modMap:
                res = min(res, i - modMap[mod])
            modMap[prefixSum] = i
        
        if res == len(nums):
            return -1
        return res
