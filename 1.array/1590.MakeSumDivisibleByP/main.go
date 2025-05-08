package main

// https://leetcode.com/problems/make-sum-divisible-by-p/description/

// prefixSums
// total

// (total - sum(i,j)) % P == 0 <=> total%P == sum(i,j)%P <=> (prefix[i] - prefix[j-1])%P = total%P
// <=> prefix[j-1]%P = prefix[i]%P - total%P <=> (prefix[i]%p-total%P+P)%P

// modMap[x]: position mod of prefix in the past

// [3,1,4,2],p=6
// total = 10, modTotal=4
// prefixSum=[3,4,8,10]

// modMap={0:-1}
// i = 0, modPrefixSum=3,modPrefixSum-modTotal=-1,modMap={3:0}
// i = 1, modPrefixSum=4,modPrefixSum-modTotal=0,modMap={3:0,4:1}
// i = 2, modPrefixSum=2,modPrefixSum-modTotal=-2,modMap={3:0,4:1,2:1}
// i = 3, modPrefixSum=4,modPrefixSum-modTotal=0,modMap={3:0,4:1,2:1,4:3}

// [6,3,5,2], p=9
// total = 16, modTotal=7
// prefixSum=[6,9,14,16]

// modMap={0:-1}
// i = 0, modPrefixSum=6, modPrefixSum-modTotal=-1,modMap={6:0}
// i = 1, modPrefixSum=0, modPrefixSum-modTotal=-7,modMap={6:0,0:1}
// i = 2, modPrefixSum=5, modPrefixSum-modTotal=-7,modMap={6:0,0:1,5:2}
// i = 3, modPrefixSum=7, modPrefixSum-modTotal=0,modMap={6:0,0:1,5:2}
func minSubarray(nums []int, p int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%p == 0 {
		return 0
	}

	prefixSum := 0
	modMap := map[int]int{0: -1}
	minLength := len(nums)
	modTotal := (total%p + p) % p
	for i, num := range nums {
		prefixSum += num
		modPrefixSum := (prefixSum%p + p) % p
		wantedMod := (modPrefixSum - modTotal + p) % p
		if _, ok := modMap[wantedMod]; ok {
			minLength = min(minLength, i-modMap[wantedMod])
		}

		modMap[modPrefixSum] = i
	}

	if minLength < len(nums) {
		return minLength
	}

	return -1
}
