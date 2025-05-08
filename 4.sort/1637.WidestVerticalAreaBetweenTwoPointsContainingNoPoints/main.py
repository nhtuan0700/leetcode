from typing import List

# https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points
# Width of vertical area = difference between x-coordinates of two consecutive points (sorted by x-axis)
# TC: O(nlogn), SC: O(1)
class Solution:
    def maxWidthOfVerticalArea(self, points: List[List[int]]) -> int:
        points.sort(key=lambda x: x[0])
        return max(points[i+1][0] - points[i][0] for i in range(len(points) - 1))
