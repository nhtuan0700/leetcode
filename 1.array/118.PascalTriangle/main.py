from typing import List

# O(n^2), SC: O(n^2)
# n: numRows
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        rows = []
        for i in range(0, numRows):
            rows.append([1] * (i+1))
            for j in range(1, i):
                rows[i][j] = rows[i-1][j] + rows[i-1][j-1]
        return rows
        