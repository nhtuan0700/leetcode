package main

// https://leetcode.com/problems/two-sum/

// [2, 7,11, 15]
// i: 0 -> n
// j: 0 -> n - 1

// TC: O(n), SC: O(n)
func twoSum1(nums []int, target int) []int {
	cached := map[int]int{}

	for i, num := range nums {
		remainNum := target - num
		if _, ok := cached[remainNum]; ok {
			return []int{cached[remainNum], i}
		}

		cached[num] = i
	}

	return nil
}

// TC: O(n^2), SC: O(1)
func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
