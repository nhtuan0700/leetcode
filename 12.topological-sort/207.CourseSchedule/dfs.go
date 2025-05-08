package main

// https://leetcode.com/problems/course-schedule/

// return true if no cycle
// TC: O(V+E), SC: O(V+E)
func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)

	for _, e := range prerequisites {
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// res := make([]int, 0)
	vis := make([]bool, numCourses)
	path := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if vis[i] {
			continue
		}

		if !dfs(i, adj, vis, path) {
			return false
		}
	}

	return true
}

func dfs(cur int, adj [][]int, vis, path []bool) bool {
	vis[cur] = true
	path[cur] = true

	for _, v := range adj[cur] {
		if !vis[v] {
			if !dfs(v, adj, vis, path) {
				return false
			}
		} else if path[v] {
			return false
		}
	}

	path[cur] = false

	return true
}
