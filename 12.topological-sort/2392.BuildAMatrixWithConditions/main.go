package main

// https://leetcode.com/problems/build-a-matrix-with-conditions/

// row: 3 1 2 -> [1, 2, 0]
// col: 3 2 1 -> [2, 1, 0]

// TC: O(V+E), SC: O(V+E)
// V: k
// E: len(rowConditions) + len(colConditions)
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowAdj := make([][]int, k)
	for _, con := range rowConditions {
		u := con[0] - 1
		v := con[1] - 1
		rowAdj[u] = append(rowAdj[u], v)
	}
	colAdj := make([][]int, k)
	for _, con := range colConditions {
		u := con[0] - 1
		v := con[1] - 1
		colAdj[u] = append(colAdj[u], v)
	}

	rows := make([]int, k)
	idx := k - 1
	vis := make([]bool, k)
	path := make([]bool, k)
	for i := 0; i < k; i++ {
		if !vis[i] {
			if !dfs(i, rowAdj, vis, path, &rows, &idx) {
				return nil
			}
		}
	}
	cols := make([]int, k)
	idx = k - 1
	vis = make([]bool, k)
	path = make([]bool, k)
	for i := 0; i < k; i++ {
		if !vis[i] {
			if !dfs(i, colAdj, vis, path, &cols, &idx) {
				return nil
			}
		}
	}

	res := make([][]int, k)
	for i := 0; i < k; i++ {
		res[i] = make([]int, k)
	}
	for i := 0; i < k; i++ {
		row := rows[i]
		col := cols[i]
		res[row][col] = i + 1
	}

	return res
}

func dfs(cur int, adj [][]int, vis, path []bool, res *[]int, idx *int) bool {
	vis[cur] = true
	path[cur] = true

	for _, v := range adj[cur] {
		if !vis[v] {
			if !dfs(v, adj, vis, path, res, idx) {
				return false
			}
		} else if path[v] {
			return false
		}
	}

	path[cur] = false

	(*res)[cur] = *idx
	*idx -= 1
	return true
}
