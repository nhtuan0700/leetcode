package main

// https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/

// up, right, down, left
var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// TC: O(m*n) SC: O(m*n)
func findMaxFish(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	vis := make([][]bool, m)
	for i := range m {
		vis[i] = make([]bool, n)
	}

	count := 0
	for i := range m {
		for j := range n {
			if vis[i][j] || grid[i][j] == 0 {
				continue
			}
			count = max(count, dfs(i, j, m, n, grid, vis))
		}
	}

	return count
}

func dfs(i, j, m, n int, grid [][]int, vis [][]bool) int {
	count := grid[i][j]
	vis[i][j] = true

	for _, d := range dirs {
		i2 := i + d[0]
		j2 := j + d[1]

		if !inside(i2, j2, m, n) || vis[i2][j2] || grid[i2][j2] == 0 {
			continue
		}

		count += dfs(i2, j2, m, n, grid, vis)
	}

	return count
}

func inside(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
