package main

// https://leetcode.com/problems/detonate-the-maximum-bombs/description/

// Idea: Build graph with connected bombs → Then need to count area each component
// bomb1 can detonate bomb2 if distance d <= radius of bomb1
// where d = sqrt((x2 - x1)^2 + (y2 - y1)^2)
// or (x2 - x1)^2 + (y2 - y1)^2 <= r1^2

// TC: O(n^2), SC: O(n^2)
func maximumDetonation(bombs [][]int) int {
	n := len(bombs)

	adj := make([][]int, n)

	// O(n^2)
	for i := 0; i < n; i++ {
		x1, y1, r := bombs[i][0], bombs[i][1], bombs[i][2]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x2, y2 := bombs[j][0], bombs[j][1]
			dx, dy := x2-x1, y2-y1

			if dx*dx+dy*dy <= r*r {
				adj[i] = append(adj[i], j)
			}
		}
	}

	res := 0
	for i := range n {
		res = max(res, dfs(adj, i, make([]bool, n)))
	}

	return res
}

func dfs(adj [][]int, cur int, vis []bool) int {
	vis[cur] = true
	count := 1
	for _, v := range adj[cur] {
		if vis[v] {
			continue
		}
		count += dfs(adj, v, vis)
	}

	return count
}
