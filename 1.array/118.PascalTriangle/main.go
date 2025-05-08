package main

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		// 0 - 1, 1 - 2, 2 - 3
		// i: currentRow, i+1: length of current row
		// first and last element is always 1
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}

func generate2(numRows int) [][]int {
	var result [][]int

	for i := 1; i <= numRows; i++ {
		newRow := make([]int, i)
		newRow[0] = 1
		newRow[i - 1] = 1 
		for j := 1; j < i - 1; j++ {
			newRow[j] = result[len(result) - 1][j-1] + result[len(result) - 1][j]
		}

		result = append(result, newRow)
	}

	return result
}

func main() {
	fmt.Println(generate2(5))
}
