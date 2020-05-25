package hard

import "github.com/diinvoke/leetcode/easy"

/*
https://zhuanlan.zhihu.com/p/77666061
123. Best Time to Buy and Sell Stock III https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	k := 2
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, k+1)
	}
	for j := 1; j <= k; j++ {
		m := prices[0]
		for i := 1; i < len(prices); i++ {
			m = min(prices[i]-dp[i][j-1], m)
			dp[i][j] = max(dp[i-1][j], prices[i]-m)
		}
	}

	return dp[len(prices)-1][k]
}

func maxProfit4(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	if k > len(prices) {
		return easy.MaxProfit2(prices)
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, k+1)
	}
	for j := 1; j <= k; j++ {
		m := prices[0]
		for i := 1; i < len(prices); i++ {
			m = min(prices[i]-dp[i][j-1], m)
			dp[i][j] = max(dp[i-1][j], prices[i]-m)
		}
	}

	return dp[len(prices)-1][k]
}
