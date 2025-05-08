package main

import "sort"

// https://leetcode.com/problems/relative-sort-array/
// Using counting sort with arr1
// then iterate arr2 to sort with countArr
// Finally append the remain elements of countArr
// TC: O(max(arr)+n+n2), SC: O(max(arr)+n)
func relativeSortArray1(arr1 []int, arr2 []int) []int {
	max := 0

	// TC: O(n)
	for _, num := range arr1 {
		if max < num {
			max = num
		}
	}

	countArr := make([]int, max+1)

	for _, num := range arr1 {
		countArr[num]++
	}

	// TC: O(n+n2)
	result := make([]int, 0)
	for _, num := range arr2 {
		for countArr[num] > 0 {
			result = append(result, num)
			countArr[num]--
		}
	}

	// TC: O(max(arr)+n)
	for i := 0; i <= max; i++ {
		for countArr[i] > 0 {
			result = append(result, i)
			countArr[i]--
		}
	}
	return result
}

// max(arr)=1002
// TC: O(n+n2+max(arr)), SC: O(max(arr))
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	countArr := make([]int, 1002)

	for _, num := range arr1 {
		countArr[num]++
	}

	idx := 0
	for _, num := range arr2 {
		for countArr[num] > 0 {
			arr1[idx] = num
			countArr[num]--
			idx++
		}
	}

	for i := range countArr {
		for countArr[i] > 0 {
			arr1[idx] = i
			countArr[i]--
			idx++
		}
	}

	return arr1
}

// if max in arr is big
// using lambda sort by order of arr2
// arr1: sort by index order of arr2 if exist, else sort by value of arr1
// TC: O(n2+nlogn) SC: O(n2)
func relativeSortArray3(arr1 []int, arr2 []int) []int {
	pos := make(map[int]int)

	for i, num := range arr2 {
		pos[num] = i
	}

	sort.SliceStable(arr1, func(i, j int) bool {
		var px, py int
		if idx, ok := pos[arr1[i]]; ok {
			px = idx
		} else {
			px = 1e9
		}

		if idx, ok := pos[arr1[j]]; ok {
			py = idx
		} else {
			py = 1e9
		}

		if px != py {
			return px < py
		}

		return arr1[i] < arr1[j]
	})

	return arr1
}
