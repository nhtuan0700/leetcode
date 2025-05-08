from typing import List

# https://leetcode.com/problems/sum-of-subarray-minimums/

# brute force
# O(n^2)
class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        n = len(arr)
        res = 0
        for i in range(n):
            minNum = arr[i]
            for j in range(i, n):
                minNum = min(arr[j], minNum)
                res = (res + minNum) % (1e9+7)
        return int(res)

# the number of arr[i] in result:
# = arr[i] * number of subarrays where arr[i] is the minimum
# = arr[i] * (left+1) * (right+1)
# left: arr[left_j] > arr[i]
# right: arr[right_j] >= arr[i]

# _ _ i _  => (2+1)*(1+1) = 6
class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        stack = []  # (idx, count)
        res = 0
        mod = 1e9 + 7
        for i, num in enumerate(arr):
            count = 1
            while stack and arr[stack[-1][0]] > num:
                v = stack.pop()
                count += v[1]
                res = (res + v[1] * (i - v[0]) * arr[v[0]]) % mod
            stack.append((i, count))

        while stack:
            v = stack.pop()
            res = (res + v[1] * (len(arr) - v[0]) * arr[v[0]]) % mod

        return int(res)
