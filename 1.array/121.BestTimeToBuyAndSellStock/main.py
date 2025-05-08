from typing import List

# https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

# on each day:
# - tracking the min price in the past
# - calculate profit = current price - min price (in the past)

# TC: O(n), SC: O(1)
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        minPrice = prices[0]
        profit = 0
        for i in range(1, len(prices)):
            profit = max(profit, prices[i] - minPrice)
            minPrice = min(minPrice, prices[i])
        return profit
