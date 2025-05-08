package main

// https://leetcode.com/problems/number-of-islands/

// up, right, down, left
var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}

// TC: O(m*n), SC: O(m*n)
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if vis[i][j] || grid[i][j] == '0' {
				continue
			}
			res++
			dfs(i, j, m, n, grid, vis)
		}
	}

	return res
}

func dfs(x, y, m, n int, grid [][]byte, vis [][]bool) {
	vis[x][y] = true

	// 4 directions
	for k := 0; k < 4; k++ {
		u := x + dx[k]
		v := y + dy[k]
		// skip out of bound cells
		// skip visited cell
		// skip invalid cell
		if !inside(u, v, m, n) || vis[u][v] || grid[u][v] == '0' {
			continue
		}

		dfs(u, v, m, n, grid, vis)
	}
}

func inside(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
