package main

// prerequisites[i] = (0,1) => 1 -> 0

// TC: O(V + E), SC: O(V)
// V: numCourses
// E: len(prerequisites)
func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, e := range prerequisites {
		adj[e[1]] = append(adj[e[1]], e[0])
		indegree[e[0]]++
	}

	res := make([]int, 0)
	vis := make([]bool, numCourses)
	path := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if vis[i] {
			continue
		}

		if !dfs(i, adj, vis, path, &res) {
			return nil
		}
	}

	l := 0
	h := numCourses - 1
	for l <= h {
		res[l], res[h] = res[h], res[l]
		l++
		h--
	}

	return res
}

func dfs(cur int, adj [][]int, vis, path []bool, res *[]int) bool {
	vis[cur] = true
	path[cur] = true

	for _, v := range adj[cur] {
		if !vis[v] {
			// detect cycle
			if !dfs(v, adj, vis, path, res) {
				return false
			}
		} else if path[v] { // detect cycle
			return false
		}
	}

	*res = append(*res, cur)
	path[cur] = false

	return true
}
