from typing import List

# https://leetcode.com/problems/prime-in-diagonal/

# every n x n matrix, there are only 2 diagonals
# nums[i][j] is on one of the diaganols if i == j or j = n - i - 1

# approach:
# - get cells in diagonals
# - check max prime number for every cell in diagonal cells


# number cell of diagonals
# 1x1: 1
# 2x2: 4
# 3x3: 5
# 4x4: 8
# ...
# nxn: 2*n - n%2
# TC: O(n*sqrt(max(nums))), SC: O(n)
class Solution:
    def diagonalPrime(self, nums: List[List[int]]) -> int:
        def isPrime(n) -> bool:
            if n < 2:
                return False
            for i in range(2, int(n**0.5) + 1):
                if n % i == 0:
                    return False
            return True

        n = len(nums)
        diagonals = set()
        for i in range(0, n):
            diagonals.add(nums[i][i])
            diagonals.add(nums[i][n - i - 1])

        res = 0
        for num in diagonals:
            if isPrime(num):
                res = max(res, num)
        return res
