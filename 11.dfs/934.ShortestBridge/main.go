package main

// https://leetcode.com/problems/shortest-bridge/

// 1. dfs -> find island 1
// 2. bfs -> find shortest path to island 2

var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, 1, 0, -1}

// TC: O(n^2), SC: O(n^2)
func shortestBridge(grid [][]int) int {
	n := len(grid)
	vis := make([][]bool, n)
	for i := range n {
		vis[i] = make([]bool, n)
	}

	found := false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n && !found; j++ {
			if grid[i][j] == 1 {
				found = true
				dfs(i, j, n, grid, vis)
			}
		}
	}

	q := make([][2]int, 0)
	for i := range n {
		for j := range n {
			if vis[i][j] {
				q = append(q, [2]int{i, j})
			}
		}
	}

	head := 0
	dis := make([][]int, n)
	for i := range n {
		dis[i] = make([]int, n)
	}

	for head < len(q) {
		val := q[head]
		head++
		x, y := val[0], val[1]
		for k := range 4 {
			u := x + dx[k]
			v := y + dy[k]

			if !inside(u, v, n) || vis[u][v] {
				continue
			}

			if grid[u][v] == 1 {
				return dis[x][y]
			}
			dis[u][v] = dis[x][y] + 1
			vis[u][v] = true
			q = append(q, [2]int{u, v})
		}
	}

	return -1
}

func dfs(x, y, n int, grid [][]int, vis [][]bool) {
	vis[x][y] = true

	for k := range 4 {
		u := x + dx[k]
		v := y + dy[k]

		if !inside(u, v, n) || vis[u][v] || grid[u][v] == 0 {
			continue
		}

		dfs(u, v, n, grid, vis)
	}
}

func inside(x, y, n int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}
