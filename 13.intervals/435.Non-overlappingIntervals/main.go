package main

// https://leetcode.com/problems/non-overlapping-intervals/description/

import "sort"

// Sort the intervals by their end time. Then, iterate through each interval.
// - If the current interval overlaps with the previous one (i.e. start < last end), remove it.
// - This ensures we always keep the interval with the earliest end time, minimizing overlap.
// TC: O(nlogn), SC: O(1)
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	count := 0
	last := int(-1e6)
	for _, in := range intervals {
		// not overlap
		if last <= in[0] {
			last = in[1]
		} else {
			// last = min(last, in[1])
			count++
		}
	}

	return count
}
