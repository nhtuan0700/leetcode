package main

import "fmt"

// min number > target
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2

		// if greater than target, ignore right
		if arr[mid] > target {
			high = mid - 1
		} else { // <= ignore left 
			low = mid + 1
		}
	}

	if high+1 == len(arr) {
		return -1
	}

	// return high + 1
	return low
}

func binarySearch2(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	pos := -1
	for low <= high {
		mid := low + (high-low)/2
		// if greater than target, ignore right
		if arr[mid] > target {
			pos = mid
			high = mid - 1
			
		} else { // <= ignore left
			low = mid + 1
		}
	}

	return pos
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(binarySearch(arr, 5))  // -1
	fmt.Println(binarySearch2(arr, 5)) // -1
}
