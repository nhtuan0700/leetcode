package main

// https://leetcode.com/problems/01-matrix/
// 1. process bfs for cells 1
//    - if mat[x][y] == 0 => dis[i][j] = abs(x - i) + abs(y - j)
// TC: O(m^2*n^2) => bad

// 2. BFS for cells 0
// up, right, down, left
var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

const inf = 1e9

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	q := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, [2]int{i, j})
			} else {
				mat[i][j] = inf
			}
		}
	}

	head := 0
	for head < len(q) {
		cur := q[head]
		head++ // pop

		for _, d := range dirs {
			i2 := cur[0] + d[0]
			j2 := cur[1] + d[1]

			if inside(i2, j2, m, n) && mat[i2][j2] == inf {
				mat[i2][j2] = mat[cur[0]][cur[1]] + 1
				q = append(q, [2]int{i2, j2})
			}
		}
	}

	return mat
}

func inside(x, y, n, m int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
