package main

// When we first visit a cell with value 1, we:
// - Find the bottom-right coordinates of the rectangular farmland.
// - Mark the entire rectangle area as 0 to avoid revisiting.

// TC: O(m*n), SC: O(1)
func findFarmland2(land [][]int) [][]int {
	m := len(land)
	n := len(land[0])

	res := make([][]int, 0)
	for i := range m {
		for j := range n {
			if land[i][j] == 0 {
				continue
			}

			endRow, endCol := i, j

			for endRow < m-1 {
				endRow++
				if land[endRow][j] == 0 {
					endRow--
					break
				}
			}

			for endCol < n-1 {
				endCol++
				if land[i][endCol] == 0 {
					endCol--
					break
				}
			}

			for i2 := i; i2 <= endRow; i2++ {
				for j2 := j; j2 <= endCol; j2++ {
					land[i2][j2] = 0
				}
			}

			res = append(res, []int{i, j, endRow, endCol})
		}
	}

	return res
}
