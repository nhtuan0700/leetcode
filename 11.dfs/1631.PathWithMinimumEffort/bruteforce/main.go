package main

// https://leetcode.com/problems/path-with-minimum-effort/
// find all posible from first cell to the last cell, keep track the maximum absolute differences
// Then we find the min value

var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}

// O(3^(m+n))
// at each node: we have m+n option, and we have 3 directions
func minimumEffortPath(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])

	// vis := make([][]bool, m)
	// for i := range m {
	//     vis[i] = make([]bool, n)
	// }

	path := make([][]bool, m)
	for i := range m {
		path[i] = make([]bool, n)
	}

	vals := dfs(0, 0, m, n, heights, path, 0)
	res := int(1e9)
	for _, v := range vals {
		res = min(res, v)
	}

	if res == 1e9 {
		return 0
	}
	return res
}

func dfs(i, j, m, n int, heights [][]int, path [][]bool, v int) []int {
	path[i][j] = true
	res := make([]int, 0)
	for k := range 4 {
		i2 := i + dx[k]
		j2 := j + dy[k]

		if !inside(i2, j2, m, n) || path[i2][j2] {
			continue
		}
		prev := v
		v = max(v, abs(heights[i][j]-heights[i2][j2]))
		if i2 == m-1 && j2 == n-1 {
			res = append(res, v)
		} else {
			res = append(res, dfs(i2, j2, m, n, heights, path, v)...)
		}
		v = prev
	}

	path[i][j] = false
	return res
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
