from typing import List

# https://leetcode.com/problems/height-checker/
# counting sort
# [2,1,1,4]
# [0,2,1,0,1]
# TC: O(101) ~ O(1), SC: O(101) ~ O(1)
class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        counts = [0] * 101
        for h in heights:
            counts[h] += 1
        curIdx = 0
        diff = 0
        for i in range(len(counts)):
            while counts[i] > 0:
                if heights[curIdx] != i:
                    diff += 1
                counts[i] -= 1
                curIdx += 1
        return diff

# O(nlogn), SC: O(n)
class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        sortedHeights = sorted(heights)
        diff = 0
        for i in range(len(heights)):
            if sortedHeights[i] != heights[i]:
                diff += 1
        return diff
