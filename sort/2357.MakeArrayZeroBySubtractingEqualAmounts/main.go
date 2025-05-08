package main

import "sort"

// https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
// TC: O(nlogn)

func minimumOperations1(nums []int) int {
	sort.Ints(nums)

	i := 0
	x := 0
	for _, num := range nums {
		num -= x
		if num > 0 {
			i++
		}
		x += num
	}

	return i
}

func minimumOperations2(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		if num > 0 {
			set[num] = true
		}
	}

	return len(set)
}
