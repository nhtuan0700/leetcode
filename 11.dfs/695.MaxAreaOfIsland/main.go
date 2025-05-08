package main

// https://leetcode.com/problems/max-area-of-island/

var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}

// TC: O(m*n) SC: O(m*n)
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	res := 0
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if vis[i][j] || grid[i][j] == 0 {
				continue
			}

			res = max(res, dfs(i, j, m, n, grid, vis))
		}
	}

	return res
}

// get area of island
func dfs(x, y, m, n int, grid [][]int, vis [][]bool) int {
	vis[x][y] = true

	area := 1
	for k := 0; k < 4; k++ {
		u := x + dx[k]
		v := y + dy[k]

		if !inside(u, v, m, n) || vis[u][v] || grid[u][v] == 0 {
			continue
		}

		area += dfs(u, v, m, n, grid, vis)
	}

	return area
}

func inside(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
