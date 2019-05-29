package medium

/*
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/

func searchRange(nums []int, target int) []int {
	left := binarySearchIndex(nums, target, true)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	right := binarySearchIndex(nums, target, false)

	return []int{left, right - 1}
}

func binarySearchIndex(nums []int, target int, isLeft bool) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target || isLeft && nums[mid] == target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
