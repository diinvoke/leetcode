package medium

/*
81. Search in Rotated Sorted Array II https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
Follow up:

This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
Would this affect the run-time complexity? How and why?
*/

func searchII(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[right] { // 左边有序
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] { // 右边有序
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // 相等，最右值向左移动
			right -= 1
		}
	}
	return false
}
