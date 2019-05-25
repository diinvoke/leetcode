package medium

/*
https://leetcode.com/problems/permutations/
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
func permute(nums []int) [][]int {
	var result [][]int
	fullPermute(nums, 0, len(nums)-1, &result)
	return result
}

func fullPermute(nums []int, left, end int, result *[][]int) {
	if left == end {
		a := make([]int, 0, len(nums))
		a = append(a, nums...)
		*result = append(*result, a)
	}

	for i := left; left < end && i <= end; i++ {
		nums[i], nums[left] = nums[left], nums[i]
		fullPermute(nums, left+1, end, result)
		nums[i], nums[left] = nums[left], nums[i]
	}
}
