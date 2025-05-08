package main

// https://leetcode.com/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray1(arr []int) int {
	l := 0
	h := len(arr) - 1

	for l <= h {
		m := l + (h-l)/2
		// if m == len(arr)-1 {
		// 	return m
		// }
		if arr[m] < arr[m+1] {  // up hill
			l = m + 1
		} else { // down hill
			h = m - 1
		}
	}

	return l
}

func peakIndexInMountainArray2(arr []int) int {
	l := 0
	h := len(arr) - 1

	for l < h {
		m := l + (h-l)/2
		if arr[m] < arr[m+1] { // up hill
			l = m + 1
		} else {  // down hill
			h = m
		}
	}

	// return l
	return h
}

// 3 <= len(arr) <= 10^5
func peakIndexInMountainArray3(arr []int) int {
	l := 1
	h := len(arr) - 2

	for l <= h {
		m := l + (h-l)/2
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return m
		}
		if arr[m] < arr[m+1] { // up hill
			l = m + 1
		} else { // down hill
			h = m - 1
		}
	}

	return -1
}
