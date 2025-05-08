package main

// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/

const (
	nMax = 42
	kMax = 1602
)

var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

const inf = 1e9

// TC: O(m*n*k), SC: O(nMax*nMax*kMax)
func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dis := make([][][]int, m)

	for i := 0; i < m; i++ {
		dis[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = make([]int, k+1)
			for l := 0; l <= k; l++ {
				dis[i][j][l] = inf
			}
		}
	}

	q := [][3]int{}
	q = append(q, [3]int{0, 0, k})
	dis[0][0][k] = 0

	head := 0
	for head < len(q) {
		val := q[head]
		i, j, k := val[0], val[1], val[2]
		head++

		for _, d := range dirs {
			i2 := i + d[0]
			j2 := j + d[1]
			k2 := k

			if !inside(i2, j2, m, n) {
				continue
			}

			if grid[i2][j2] == 1 {
				if k > 0 {
					k2 = k - 1
				} else {
					k2 = -1
				}
			}

			if k2 >= 0 && dis[i2][j2][k2] >= inf {
				dis[i2][j2][k2] = dis[i][j][k] + 1
				q = append(q, [3]int{i2, j2, k2})
			}
		}
	}

	var res int = inf
	for i := 0; i <= k; i++ {
		res = min(res, dis[m-1][n-1][i])
	}

	if res >= inf {
		return -1
	}

	return res
}

func inside(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
