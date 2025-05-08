package main

import "sort"

// https://leetcode.com/problems/merge-intervals/

// TC: O(nlogn), SC: O(n)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		lastInterval := result[len(result)-1]
		currInterval := intervals[i]
		// overlap: B.L <= A.R -> merge
		if currInterval[0] <= lastInterval[1] {
			lastInterval[1] = max(currInterval[1], lastInterval[1])
		} else { // A.R < B.L
			result = append(result, currInterval)
		}
	}

	return result
}
