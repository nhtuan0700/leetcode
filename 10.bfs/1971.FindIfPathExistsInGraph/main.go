package main

// https://leetcode.com/problems/find-if-path-exists-in-graph/description/
import "container/list"

// TODO: Can resolve by dfs ?

// TC: O(n+len(edges)), SC: O(n)
func validPath(n int, edges [][]int, source int, destination int) bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	vis := make([]bool, n)
	q := list.New()

	q.PushBack(source)
	vis[source] = true

	for q.Front() != nil {
		v := q.Front().Value.(int)
		q.Remove(q.Front()) // pop

		for _, u := range adj[v] {
			if !vis[u] {
				vis[u] = true
				q.PushBack(u)
			}
		}
	}

	return vis[destination]
}
