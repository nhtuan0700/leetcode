package main

import "sort"

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/
// sort array
// map[nums[i]] = i if empty
// loops num: nums , result = append(result, map[num])
// TC: O(nlogn), SC: O(3n) = O(n)
func smallerNumbersThanCurrent1(nums []int) []int {
	result := make([]int, len(nums))
	cachedNum := map[int]int{}
	tempNums := make([]int, len(nums))
	copy(tempNums, nums)
	sort.Ints(tempNums)
	for i, num := range tempNums {
		if _, ok := cachedNum[num]; !ok {
			cachedNum[num] = i
		}
	}

	for i, num := range nums {
		result[i] = cachedNum[num]
	}

	return result
}

//  1. tempNums: [][]int{} [[value, index]]
//  2. sort tempNums by value
//  3. result: []int{}
//  4. loop tempNums i -> n, prev = 0:
//     i > 0 && tempNums[i][0] > tempNums[i-1][0] -> result[tempNums[i][1]] = i, prev=i
//     result[tempNums[i][1]] = prev
//
// TC: O(n), SC: O(2n) = O(n)

// [8,1,2,2,3]
// [[8,0],[1,1],[2,2],[2,3],[3,4]]
// [[1,1],[2,2],[2,3],[3,4],[8,0]]
// [4,0,1,1,3]
func smallerNumbersThanCurrent2(nums []int) []int {
	tempArr := make([][]int, len(nums))
	for i, num := range nums {
		subArr := make([]int, 2)
		subArr[0] = num
		subArr[1] = i
		tempArr[i] = subArr
	}

	sort.SliceStable(tempArr, func(i, j int) bool { return tempArr[i][0] < tempArr[j][0] })
	prev := 0
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		val := tempArr[i][0]
		idx := tempArr[i][1]
		if i > 0 && val > tempArr[i-1][0] {
			result[idx] = i
			prev = i
		} else {
			result[idx] = prev
		}
	}

	return result
}

// [8,1,2,2,3]
// [[8,0],[1,1],[2,2],[2,3],[3,4]]
// [[1,1],[2,2],[2,3],[3,4],[8,0]]
// [4,_,1,1,3]
func smallerNumbersThanCurrent3(nums []int) []int {
	tempArr := make([][]int, len(nums))
	for i, num := range nums {
		subArr := make([]int, 2)
		subArr[0] = num
		subArr[1] = i
		tempArr[i] = subArr
	}

	sort.SliceStable(tempArr, func(i, j int) bool { return tempArr[i][0] < tempArr[j][0] })
	result := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if tempArr[i][0] != tempArr[i-1][0] {
			result[tempArr[i][1]] = i
		} else {
			result[tempArr[i][1]] = result[tempArr[i-1][1]]
		}
	}

	return result
}

// [8,1,2,2,3]
// [0,1,2,1,0,0,0,0,8]
// [0,0,1,3,4,4,4,4,4]
// [0,1,3,4,4,4,4,4,8]
// => [4,0,1,1,3]
// TC: O(max(arr) + n), SC: O(max(arr))
func smallerNumbersThanCurrent4(nums []int) []int {
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	countArr := make([]int, max+1)
	for _, num := range nums {
		countArr[num]++
	}

	prefixSum := 0
	for i, num := range countArr {
		countArr[i] = prefixSum
		prefixSum += num
	}

	for i, num := range nums {
		nums[i] = countArr[num]
	}

	return nums
}
