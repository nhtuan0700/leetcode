package main

// https://leetcode.com/problems/find-lucky-integer-in-an-array/

// hash table
// TC: O(n) SC: O(n)
func findLucky1(arr []int) int {
	freqs := make(map[int]int)

	for _, num := range arr {
		freqs[num]++
	}

	result := -1
	for k, v := range freqs {
		if k == v {
			result = max(result, v)
		}
	}

	return result
}

// counting
// TC: O(n) SC: O(n)
func findLucky2(arr []int) int {
	freqs := make([]int, 501)

	for _, num := range arr {
		freqs[num]++
	}

	result := -1
	for _, num := range arr {
		if freqs[num] == num {
			result = max(result, num)
		}
	}

	return result
}
