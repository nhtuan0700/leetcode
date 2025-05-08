package main

import "sort"

// https://leetcode.com/problems/find-right-interval/
// TC: O(nlogn), SC: O(n)
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	sorted := make([][]int, n) // [2] => {start, originalIndex}

	for i, in := range intervals {
		sorted[i] = []int{in[0], i}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	// [1,2] [2,3] [3,4]
	res := make([]int, n)
	for i := 0; i < n; i++ {
		right := intervals[i][1]
		pos := binarySearch(sorted, right)
		if pos == -1 {
			res[i] = -1
		} else {
			res[i] = pos
		}
	}

	return res
}

// find right interval
func binarySearch(intervals [][]int, target int) int {
	l, r := 0, len(intervals)-1
	res := -1
	for l <= r {
		m := l + (r-l)/2
		if intervals[m][0] >= target {
			res = intervals[m][1] // original index
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return res
}
