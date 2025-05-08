package main

// TC: O(n+E), SC: O(n)
// E = len(edges)
func validPath(n int, edges [][]int, source int, destination int) bool {
	if len(edges) == 0 {
		return true
	}

	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	vis := make([]bool, n)
	dfs(adj, vis, source)

	return vis[destination]
}

func dfs(adj [][]int, vis []bool, cur int) {
	vis[cur] = true

	for _, v := range adj[cur] {
		if vis[v] {
			continue
		}

		dfs(adj, vis, v)
	}
}
