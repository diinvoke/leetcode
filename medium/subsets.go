package medium

/*
78. Subsets https://leetcode.com/problems/subsets/

Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func subsets(nums []int) [][]int {
	subset := make([]int, 0)
	subsets := make([][]int, 0)
	subsetsRecursive(len(nums), 0, nums, &subset, &subsets)
	return subsets
}

func subsetsRecursive(length, start int, nums []int, subset *[]int, subsets *[][]int) {
	t := make([]int, len(*subset))
	copy(t, *subset)
	*subsets = append(*subsets, t)
	if len(*subset) == length {
		return
	}
	for i := start; i < length; i++ {
		*subset = append(*subset, nums[i])
		subsetsRecursive(length, i+1, nums, subset, subsets)
		*subset = (*subset)[:len(*subset)-1]
	}
}
