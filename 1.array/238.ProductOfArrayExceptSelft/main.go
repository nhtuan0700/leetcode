package main

// https://leetcode.com/problems/product-of-array-except-self/description/

// [1,2,3,4]
// i = 0; 2*3*4
// i = 1; 1*3*4
// i = 2; 1*2*4

// overview = productLefts[i] + productRight
// keyword: except itselft

// result = [1,1,2,6] (each i is product of nums at left i except itself)
// productRight = 1
// i: n - 1 -> 0, result[i]*=productRight,productRight*=nums[i]
// i = 3, result[3]=1*6=6, productRight=4
// i = 2, result[2]=4*2=8, productRight=12
// i = 1, result[1]=12*1=12, productRight=24
func productExceptSelf(nums []int) []int {
	productLeft := 1
	productRight := 1
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = productLeft
		productLeft *= num
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= productRight
		productRight *= nums[i]
	}

	return result
}
