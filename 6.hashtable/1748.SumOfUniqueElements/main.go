package main

// https://leetcode.com/problems/sum-of-unique-elements/description/

// hashmap
// TC: O(n), SC: O(n)
func sumOfUnique1(nums []int) int {
	freqs := map[int]int{}

	result := 0
	for _, num := range nums {
		freqs[num]++
		if freqs[num] == 1 {
			result += num
		} else if freqs[num] == 2 {
			result -= num
		}
	}

	return result
}

// counting array
// TC: O(n), SC: O(max(arr))
func sumOfUnique2(nums []int) int {
	counts := make([]int, 101)

	result := 0
	for _, num := range nums {
		counts[num]++
		if counts[num] == 1 {
			result += num

		} else if counts[num] == 2 {
			result -= num
		}
	}

	return result
}
