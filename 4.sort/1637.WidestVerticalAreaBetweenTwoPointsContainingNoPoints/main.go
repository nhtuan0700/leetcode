package main

import "sort"

// https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/description/
// TC: O(nlogn)
func maxWidthOfVerticalArea(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	maxWidth := 0
	for i := 1; i < len(points); i++ {
		maxWidth = max(maxWidth, points[i][0]-points[i-1][0])
	}

	return maxWidth
}
