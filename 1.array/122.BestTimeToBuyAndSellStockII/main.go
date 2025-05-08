package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/

// [7,1,4,3,5,4], 1,5=4; 3,6=3
// [7,1,2,5,6,4]

// buy low, sell high
// Whenever the price increases from one day to the next, we treat it as a buy on the previous day and a sell on the current day to earn profit from that rise.
// We repeat this for every such increase.

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}

	return profit
}
