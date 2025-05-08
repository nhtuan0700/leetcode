from typing import List

# https://leetcode.com/problems/subarray-sums-divisible-by-k/description/

# sum(i, j) % k == 0
# sum(j)%k = sum(i-1)%k

# TC: O(n), SC: O(n)
class Solution:
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        res = 0
        modMap = {0:1}
        modPrefix = 0
        for num in nums:
            modPrefix = (modPrefix + num) %k
            if modPrefix in modMap:
                res += modMap[modPrefix]

            modMap[modPrefix] = modMap.get(modPrefix, 0) + 1

        return res
        