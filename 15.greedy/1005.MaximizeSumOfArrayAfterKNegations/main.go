package main

// https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/

import "sort"

// problem: largest sum after k negations
// idea:
// 1. sort array to handle negative numbers
// 2. Flip the most negative numbers if having
// 3. after that we have all positive numbers in the entire array
// 4. if k is remaining, then we check if k is even
// - true: we will return sum of the array
// - false: flip the smallest number to minize loss

// Greedy:
// - Flip the most negative numbers first to maximize gain
// - If k is odd after that, flip the smallest number to minimize loss

// TC: O(nlogn), SC: O(1)
func largestSumAfterKNegations(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	sum := 0
	minNum := 10_000
	for i := 0; i < n; i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		minNum = min(minNum, nums[i])
		sum += nums[i]
	}

	if k%2 == 0 {
		return sum
	}

	return sum - 2*minNum
}
