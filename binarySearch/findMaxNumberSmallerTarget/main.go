package main

import "fmt"

// max number < target
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2

		// if smaller than target, ignore left
		if arr[mid] < target {
			low = mid + 1
		} else { // >= ignore right
			high = mid - 1
		}
	}

	// return low - 1
	return high
}

func binarySearch2(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	pos := -1
	for low <= high {
		mid := low + (high-low)/2

		// if smaller than target, ignore left
		if arr[mid] < target {
			low = mid + 1
			pos = mid
		} else { // >= ignore right
			high = mid - 1
		}
	}

	return pos
}

func main() {
	arr := []int{0, 1, 2, 3, 4}

	fmt.Println(binarySearch(arr, 3))  // 2
	fmt.Println(binarySearch2(arr, 3)) // 2
}
