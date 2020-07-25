package medium

import "sort"

/*
90. Subsets II https://leetcode.com/problems/subsets-ii/

Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/

func subsetsWithDup(nums []int) [][]int {
	subset := make([]int, 0)
	subsets := make([][]int, 0)
	sort.Ints(nums)
	subsetsWithDupRecursive(len(nums), 0, nums, &subset, &subsets)
	return subsets
}

func subsetsWithDupRecursive(length, start int, nums []int, subset *[]int, subsets *[][]int) {
	t := make([]int, len(*subset))
	copy(t, *subset)
	*subsets = append(*subsets, t)
	if len(*subset) == length {
		return
	}
	for i := start; i < length; i++ {
		*subset = append(*subset, nums[i])
		subsetsWithDupRecursive(length, i+1, nums, subset, subsets)
		*subset = (*subset)[:len(*subset)-1]
		for i+1 < length && nums[i] == nums[i+1] {
			i++
		}
	}
}
