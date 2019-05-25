package medium

/*
https://leetcode.com/problems/next-permutation/
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i -= 1
	}
	j := len(nums) - 1
	if i >= 0 {
		for j >= 0 && nums[j] <= nums[i] {
			j -= 1
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	revert(nums, i+1, len(nums)-1)
}

func revert(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}
}
