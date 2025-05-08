package main

// https://leetcode.com/problems/subarray-sums-divisible-by-k/

// [4,5,0,-2,-3,1]
// [4,9,9,7,4,5]
// prefixSums = ...

// (prefix[i] - prefix[j-1]) % k = 0 <=> prefix[i]%k == prefix[j-1]%k
// count[x]: number of mod of perfix[i]

func subarraysDivByK(nums []int, k int) int {
	for i := 1; i < len(nums); i++ {
			nums[i] += nums[i-1]
	}

	result := 0
	count := map[int]int{0:1}
	for _, prefixSum := range nums {
			mod := (prefixSum%k + k) %k 
			result += count[mod]
			count[mod]++
	}


	return result
}
