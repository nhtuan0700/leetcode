package main

import "fmt"

// https://leetcode.com/problems/find-pivot-index/description/

// [1,7,3,6,5,6] => 3

// we will calculate the sum of all number
// traverse from left to right, we will calculate sumLeft
// 1. descrease total by substracting num and get the sumRight
// 2. sumRight = total - sumLeft - nums[i]
// sumRight == sumLeft => i

// TC: O(n), SC: O(1)

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	for i, num := range nums {
		total -= num
		if total == leftSum {
			return i
		}
		leftSum += num
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1,7,3,6,5,6}))
}
