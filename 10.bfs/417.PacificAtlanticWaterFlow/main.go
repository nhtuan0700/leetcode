package main

// https://leetcode.com/problems/pacific-atlantic-water-flow/description/

// Find path that water can flow to pacific
// Find path that water can flow to atlantic
// - For each cell, bfs if current cell <= neighbors
// result's cell is in pacific & atlantic 
var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// TC: O(m*n), SC: O(m*n)
func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])

	pacific := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
	}
	atlantic := make([][]bool, m)
	for i := 0; i < m; i++ {
		atlantic[i] = make([]bool, n)
	}

	// vertical
	for i := 0; i < m; i++ {
		bfs(heights, i, 0, pacific)
		bfs(heights, i, n-1, atlantic)
	}

	// horizontal
	for i := 0; i < n; i++ {
		bfs(heights, 0, i, pacific)
		bfs(heights, m-1, i, atlantic)
	}

	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func bfs(heights [][]int, i, j int, vis [][]bool) {
	m := len(heights)
	n := len(heights[0])
	q := [][2]int{}
	q = append(q, [2]int{i, j})
	vis[i][j] = true

	head := 0
	for head < len(q) {
		val := q[head]
		head++
		i, j := val[0], val[1]

		for _, d := range dirs {
			i2 := i + d[0]
			j2 := j + d[1]

			if inside(i2, j2, m, n) && !vis[i2][j2] && heights[i2][j2] >= heights[i][j] {
				vis[i2][j2] = true
				q = append(q, [2]int{i2, j2})
			}
		}
	}
}

func inside(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
