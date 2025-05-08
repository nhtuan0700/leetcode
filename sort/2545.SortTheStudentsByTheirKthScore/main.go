package main

import "sort"

// https://leetcode.com/problems/sort-the-students-by-their-kth-score/

func sortTheStudents(score [][]int, k int) [][]int {
	sort.SliceStable(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})

	return score
}
