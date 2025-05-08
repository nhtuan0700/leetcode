package main

// https://leetcode.com/problems/find-eventual-safe-states/description/

// safe node: have no cycle paths

// TC: O(n+E), S: O(n)
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	vis := make([]bool, n)
	path := make([]bool, n)

	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		dfs(i, graph, vis, path)
	}

	res := make([]int, 0)
	for i, p := range path {
		if !p {
			res = append(res, i)
		}
	}
	return res
}

func dfs(cur int, graph [][]int, vis, path []bool) bool {
	vis[cur] = true
	path[cur] = true

	for _, v := range graph[cur] {
		if !vis[v] {
			if !dfs(v, graph, vis, path) {
				// detect cyle
				return false
			}
		} else if path[v] {
			return false
		}
	}

	path[cur] = false

	return true
}
