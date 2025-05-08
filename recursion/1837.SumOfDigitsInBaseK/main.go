package main

// https://leetcode.com/problems/sum-of-digits-in-base-k/

// 34
// 34 / 6 = 5 (4)
// 5 / 6 = 0 (5)

func sumBase(n int, k int) int {
	if n == 0 {
			return 0
	}

	return n%k + sumBase(n / k, k)
}
