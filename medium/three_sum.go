package medium

import "github.com/diinvoke/leetcode/tools"

/*
https://leetcode.com/problems/3sum/

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note:
The solution set must not contain duplicate triplets.

Example:
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func ThreeSum(nums []int) [][]int {
	tools.QuickSort(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i += 1 {
		// 相同直接跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		ret := find(nums, i+1, len(nums)-1, nums[i])
		result = append(result, ret...)
	}

	return result

}

func find(nums []int, left, right int, num int) [][]int {
	target := 0 - num
	var ret [][]int
	for left < right {
		if nums[left]+nums[right] == target {
			ret = append(ret, []int{num, nums[left], nums[right]})
			for left < right && nums[left] == nums[left+1] {
				left += 1
			}

			for left < right && nums[right] == nums[right-1] {
				right -= 1
			}

			left += 1
			right -= 1
			continue
		}

		if nums[left]+nums[right] < target {
			left += 1
		} else {
			right -= 1
		}
	}

	return ret
}
