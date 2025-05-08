package main

import "fmt"

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

// [4,3,2,7,8,2,3,1] => [5,6]

// we will convert the position of the number value into negative number to mark it was appeared in array

// [4,3,-2,-7,8,2,3,1]
// i = 0; nums[nums[i-1]] = negative of nums[i-1]
// Finally, traverse from i: 1 -> n, nums[i-1] > 0 => push to the result array

// TC: O(n), SC: O(1) exclude variable of result

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		index := num - 1
		if num < 0 {
			index = -num - 1
		}
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	var result []int
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
