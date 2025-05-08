package main

// [1,2,3] [2,5,6]
// [1,2,2,3,5,6]
// len (m+n)
// TC: O(n+m), SC: O(n+m)
func merge1(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	k := 0

	result := make([]int, m+n)
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			result[k] = nums1[i]
			i++
		} else {
			result[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		result[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		result[k] = nums2[j]
		j++
		k++
	}

	copy(nums1, result)
}

// Start from right to avoid overwriting arr1 with len m
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
