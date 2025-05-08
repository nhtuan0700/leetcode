package main

// https://leetcode.com/problems/valid-sudoku/description/z
// 9x9
// at each cell, we will check row, column, sub-box
// map[string][2]int{[row,column]}

// lowerBound = (i // 3) * 2, upperBound = lowerBound + 2
// invalid: i == row || j == column || lowerBound(i) <= row <= upperBound(i) || lowerBound(j) <= column <= upperBound(j)

func isValidSudoku(board [][]byte) bool {
	tempMap := make(map[byte][][]int)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}

			for _, vals := range tempMap[val] {
				row := vals[0]
				col := vals[1]

				lowerBoundRow := (i / 3) * 3
				lowerBoundCol := (j / 3) * 3
				if i == row || j == col ||
					(lowerBoundRow <= row && row <= lowerBoundRow+2 &&
						lowerBoundCol <= col && col <= lowerBoundCol+2) {
					return false
				}
			}
			tempMap[val] = append(tempMap[val], []int{i, j})
		}
	}

	return true
}
