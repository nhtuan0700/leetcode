from typing import List
# https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

# TC: O(nlogn), SC: O(n)
class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        n = len(nums)
        sortedNums = nums.copy()
        for (i, num) in enumerate(sortedNums):
            sortedNums[i] = (num, i)
        
        sortedNums.sort(key=lambda x:x[0])
        res = [0] * n
        i = 0
        while i < n:
            (num, idx) = sortedNums[i]
            res[idx] = i
            j = i
            while j + 1 < n and sortedNums[j+1][0] == num:
                res[sortedNums[j+1][1]] = i
                j += 1
            i = j + 1
        
        return res
