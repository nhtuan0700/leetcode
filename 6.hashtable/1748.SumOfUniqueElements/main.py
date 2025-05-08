from typing import List

# https://leetcode.com/problems/sum-of-unique-elements

class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        freqs = [0] * 101
        for num in nums:
            freqs[num] += 1
        return sum([i for i in range(len(freqs)) if freqs[i] == 1])
