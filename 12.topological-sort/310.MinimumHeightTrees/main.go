package main

// https://leetcode.com/problems/minimum-height-trees/

// Idea: because most of tree, there are either 1 or 2 central nodes that can be the roots of MHTs
// then we will triming the leaves node until 1 or 2 nodes remain

// TC: O(V+E) ~ O(n), SC: O(V+E) ~ O(n)
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([][]int, n)
	inDegree := make([]int, n)

	for _, e := range edges {
		u := e[0]
		v := e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		inDegree[u]++
		inDegree[v]++
	}

	q := make([]int, 0)
	// add leaf node first
	for i := 0; i < n; i++ {
		if inDegree[i] == 1 {
			q = append(q, i)
		}
	}

	head := 0
	// 1, 2 last leaf nodes after trimming leaf nodes are the result
	for n > 2 {
		size := len(q) - head
		n -= size

		for size > 0 {
			cur := q[head]
			head++ // pop
			size--
			for _, v := range adj[cur] {
				inDegree[v]--
				if inDegree[v] == 1 {
					q = append(q, v)
				}
			}
		}
	}

	res := make([]int, 0)
	for head < len(q) {
		res = append(res, q[head])
		head++
	}

	return res
}
