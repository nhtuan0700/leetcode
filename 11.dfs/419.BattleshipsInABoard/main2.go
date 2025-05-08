package main

// https://leetcode.com/problems/battleships-in-a-board/

// because battleships can only be placed horizontally or vertically
// => we need only count the first cell that is the most top-left cell of battleship
func countBattleships2(board [][]byte) int {
	count := 0
	m := len(board)
	n := len(board[0])

	for i := range m {
		for j := range n {
			if board[i][j] == '.' {
				continue
			}
			if isFirstCell(i, j, board) {
				count++
			}
		}
	}

	return count
}

// most top-left cell
func isFirstCell(i, j int, board [][]byte) bool {
	// vertical
	if i > 0 && board[i-1][j] == 'X' {
		return false
	}

	// vertical
	if j > 0 && board[i][j-1] == 'X' {
		return false
	}
	return true
}
