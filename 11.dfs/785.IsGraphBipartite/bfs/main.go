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
		colors[i] = -1
		if !bfs(i, graph, colors) {
			return false
		}
	}

	return true
}

func bfs(cur int, graph [][]int, colors []int) bool {
	n := len(graph)
	q := make([]int, n)
	q = append(q, cur)
	head := 0
	for head < len(q) {
		val := q[head]
		head++

		for _, v := range graph[val] {
			if colors[v] == 0 {
				colors[v] = -colors[val]
				q = append(q, v)
			} else if colors[v] == colors[val] {
				return false
			}
		}
	}

	return true
}
