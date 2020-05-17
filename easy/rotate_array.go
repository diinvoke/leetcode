package easy

/*
189. Rotate Array https://leetcode.com/problems/rotate-array/description/

Given an array, rotate the array to the right by k steps, where k is non-negative.

Follow up:
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?

Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/

func RotateArray(nums []int, k int) {
	temp := nums
	for i := 0; i < k; i++ {
		temp = append(temp[len(temp)-1:], temp[0:len(temp)-1]...)
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = temp[i]
	}
}

func RotateArray2(nums []int, k int) {
	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a[(i+k)%len(nums)] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = a[i]
	}
}
