package main

import "sort"

// https://leetcode.com/problems/find-target-indices-after-sorting-array/

// nums: [1,2,5,2,3], target = 2
// sorted: [1,2,2,3,5]
// output: sort nums, return indices of sorted nums with nums[i] = target
// TC: O(nlogn + n) = O(n(logn+1)) = O(nlogn), SC: O(n)
func targetIndices1(nums []int, target int) []int {
	sort.Ints(nums)
	result := make([]int, 0)
	for i, num := range nums {
		if num > target {
			break
		}

		if num == target {
			result = append(result, i)
		}
	}

	return result
}


// Way2:
// After sorting, 
// result's start index: number of the elements less than pivot
// result's end index: number of the elements equal to pivot + start index
// we dont need to process sorting the array
// TC: O(n), SC: O(n)
func targetIndices2(nums []int, target int) []int {
	smaller := 0
	eq := 0
	result := make([]int, 0)

	for _, num := range nums {
		if num < target {
			smaller++
		}
		if num == target {
			eq++
		}
	}

	for i := smaller; i < smaller+eq; i++ {
		result = append(result, i)
	}

	return result
}
