package main

import "sort"

// https://leetcode.com/problems/insert-interval/

// Can do bs

// 3 steps:
// 1. Add intervals that is before newInterval
// 2. Merge newInterval with all overlapping intervals
// 3. Add the rest of elements that comes after newInterval
// TC: O(nlogn), SC: O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < len(intervals) && newInterval[1] >= intervals[i][0] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}

// O(n)
func insert2(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)

	inserted := false
	for _, in := range intervals {
		// left
		if in[1] < newInterval[0] {
			res = append(res, in)
			continue
		} else if in[0] <= newInterval[1] && in[1] >= newInterval[0] {
			newInterval[0] = min(newInterval[0], in[0])
			newInterval[1] = max(newInterval[1], in[1])
		} else { // right
			if !inserted {
				res = append(res, newInterval)
				inserted = true
			}
			res = append(res, in)
		}
	}

	if !inserted {
		res = append(res, newInterval)
	}

	return res
}

// TC: O(nlogn)
func insert3(intervals [][]int, newInterval []int) [][]int {
	copyIntervals := intervals
	copyIntervals = append(copyIntervals, newInterval)
	sort.Slice(copyIntervals, func(i, j int) bool { return copyIntervals[i][0] < copyIntervals[j][0] })

	res := [][]int{copyIntervals[0]}
	for _, in := range copyIntervals {
		if res[len(res)-1][1] >= in[0] {
			res[len(res)-1][0] = min(res[len(res)-1][0], in[0])
			res[len(res)-1][1] = max(res[len(res)-1][1], in[1])
		} else {
			res = append(res, in)
		}
	}
	return res
}
