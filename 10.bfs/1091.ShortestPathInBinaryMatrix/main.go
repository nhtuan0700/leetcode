package main

// https://leetcode.com/problems/shortest-path-in-binary-matrix/

// (i-1, j-1) (i-1, j) (i-1, j+1)
// (i, j-1)            (i, j+1)
// (i+1, j-1) (i+1, j) (i+1, j+1)

const inf = 1e9

var dirs = [8][2]int{
	[2]int{-1, -1},
	[2]int{-1, 0},
	[2]int{-1, 1},
	[2]int{0, -1},
	[2]int{0, 1},
	[2]int{1, -1},
	[2]int{1, 0},
	[2]int{1, 1},
}

// TC: O(n^2), SC: O(n^2)
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	q := [][2]int{}
	q = append(q, [2]int{0, 0})
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = inf
		}
	}
	dis[0][0] = 0

	i := 0
	for i < len(q) {
		val := q[i]
		i++ // pop

		for _, d := range dirs {
			i2 := val[0] + d[0] // x
			j2 := val[1] + d[1] // y

			if inside(i2, j2, n) && grid[i2][j2] == 0 && dis[i2][j2] == inf {
				dis[i2][j2] = dis[val[0]][val[1]] + 1
				q = append(q, [2]int{i2, j2})
			}
		}
	}
	if dis[n-1][n-1] >= inf {
		return -1
	}
	return dis[n-1][n-1] + 1
}

func inside(x, y, n int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}
