from typing import List

# https://leetcode.com/problems/continuous-subarray-sum/

# TC: O(n), SC: O(n)
class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        modIndexMap = {0: -1}
        modPrefixSum = 0
        for i, num in enumerate(nums):
            modPrefixSum = (modPrefixSum + num) % k
            mod = modPrefixSum % k
            if mod in modIndexMap and i - modIndexMap[mod] >= 2:
                return True
            modIndexMap.setdefault(mod, i)
        return False
        