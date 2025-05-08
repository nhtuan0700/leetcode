package main

// https://leetcode.com/problems/valid-mountain-array/description/
// TC: O(n)
func validMountainArray1(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	peakTrend := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
		if arr[i] > arr[i-1] { // up hill
			if peakTrend < 0 {
				return false
			}
			peakTrend = 1
		} else { // down hill
			if peakTrend == 0 {
				return false
			}
			peakTrend = -1
		}
	}

	return peakTrend == -1
}

func validMountainArray2(arr []int) bool {
	decreased := false
	increased := false
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
		if arr[i] < arr[i+1] { // up hill
			if decreased {
				return false
			}
			increased = true
		} else { // down hill
			if !increased {
				return false
			}
			decreased = true
		}
	}

	return increased && decreased
}
