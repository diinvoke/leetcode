package medium

/*
https://leetcode.com/problems/permutations-ii/
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:

Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

func permuteUnique(nums []int) [][]int {
	var result [][]int
	fullPermuteUnique(nums, 0, len(nums)-1, &result)
	return result
}

func fullPermuteUnique(nums []int, left, end int, result *[][]int) {
	if left == end {
		a := make([]int, 0, len(nums))
		a = append(a, nums...)
		*result = append(*result, a)
	}

	for i := left; i <= end; i++ {
		if permuteUniqueIgnore(nums, i, end) {
			continue
		}

		nums[i], nums[left] = nums[left], nums[i]
		fullPermuteUnique(nums, left+1, end, result)
		nums[i], nums[left] = nums[left], nums[i]
	}
}

func permuteUniqueIgnore(nums []int, start, end int) bool {
	for i := start + 1; i <= end; i++ {
		if nums[i] == nums[start] {
			return true
		}
	}

	return false
}
