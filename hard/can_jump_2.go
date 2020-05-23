package hard

/*
45. Jump Game II https://leetcode.com/problems/jump-game-ii/description/

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

Example:

Input: [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
    Jump 1 step from index 0 to 1, then 3 steps to the last index.
Note:

You can assume that you can always reach the last index.
*/

func jump(nums []int) int {
	resIdx, curIdx, i := 0, 0, 0
	for curIdx < len(nums)-1 {
		resIdx++
		pre := curIdx
		for ; i <= pre; i++ {
			curIdx = max(curIdx, i+nums[i])
		}
	}
	return resIdx
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//
//	return b
//}

func max(a, b int) int {
	if min(a, b) == a {
		return b
	}

	return a
}
