from typing import List

class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        prefixSum = 0
        prefixMap = {0:1}
        res = 0
        for num in nums:
            prefixSum += num
            complete = prefixSum - k
            if complete in prefixMap:
                res += prefixMap[complete]
            prefixMap[prefixSum] = prefixMap.get(prefixSum, 0) + 1
        return res
