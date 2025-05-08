package main

// https://leetcode.com/problems/subarray-sum-equals-k/description/

// sum(j, i) = prefix[i] - prefix[j-1] = k <=> prefix[j-1] = prefix[i] - k
// count[x]: số lần xuất hiện của prefix[j-1]  thỏa mãn prefix[j-1] = prefix[i]-k

// for each i: nums; 
//  - result += số  lượng của prefix[j-1]
//  - lưu count của mỗi prefix tại i để đánh dấu số lần xuất hiện của mỗi prefix sum

func subarraySum(nums []int, k int) int {
	var result = 0
	var count = map[int]int{0:1}

	prefixSum := 0
	// convert nums into array of prefix sum
	for i, num := range nums {
			prefixSum += num
			nums[i] = prefixSum
	}

	for _, prefixSum := range nums {
			result += count[prefixSum - k]
			count[prefixSum]++
	}


	return result
}
