package main

import "fmt"

func binarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high - low) /2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	
	return false
}

func main() {
	arr := []int{0, 1, 2, 3, 4}

	fmt.Println(binarySearch(arr, 5))
}
