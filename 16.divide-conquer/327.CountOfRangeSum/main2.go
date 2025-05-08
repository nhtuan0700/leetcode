package main

// prefix sum and brute force
// TC: O(n^2), SC: O(n)
func countRangeSum2(nums []int, lower int, upper int) int {
	n := len(nums)
	sums := make([]int, n+1)

	for i := range nums {
		sums[i+1] = nums[i] + sums[i]
	}

	res := 0

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			diff := sums[j] - sums[i-1]
			if diff >= lower && diff <= upper {
				res++
			}
		}
	}

	return res
}
