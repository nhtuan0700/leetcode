package main

// https://leetcode.com/problems/course-schedule/

// [0,1] => 1->0
// return true if no cycle
// TC: O(V+E), SC: O(V+E)
func canFinish2(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, e := range prerequisites {
		adj[e[1]] = append(adj[e[1]], e[0])
		inDegree[e[0]]++
	}

	q := make([]int, 0)
	count := 0
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			count++
			q = append(q, i)
		}
	}

	head := 0
	for head < len(q) {
		cur := q[head]
		head++

		for _, v := range adj[cur] {
			inDegree[v]--

			if inDegree[v] == 0 {
				count++
				q = append(q, v)
			}
		}
	}

	return count == numCourses
}
