package main

// https://leetcode.com/problems/surrounded-regions/description/

var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// TC: O(m*n), SC(m*n)
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	regions := make([][]bool, m)
	for i := 0; i < m; i++ {
		regions[i] = make([]bool, n)
	}

	q := make([][2]int, 0)
	// O(m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && isOnEdge(i, j, m, n) {
				regions[i][j] = true
				q = append(q, [2]int{i, j})
			}
		}
	}

	head := 0
	for head < len(q) {
		val := q[head]
		i, j := val[0], val[1]
		head++

		for _, d := range dirs {
			i2 := i + d[0]
			j2 := j + d[1]

			if inside(i2, j2, m, n) {
				if board[i2][j2] == 'X' {
					continue
				}
				if !regions[i2][j2] {
					regions[i2][j2] = true
					q = append(q, [2]int{i2, j2})
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !regions[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func inside(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func isOnEdge(i, j, m, n int) bool {
	return i == 0 || i == m-1 || j == 0 || j == n-1
}
