package main

// https://leetcode.com/problems/battleships-in-a-board/

// up, right, down, left
var dx = [4]int{-1, 0, 1, 0}
var dy = [4]int{0, -1, 0, 1}

// TC: O(m*n) SC: O(m*n)
func countBattleships(board [][]byte) int {
	m := len(board)
	n := len(board[0])

	vis := make([][]bool, m)
	for i := range m {
		vis[i] = make([]bool, n)
	}

	count := 0
	for i := range m {
		for j := range n {
			if vis[i][j] || board[i][j] == '.' {
				continue
			}

			dfs(i, j, m, n, board, vis)
			count++
		}
	}

	return count
}

func dfs(x, y, m, n int, board [][]byte, vis [][]bool) {
	vis[x][y] = true

	for k := range 4 {
		u := x + dx[k]
		v := y + dy[k]

		if !inside(u, v, m, n) || vis[u][v] || board[u][v] == '.' {
			continue
		}

		dfs(u, v, m, n, board, vis)
	}
}

func inside(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
