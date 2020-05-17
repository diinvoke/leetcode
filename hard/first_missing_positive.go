package hard

/*
41. First Missing Positive https://leetcode.com/problems/first-missing-positive/description/

Given an unsorted integer array, find the smallest missing positive integer.

Example 1:
Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2

Example 3:
Input: [7,8,9,11,12]
Output: 1

Note:
Your algorithm should run in O(n) time and uses constant extra space.
*/

func FirstMissingPositive(nums []int) int {
	has := make(map[int]struct{})
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
		has[num] = struct{}{}
	}

	for i := 1; i <= max; i++ {
		if _, ok := has[i]; !ok {
			return i
		}
	}

	return max + 1
}

func FirstMissingPositive2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] <= len(nums) && nums[i] > 0 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
