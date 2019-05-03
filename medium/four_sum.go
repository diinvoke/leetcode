package medium

import (
	"github.com/diinvoke/leetcode/tools"
)

/*
https://leetcode.com/problems/4sum/
Given an array nums of n integers and an integer target,
are there elements a, b, c, and d in nums such that a + b + c + d = target?
Find all unique quadruplets in the array which gives the sum of target.

Note:
The solution set must not contain duplicate quadruplets.

Example:
Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func FourSum(nums []int, target int) [][]int {
	return fourSum(nums, target)
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	tools.QuickSort(nums)

	numLen := len(nums)
	for i := 0; i < numLen-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < numLen-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, numLen-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
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

				if sum < target {
					left += 1
				} else {
					right -= 1
				}
			}
		}
	}

	return result
}
