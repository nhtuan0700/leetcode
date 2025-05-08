package main

import "sort"

// https://leetcode.com/problems/height-checker/
// TC: O(nlogn), SC: O(n)
func heightChecker1(heights []int) int {
	tempHeights := make([]int, len(heights))
	copy(tempHeights, heights)
	sort.Ints(tempHeights)

	result := 0
	for i := range heights {
		if heights[i] != tempHeights[i] {
			result++
		}
	}

	return result
}

func heightChecker2(heights []int) int {
	freq := make([]int, 101)
	for _, h := range heights {
		freq[h]++
	}

	diff := 0
	// current index of heights
	i := 0
	for h, f := range freq {
		for f > 0 {
			// expected[i] = h
			if heights[i] != h {
				diff++
			}
			i++
			f--
		}
	}

	return diff
}
