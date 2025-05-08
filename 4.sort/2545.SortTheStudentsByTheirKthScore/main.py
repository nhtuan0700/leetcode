from typing import List

# https://leetcode.com/problems/sort-the-students-by-their-kth-score/

# TC: O(mlogm), SC: O(1)
class Solution:
    def sortTheStudents(self, score: List[List[int]], k: int) -> List[List[int]]:
        score.sort(key=lambda x: x[k], reverse=True)
        return score
