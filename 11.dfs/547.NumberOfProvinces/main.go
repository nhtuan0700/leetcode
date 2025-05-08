package main

// https://leetcode.com/problems/number-of-provinces/

// input: adjacency matrix
// idea: if the current node is not visited, it denotes that is the new group, then increase result
//   - need to keep track visited node to prevent check node again because it was belonged a group before
//
// TC: O(n^2), SC: O(n)
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	vis := make([]bool, n)
	res := 0
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		res++
		dfs(i, n, isConnected, vis)
	}

	return res
}

func dfs(cur int, n int, isConnected [][]int, vis []bool) {
	vis[cur] = true

	for i := 0; i < n; i++ {
		if isConnected[cur][i] == 0 || vis[i] {
			continue
		}

		dfs(i, n, isConnected, vis)
	}
}
