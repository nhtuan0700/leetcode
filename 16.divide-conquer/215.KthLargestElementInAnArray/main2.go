package main

import "math/rand"

func findKthLargest2(nums []int, k int) int {
	return quickSelect(0, len(nums)-1, k, nums)
}

func quickSelect(l, r, k int, nums []int) int {
	if l == r {
		return nums[l]
	}
	i, j := l, r
	pivot := nums[rand.Intn(r-l+1)+l]
	for i < j {
		for nums[i] > pivot {
			i++
		}
		for nums[j] < pivot {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	// j <= i, j có thể là phần tử cuối cùng >= pivot hoặc là cái đầu tiên mà nó < pivot
	// [5 4 6 3 1 2], pivot = 6
	if nums[j] < pivot {
		j--
	}

	// l -> j: >= pivot
	// j+1 -> r: <= pivot
	count := j - l + 1
	if k <= count {
		return quickSelect(l, j, k, nums)
	} else {
		return quickSelect(j+1, r, k - count, nums)
	}

	// if k-1 <= j {
	// 	return quickSelect(l, j, k, nums)
	// } else {
	// 	return quickSelect(j+1, r, k, nums)
	// }
}
