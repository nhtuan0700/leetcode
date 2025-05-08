package main

// https://leetcode.com/problems/detect-cycles-in-2d-grid/

//	var dirs = [4][2]int{
//		{-1, 0},
//		{0, 1},
//		{1, 0},
//		{0, -1},
//	}
var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}

// TC: O(m*n) SC: O(m*n)
func containsCycle(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])

	vis := make([][]bool, m)
	for i := range m {
		vis[i] = make([]bool, n)
	}

	for i := range m {
		for j := range n {
			if vis[i][j] {
				continue
			}

			if dfs(i, j, i, j, m, n, grid, vis) {
				return true
			}

		}
	}

	return false
}

func dfs(pi, pj, ci, cj, m, n int, grid [][]byte, vis [][]bool) bool {
	vis[ci][cj] = true

	for k := range 4 {
		i2 := ci + dx[k]
		j2 := cj + dy[k]

		// skip:
		// - outbound of grid
		// - not same value to current cell
		// - last visited cell
		if !inside(i2, j2, m, n) || grid[i2][j2] != grid[ci][cj] || (i2 == pi && j2 == pj) {
			continue
		}

		if vis[i2][j2] {
			return true
		}

		if dfs(ci, cj, i2, j2, m, n, grid, vis) {
			return true
		}
	}

	return false
}

func inside(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
