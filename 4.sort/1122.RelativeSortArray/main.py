from typing import List
from functools import cmp_to_key

# https://leetcode.com/problems/relative-sort-array/
# counting sort
# TC: O(1000) ~ O(1), SC: O(1001) ~ O(1)
class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        counts = [0] * 1001
        for num in arr1:
            counts[num] += 1
        idx = 0
        for num in arr2:
            while counts[num] > 0:
                arr1[idx] = num
                counts[num] -= 1
                idx += 1
        for i in range(len(counts)):
            while counts[i] > 0:
                arr1[idx] = i
                counts[i] -= 1
                idx += 1

        return arr1

# TC: O(n2logn2), SC: O(1)
# n2: len(arr2)
class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        criticalSort = {}
        for (i, num) in enumerate(arr2):
            criticalSort[num] = i
        
        def compare(a, b):
            nonlocal criticalSort
            if a in criticalSort and b in criticalSort:
                return criticalSort[a] - criticalSort[b]
            # b not in arr2
            if a in criticalSort:
                return -1
            if b in criticalSort:
                return 1
            return a - b

        arr1.sort(key=cmp_to_key(compare))
        return arr1
