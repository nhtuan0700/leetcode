from collections import defaultdict
from typing import List


#https://leetcode.com/problems/valid-sudoku/

# TC: O(n^2) SC: O(n^2)
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # dict[num][row,col]
        numMap = defaultdict(list)
        for i in range(9):
            for j in range(9):
                num = board[i][j]
                if num == '.':
                    continue
                
                for (row, col) in numMap[num]:
                    lowerBoundRow = i // 3 * 3
                    lowerBoundCol = j // 3 * 3

                    if i == row or j == col or (lowerBoundRow <= row and row <= lowerBoundRow + 2 and lowerBoundCol <= col and col <= lowerBoundCol + 2):
                        return False
                
                numMap[num].append((i, j))
        return True

class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # dict[num][row,col]
        rows = defaultdict(set)
        cols = defaultdict(set)
        boxes = defaultdict(set)
        for r in range(9):
            for c in range(9):
                num = board[r][c]
                if num == '.':
                    continue
                
                if num in rows[r] or num in cols[c] or num in boxes[(r//3*3, c//3*3)]:
                    return False
                rows[r].add(num)
                cols[c].add(num)
                boxes[(r//3*3, c//3*3)].add(num)
        return True
