package main

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// O(logn)
func searchRange1(nums []int, target int) []int {
	midPos := -1
	low := 0
	high := len(nums) - 1

	// O(logn)
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			midPos = mid
			break
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if midPos == -1 {
		return []int{-1, -1}
	}

	start := midPos
	end := midPos
	// (low, high)
	// left to find start index
	lLow := low
	lHigh := midPos
	for lLow <= lHigh {
		mid := lLow + (lHigh-lLow)/2
		if nums[mid] == target {
			start = mid
			lHigh = mid - 1
		} else { // nums[mid] < target
			lLow = mid + 1
		}
	}

	// right to find end index
	rLow := midPos
	rHigh := high
	for rLow <= rHigh {
		mid := rLow + (rHigh-rLow)/2
		if nums[mid] == target {
			end = mid
			rLow = mid + 1
		} else { // nums[mid] > target
			rHigh = mid - 1
		}
	}

	return []int{start, end}
}

// recursion combined to binary search
// nums [5,7,7,8,8,10], target = 8
// recursion binarySearch to find the mid where nums[mid] == target
// then we have 2 pointer start, end = mid
// decrease start until nums[start] < target
// increase end until nums[start] > target
func searchRange2(nums []int, target int) []int {
	return bsearch1(nums, target, 0, len(nums)-1)
}
// TC: O(logn), worstcase: O(n)
func bsearch1(nums []int, target, l, r int) []int {
	if l > r {
		return []int{-1, -1}
	}

	m := l + (r-l)/2
	if nums[m] == target {
		start := m
		end := m
		for nums[start] == target {
			start--
			if start < 0 {
				break
			}
		}

		for nums[end] == target {
			end++
			if end == len(nums) {
				break
			}
		}

		return []int{start + 1, end - 1}
	}

	if nums[m] > target {
		return bsearch1(nums, target, l, m-1)
	}
	return bsearch1(nums, target, m+1, r)
}

func bsearch2(nums []int, target, l, r int) []int {
	if l > r {
		return []int{-1, -1}
	}

	m := l + (r-l)/2
	if nums[m] == target {
		start := m
		end := m
		// (l [start end] h)
		// (l -> start)
		// (end -> h)
		ll := l
		lh := m
		for ll <= lh {
			lm := ll + (lh-ll)/2
			if nums[lm] == target {
				start = lm
				lh = lm - 1
			} else { // nums[mid] < target
				ll = lm + 1
			}
		}

		rl := m
		rr := r
		for rl <= rr {
			rm := rl + (rr-rl)/2
			if nums[rm] == target {
				end = rm
				rl = rm + 1
			} else { // nums[mid] > target
				rr = rm - 1
			}
		}

		return []int{start, end}
	}

	if nums[m] > target {
		return bsearch2(nums, target, l, m-1)
	}
	return bsearch2(nums, target, m+1, r)
}
