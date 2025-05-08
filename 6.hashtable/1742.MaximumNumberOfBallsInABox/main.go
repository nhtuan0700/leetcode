package main

// https://leetcode.com/problems/maximum-number-of-balls-in-a-box/

// hash table
// TC: O(n), SC: O(n)
func countBalls1(lowLimit int, highLimit int) int {
	counts := map[int]int{}

	result := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		current := i
		for current > 0 {
			sum += current % 10
			current /= 10
		}

		counts[sum]++
		result = max(counts[sum], result)
	}

	return result
}

// counting
// TC: O(n), SC: O(1)
func countBalls2(lowLimit int, highLimit int) int {
	// 10^5 ~ 99999 = 54
	counts := make([]int, 55)

	result := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		current := i
		for current > 0 {
			sum += current % 10
			current /= 10
		}

		counts[sum]++
		result = max(counts[sum], result)
	}

	return result
}
