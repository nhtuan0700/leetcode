package main

// https://leetcode.com/problems/number-of-operations-to-make-network-connected/description/

// TC: O(n), SC: O(n^2)
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 { // edge case
		return -1
	}
	adj := make([][]int, n)

	for _, e := range connections {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	vis := make([]bool, n)
	path := make([]bool, n)
	unConnected := 0
	cycleCount := 0
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}

		unConnected++
		cycleCount += dfs(i, -1, adj, vis, path)
	}

	if cycleCount >= unConnected-1 {
		return unConnected - 1
	}

	return -1
}

// count cycle
func dfs(cur, prev int, adj [][]int, vis, path []bool) int {
	cycleCount := 0
	vis[cur] = true
	path[cur] = true

	for _, v := range adj[cur] {
		if !vis[v] {
			cycleCount += dfs(v, cur, adj, vis, path)
		} else if v != prev && path[v] {
			cycleCount++
		}
	}

	path[cur] = false

	return cycleCount
}
