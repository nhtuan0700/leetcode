from typing import List

# https://leetcode.com/problems/subarray-sum-equals-k/

# 1,1,1,-3,2,1,-1,2
# 1,2,3,0,2,3,2,4
# 1 -> occurences {0: 1, 1: 1} cnt = 0
# 2 -> occurences {0: 1, 1: 1, 2: 1} cnt = 1
# 3 -> occurences {0: 1, 1: 1, 2: 1, 3: 1} cnt = 2
# 0 -> occurences {0: 2, 1: 1, 2: 1, 3: 1} cnt = 2
# 2 -> occurences {0: 2, 1: 1, 2: 2, 3: 1} cnt = 4
# 3 -> occurences {0: 2, 1: 1, 2: 2, 3: 2} cnt = 5
# 2 -> occurences {0: 2, 1: 1, 2: 3, 3: 2} cnt = 7
# 4 -> occurences {0: 2, 1: 1, 2: 3, 3: 2, 4:2} cnt = 10

# sum(j,i) = sum(0,i) - sum(0,j-1) = k
# sum(0,j-1) = sum(0,i) - k
# with j < i

# 2. at each step i (0->n) in prefixSum:
#   prefixSum
#   find occurence number of sum(0,j-1) => use hashMap

# TC: O(n) SC: O(n)
class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        occurrences = {0: 1}
        cnt = 0
        prefixSum = 0
        for num in nums:
            prefixSum += num
            cnt += occurrences.get(prefixSum - k, 0)
            occurrences[prefixSum] = occurrences.get(prefixSum, 0) + 1
        return cnt
