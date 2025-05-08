package main

import "sort"

// https://leetcode.com/problems/contains-duplicate/
// TC: O(n), SC: O(n)

func containsDuplicate1(nums []int) bool {
	freqs := map[int]int{}
	for _, num := range nums {
		if freqs[num] == 1 {
			return true
		}

		freqs[num]++
	}

	return false
}

// sort
// TC: O(nlogn+n), SC: O(1)
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
