package main

import "fmt"

// https://leetcode.com/problems/remove-element/

// [3,2,2,3] val = 3
// [2,2,_,_]

// Use 2 pointer (i, k).
// i: iteration position
// k: We will to move the element that is not equal to val, and return the number of the moved time

func removeElement(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
