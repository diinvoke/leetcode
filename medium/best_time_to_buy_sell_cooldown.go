package medium

import "math"

/*
309. Best Time to Buy and Sell Stock with Cooldown https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
Example:

Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buy, sell := math.MinInt16, 0
	preBuy, preSell := 0, 0
	for i := 0; i < len(prices); i++ {
		price := prices[i]
		preBuy = buy
		buy = max(preSell-price, preBuy)
		preSell = sell
		sell = max(preBuy+price, preSell)
	}

	return sell
}
