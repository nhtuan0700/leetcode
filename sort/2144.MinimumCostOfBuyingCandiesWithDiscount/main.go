package main

// https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/

import "sort"

// O (nlog)
func minimumCost(cost []int) int {
	sort.Ints(cost)

	result := 0
	buyNum := 0
	for i := len(cost) - 1; i >= 0; i-- {
		if buyNum < 2 {
			result += cost[i]
			buyNum++
		} else {
			buyNum = 0
		}
	}

	return result
}
