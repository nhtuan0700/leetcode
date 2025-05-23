package main

// https://leetcode.com/problems/single-element-in-a-sorted-array/description/

// mid := low + (high - low / 2)
// loop low < high
//
//	if nums[mid] != nums[mid - 1] && nums[mid] !== nums[mid+1] => return nums[mid]
//	if nums[mid] == nums[mid-1] (leftside)
//	    if low -> mid have even number: low = mid + 1
//	    else low -> mid have odd number: high = mid
//	if nums[mid] == nums[mid-1]  (rightside)
//	    if mid -> high have even number: high = mid - 1
//	    else mid -> high have odd number: low = mid
//
// the result is always on the left/right side of middle posisiton with low -> high having odd number
func singleNonDuplicate1(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		equalLeft := mid > 0 && nums[mid] == nums[mid-1]
		equalRight := mid < len(nums)-1 && nums[mid] == nums[mid+1]
		if !equalLeft && !equalRight {
			return nums[mid]
		}
		if equalLeft {
			// the low -> mid having even number
			if (mid-low+1)%2 == 0 {
				low = mid + 1
			} else { // odd number
				high = mid
			}
		} else {
			// the mid -> high having even number
			if (high-mid+1)%2 == 0 {
				high = mid - 1
			} else { // odd
				low = mid
			}
		}
	}

	return -1
}

func singleNonDuplicate2(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}
		if nums[mid] == nums[mid-1] {
			// the low -> mid having even number
			if (mid-low+1)%2 == 0 {
				low = mid + 1
			} else { // odd number
				high = mid
			}
		} else {
			// the mid -> high having even number
			if (high-mid+1)%2 == 0 {
				high = mid - 1
			} else { // odd
				low = mid
			}
		}
	}

	// return nums[high]
	return nums[low]
}

// not need
func singleNonDuplicate3(nums []int) int {
	low := 0
	high := len(nums) - 1
	mid := (high - low) / 2
	for low < high {
		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		if nums[mid] == nums[mid-1] { // equal left
			// the low -> mid having even number
			if (mid-low+1)%2 == 0 {
				low = mid + 1
			} else { // odd number
				high = mid
			}
		} else { // equal right
			// the mid -> high having even number
			if (high-mid+1)%2 == 0 {
				high = mid - 1
			} else { // odd
				low = mid
			}
		}

		mid = low + (high-low)/2
	}

	return nums[low]
}
