package medium

import "github.com/diinvoke/leetcode/tools"

/*
https://leetcode.com/problems/3sum-closest/

Given an array nums of n integers and an integer target,
find three integers in nums such that the sum is closest to target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/

func ThreeSumClosest(nums []int, target int) int {
	return threeSumClosest(nums, target)
}

func threeSumClosest1(nums []int, target int) int {
	tools.QuickSort(nums)
	var closest int

	for i := 0; i < len(nums) && i < 3; i++ {
		closest += nums[i]
	}

	if len(nums) <= 3 {
		return closest
	}

	min := abs(target - closest)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(target-sum) < min {
				min = abs(target - sum)
				closest = sum
			}

			if sum < target {
				left += 1
			} else {
				right -= 1
			}
		}
	}

	return closest
}

func threeSumClosest(nums []int, target int) int {
	tools.QuickSort(nums)
	var closest int

	for i := 0; i < len(nums) && i < 3; i++ {
		closest += nums[i]
	}

	if len(nums) <= 3 {
		return closest
	}

	for i := 0; i < len(nums)-2; i++ {
		min := nums[i] + closestFind(nums, i+1, len(nums)-1, target-nums[i])
		if abs(target-min) < abs(target-closest) {
			closest = min
		}
	}

	return closest
}

func closestFind(nums []int, left, right, target int) int {
	closest := nums[left] + nums[right]
	min := abs(target - closest)
	for left < right {
		sum := nums[left] + nums[right]
		if abs(target-sum) < min {
			min = abs(target - sum)
			closest = sum
		}

		if sum < target {
			left += 1
		} else {
			right -= 1
		}
	}

	return closest
}

func closestFind1(nums []int, left, right, target int) int {
	result := nums[left] + nums[right]
	min := target - result

	for left < right {
		if abs(target-(nums[left]+nums[right])) <= abs(min) {
			result = nums[left] + nums[right]
			min = target - result
		}

		if target < (nums[left] + nums[right]) {
			right -= 1
			continue
		}
		if target >= (nums[left] + nums[right]) {
			left += 1
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
