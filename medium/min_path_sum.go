package medium

/*
64. Minimum Path Sum https://leetcode.com/problems/minimum-path-sum/

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = grid[i][0]
		if i != 0 {
			dp[i][0] += dp[i-1][0]
		}
	}
	for i := 0; i < n; i++ {
		dp[0][i] = grid[0][i]
		if i != 0 {
			dp[0][i] += dp[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cost := grid[i][j]
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + cost
		}
	}

	return dp[m-1][n-1]
}
