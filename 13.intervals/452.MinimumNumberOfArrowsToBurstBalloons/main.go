package main

import "sort"

// Greedy idea: always shoot at the earliest possible end to cover as many balloons as possible.
// TC: O(nlogn), SC: O(1)
func findMinArrowShots(points [][]int) int {
	// [1,6],[2,8],[7,15],[10,16],[13,14]
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })

	count := 0
	last := int(-1 << 32)
	for _, p := range points {
		if last < p[0] {
			last = p[1]
			count++
		}
	}

	return count
}
