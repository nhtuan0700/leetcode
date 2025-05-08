package main

import "fmt"

// https://leetcode.com/problems/remove-element/description/

// input: prices, prices[i]: stock on i day
// maxProfit = prices[j] - prices[i], j: sell day, i: buy day, i < j
// or maxProfit = max(prices) - min(prices)
// At each time (prices[i]), you will sell with the min price in the past to find the maximum of profit
// ex: [7,1,5,3,6,4]
// TC: O(n)
// SC: O(1)

func maxProfit(prices []int) int {
	minPriceSofar := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > minPriceSofar {
			maxProfit = max(maxProfit, prices[i]-minPriceSofar)
		}
		minPriceSofar = min(minPriceSofar, prices[i])
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
