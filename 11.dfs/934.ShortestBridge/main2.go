package main

// https://leetcode.com/problems/shortest-bridge/

// 1. dfs -> find island 1
// 2. bfs -> find shortest path to island 2

// TC: O(n^2), SC: O(n^2)
func shortestBridge2(grid [][]int) int {
	n := len(grid)
	vis := make([][]bool, n)
	for i := range n {
		vis[i] = make([]bool, n)
	}

	q := make([][2]int, 0)
	head := 0

	found := false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n && !found; j++ {
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
				vis[i][j] = true
				found = true
			}
		}
	}

	for head < len(q) {
		val := q[head]
		head++

		x, y := val[0], val[1]
		for k := range 4 {
			u := x + dx[k]
			v := y + dy[k]

			if !inside(u, v, n) || vis[u][v] || grid[u][v] == 0 {
				continue
			}

			vis[u][v] = true
			q = append(q, [2]int{u, v})
		}
	}

	// we could use newQueue for clear
	// for i := range n {
	// 	for j := range n {
	// 		if vis[i][j] {
	// 			q = append(q, [2]int{i, j})
	// 		}
	// 	}
	// }

	head = 0
	res := 0
	for head < len(q) {
		size := len(q)

		for i := head; i < size; i++ {
			val := q[i]
			x, y := val[0], val[1]
			for k := range 4 {
				u := x + dx[k]
				v := y + dy[k]

				if !inside(u, v, n) || vis[u][v] {
					continue
				}

				if grid[u][v] == 1 {
					return res
				}
				vis[u][v] = true
				q = append(q, [2]int{u, v})
			}
		}
		head = size
		res++
	}

	return -1
}
