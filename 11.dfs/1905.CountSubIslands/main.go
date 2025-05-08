package main

// https://leetcode.com/problems/count-sub-islands/

var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}

// dfs for grid2
// during dfs each cell
// - make sure every cell-1 in grid2 are also in grid1
// - if having at least one cell-1 in grid2 is in grid1 (cell = 0) => result of dfs is false

// TC: O(m*n), SC: O(m*n)
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	n := len(grid1[0])

	vis := make([][]bool, m)
	for i := range m {
		vis[i] = make([]bool, n)
	}

	count := 0
	for i := range m {
		for j := range n {
			if vis[i][j] || grid2[i][j] == 0 {
				continue
			}

			if dfs(i, j, m, n, grid1, grid2, vis) {
				count++
			}
		}
	}

	return count
}

// check if the grid2 is sub-island of grid 1
func dfs(x, y, m, n int, grid1, grid2 [][]int, vis [][]bool) bool {
	vis[x][y] = true
	isSub := grid1[x][y] == 1

	for k := range 4 {
		u := x + dx[k]
		v := y + dy[k]

		if !inside(u, v, m, n) || vis[u][v] || grid2[u][v] == 0 {
			continue
		}

		isSub = dfs(u, v, m, n, grid1, grid2, vis) && isSub
	}

	return isSub
}

func inside(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
