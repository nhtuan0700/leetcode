from typing import List

# TC: O(logm + logn) = O(log(m*n)), SC: O(1)
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        if target < matrix[0][0] or target > matrix[m-1][n-1]:
            return False
        # row
        l, r = 0, m - 1
        row = 0
        while l <= r:
            mid = l + (r - l) // 2
            if matrix[mid][0] > target:
                r = mid - 1
            else:
                row = mid
                l = mid + 1
        l, r = 0, n - 1
        while l <= r:
            mid = l + (r - l) // 2
            if matrix[row][mid] == target:
                return True
            if matrix[row][mid] > target:
                r = mid - 1
            else:
                l = mid + 1


        return False

