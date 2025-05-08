package main

// https://leetcode.com/problems/range-sum-query-immutable/

type NumArray struct {
	prefixSum []int
}

// precalculate the sum of leftside
// [-2,0,3,-5,2,-1]
// [-2,-2,1,-4,-2,-3]
func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0
	for i := 1; i< len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	return NumArray{
		prefixSum: prefixSum,
	}
}

// sum[left:right] = prefixSum[right] - prefixSum[left]
func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

/**
* Your NumArray object will be instantiated and called as such:
* obj := Constructor(nums);
* param_1 := obj.SumRange(left,right);
 */
