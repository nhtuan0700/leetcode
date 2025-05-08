package main

// https://leetcode.com/problems/search-in-rotated-sorted-array/

// TC: O(logn)
func search1(nums []int, target int) int {
	l := 0
	h := len(nums) - 1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] == target {
			return m
		}

		if nums[m] < target { // right
			if nums[l] > nums[m] && nums[l] <= target {
				h = m - 1
				continue
			}

			l = m + 1
		} else { // left
			if nums[h] < nums[m] && nums[h] >= target {
				l = m + 1
				continue
			}

			h = m - 1
		}
	}

	return -1
}

// TC: O(logn)
func search2(nums []int, target int) int {
	l := 0
	h := len(nums) - 1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] { // left side sorted
			if nums[l] <= target && target < nums[m] {
				h = m - 1
			} else {
				l = m + 1
			}
		} else { // right side sorted
			if nums[m] < target && target <= nums[h] {
				l = m + 1
			} else {
				h = m - 1
			}
		}
	}

	return -1
}
