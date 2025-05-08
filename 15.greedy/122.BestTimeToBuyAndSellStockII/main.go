package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/

// Greedy: On each day, if the stock price on the next day is high than the current day's price,
// we will buy the stock today and sell it in next day to gain profit

// TC: O(n), SC: O(1)
func maxProfit(prices []int) int {
	n := len(prices)
	profit := 0
	for i := 0; i < n-1; i++ {
		if prices[i] < prices[i+1] {
			profit += (prices[i+1] - prices[i])
		}
	}

	return profit
}
