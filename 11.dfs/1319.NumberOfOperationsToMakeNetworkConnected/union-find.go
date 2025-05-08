package main

// https://leetcode.com/problems/number-of-operations-to-make-network-connected/description/

// idea: union find
// We know, the minimum number of edges requires for a graph with n nodes to remain connected is n - 1
// Similarly, if there are k componenets in a graph, we need k-1 edges to connect every component

// TC: O(n), SC: O(n^2)
func makeConnected2(n int, connections [][]int) int {
	if len(connections) < n-1 { // edge case
		return -1
	}
	adj := make([][]int, n)

	for _, e := range connections {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	vis := make([]bool, n)
	components := 0
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}

		dfs2(i, adj, vis)
		components++
	}

	return components - 1
}

func dfs2(cur int, adj [][]int, vis []bool) {
	vis[cur] = true

	for _, v := range adj[cur] {
		if vis[v] {
			continue
		}
		dfs2(v, adj, vis)
	}
}
