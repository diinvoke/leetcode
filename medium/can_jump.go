package medium

/*
55. Jump Game https://leetcode.com/problems/jump-game/description/

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.



Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/

func CanJump(nums []int) bool {
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i-1]) - 1
		if dp[i] < 0 {
			return false
		}
	}
	return true
}

func CanJump2(nums []int) bool {
	reachIndex := 0
	for i, num := range nums {
		if i > reachIndex || reachIndex >= len(nums)-1 {
			break
		}

		reachIndex = max(reachIndex, i+num)
	}

	return reachIndex >= len(nums)-1
}
