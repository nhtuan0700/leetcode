package main

import "math/rand"

// use quick select
// TC: average O(n), worst case: O(n^2)
// SC: O(1)
func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	k -= 1 // convert to 0-indexed
	for l < r {
		// fmt.Println(nums)
		pivotIndex := partition(nums, l, r)
		if pivotIndex == k {
			return nums[pivotIndex]
		}

		if pivotIndex > k {
			r = pivotIndex - 1
		} else {
			l = pivotIndex + 1
		}
	}

	return -1
}

func partition(nums []int, l, r int) int {
	p := rand.Intn(r-l+1) + l
	nums[p], nums[r] = nums[r], nums[p]
	pivot := nums[r] // select pivot

	i := l

	for j := i; j < r; j++ {
		if nums[j] > pivot {
			nums[i], nums[j] = nums[j], nums[i] // swap
			i++
		}
	}

	nums[i], nums[r] = nums[r], nums[i] // swap

	return i
}
