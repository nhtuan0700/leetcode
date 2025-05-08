package main

// https://leetcode.com/problems/arranging-coins/

// sum = n * (n + 1) / 2

func arrangeCoins(n int) int {
	low := 1
	high := n

	for low <= high {
		mid := low + (high-low)/2
		sum := mid * (mid + 1) / 2
		if sum == n {
			return mid
		}
		if sum > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return high
}

func arrangeCoins2(n int) int {
	low := 1
	high := n

	for low <= high {
		mid := low + (high-low)/2
		sum := mid * (mid + 1) / 2
		if sum == n {
			return mid
		}
		if sum > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low - 1
}
