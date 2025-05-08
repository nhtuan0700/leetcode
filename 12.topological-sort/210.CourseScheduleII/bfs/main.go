package main

// prerequisites[i] = (0,1) => 1 -> 0

// https://leetcode.com/problems/course-schedule-ii/description/

// TC: O(numCourses + len(prerequistes)) SC: O(numCourses)
func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, e := range prerequisites {
		adj[e[1]] = append(adj[e[1]], e[0])
		indegree[e[0]]++
	}

	q := make([]int, 0)
	for i, d := range indegree {
		if d == 0 {
			q = append(q, i)
		}
	}

	head := 0
	res := make([]int, numCourses)
	idx := 0
	for head < len(q) {
		cur := q[head]
		head++
		res[idx] = cur
		idx++
		for _, v := range adj[cur] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if idx < numCourses {
		return nil
	}

	return res
}
