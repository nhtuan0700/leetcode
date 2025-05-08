package main

// https://leetcode.com/problems/find-all-groups-of-farmland/

// traverse land if cell is not vis
// - add its coordinate as top-left corner
// - dfs to find right corner

var dirs = [2][2]int{
	{0, 1},
	{1, 0},
}

// TC: O(m*n), SC: O(m*n)
func findFarmland(land [][]int) [][]int {
	m := len(land)
	n := len(land[0])

	vis := make([][]bool, m)
	for i := range m {
		vis[i] = make([]bool, n)
	}

	res := make([][]int, 0)
	for i := range m {
		for j := range n {
			if vis[i][j] || land[i][j] == 0 {
				continue
			}
			bot, right := dfs(i, j, m, n, land, vis)
			res = append(res, []int{i, j, bot, right})
		}
	}

	return res
}

func dfs(i, j, m, n int, land [][]int, vis [][]bool) (int, int) {
	vis[i][j] = true
	// bot: bot, right: right
	bot, right := i, j

	for _, d := range dirs {
		dx, dy := d[0], d[1]
		i2 := i + dx
		j2 := j + dy

		if !inside(i2, j2, m, n) || vis[i2][j2] || land[i2][j2] == 0 {
			continue
		}

		x2, y2 := dfs(i2, j2, m, n, land, vis)
		bot = max(bot, x2)
		right = max(right, y2)
	}

	return bot, right
}

func inside(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
