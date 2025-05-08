package main

// https://leetcode.com/problems/is-graph-bipartite/

// at current node, we assume in a set, then adjacent nodes must be in the remaining set

// TC: O(n + E), SC: O(n)
func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n) // 0: unvisited, -1: set A, 1: set B

	for i := 0; i < n; i++ {
		if colors[i] != 0 {
			continue
		}
		if !dfs(i, graph, colors, -1) {
			return false
		}
	}

	return true
}

func dfs(cur int, graph [][]int, colors []int, color int) bool {
	colors[cur] = color
	nextColor := -color
	for _, v := range graph[cur] {
		if colors[v] == 0 {
			if !dfs(v, graph, colors, nextColor) {
				return false
			}
		} else if color == colors[v] {
			return false
		}
	}

	return true
}
