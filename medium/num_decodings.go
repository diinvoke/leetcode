package medium

/*
91. Decode Ways https://leetcode.com/problems/decode-ways/

A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i, c := range s {
		if c == '0' {
			dp[i+1] = 0
		} else {
			dp[i+1] = dp[i]
		}

		if i >= 1 && (s[i-1] == '1' || (s[i-1] == '2' && c <= '6')) {
			dp[i+1] += dp[i-1]
		}
	}

	return dp[len(s)]
}
