package main

// https://leetcode.com/problems/search-a-2d-matrix/description/

// O(log(m*n))
// using bsearch
// - find the row
// - find the column
// return true if row >= 0, col <= len()
// similar: largest number < target
// TC: O(logm + logn) = O(log(m*n))
func searchMatrix1(matrix [][]int, target int) bool {
	startRow := 0
	endRow := len(matrix) - 1

	// O(logm)
	for startRow <= endRow {
		midRow := startRow + (endRow-startRow)/2
		firstCell := matrix[midRow][0]
		if firstCell == target {
			return true
		}
		if firstCell < target {
			startRow = midRow + 1
		} else {
			endRow = midRow - 1
		}
	}

	if endRow < 0 {
		return false
	}

	currentRow := endRow
	l := 0
	h := len(matrix[currentRow]) - 1
	// O(logn)
	for l <= h {
		m := l + (h-l)/2
		if matrix[currentRow][m] == target {
			return true
		}
		if matrix[currentRow][m] > target {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

// flat matrix
// TC: O(log(m*n))
func searchMatrix2(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if matrix[m-1][n-1] < target {
		return false
	}

	l := 0
	h := (m * n) - 1
	for l <= h {
		m := l + (h-l)/2
		row := m / n
		col := m % n
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return false
}
