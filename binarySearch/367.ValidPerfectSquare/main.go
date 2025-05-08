package main

// https://leetcode.com/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	low := 1
	high := num

	for low <= high {
		mid := low + (high-low)/2
		square := mid * mid
		if square == num {
			return true
		}
		if square > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
