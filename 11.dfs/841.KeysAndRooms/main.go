package main

// https://leetcode.com/problems/keys-and-rooms/description/

// TC: O(n + E), SC: O(n)
// worse case: E = n^2
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	vis := make([]bool, n)

	dfs(0, rooms, vis)

	for _, v := range vis {
		if !v {
			return false
		}
	}

	return true
}

func dfs(cur int, rooms [][]int, vis []bool) {
	vis[cur] = true

	for _, r := range rooms[cur] {
		if vis[r] {
			continue
		}

		dfs(r, rooms, vis)
	}
}
