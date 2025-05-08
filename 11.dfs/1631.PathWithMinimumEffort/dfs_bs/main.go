package main

// https://leetcode.com/problems/path-with-minimum-effort/

// binary search to find the minimum thresold k
// each k:
// - dfs to find possible path from (0,0) -> (m-1,n-1) with |diff| <= k
// If having path: descrease k
// otherwise: increase k

var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// TC: O(m*n*log(maxHeights)), SC: O(m*n)
func minimumEffortPath(heights [][]int) int {
	l, h := 0, 1000000
	ans := h

	for l <= h {
		k := l + (h-l)/2
		if canReach(heights, k) {
			ans = k
			h = k - 1
		} else {
			l = k + 1
		}
	}

	return ans
}

func canReach(heights [][]int, k int) bool {
	m := len(heights)
	n := len(heights[0])

	vis := make([][]bool, m)

	for i := range m {
		vis[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, heights, vis, k)
}

func dfs(i, j, m, n int, heights [][]int, vis [][]bool, k int) bool {
	if i == m-1 && j == n-1 {
		return true
	}

	vis[i][j] = true

	for _, d := range dirs {
		i2 := i + d[0]
		j2 := j + d[1]

		if !inside(i2, j2, m, n) || vis[i2][j2] {
			continue
		}

		diff := abs(heights[i][j] - heights[i2][j2])
		if diff <= k {
			if dfs(i2, j2, m, n, heights, vis, k) {
				return true
			}
		}
	}

	return false
}

func inside(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
