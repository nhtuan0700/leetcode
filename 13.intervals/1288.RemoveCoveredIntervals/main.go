package mai

import "sort"

// https://leetcode.com/problems/remove-covered-intervals/

// TC: O(nlogn), SC: O(1)
func removeCoveredIntervals1(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	count := 1
	// [1,4] [2,8] [3,6]

	for i := 1; i < len(intervals); i++ {
		in := intervals[i]
		if cur[1] >= in[1] || cur[0] == in[0] {
			cur[0] = min(cur[0], in[0])
			cur[1] = max(cur[1], in[1])
		} else {
			cur = in
			count++
		}
	}
	return count
}

// TC: O(nlogn), SC: O(1)
func removeCoveredIntervals2(intervals [][]int) int {
	// make sure sooner and wider intervals always before
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1] // wider comes first
		}
		return intervals[i][0] < intervals[j][0] // earlier comes first
	})

	end := -1
	count := 0
	// [1,4] [1,2] [3,4]
	for _, in := range intervals {
		if end < in[1] {
			count++
			end = in[1]
		}
	}

	return count
}
