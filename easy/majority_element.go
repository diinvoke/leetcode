package easy

/*
169. Majority Element https://leetcode.com/problems/majority-element/description/

Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2
*/

func majorityElement(nums []int) int {
	set := make(map[int]int)

	for _, num := range nums {
		set[num] += 1
	}

	ret, max := 0, 0
	for num, v := range set {
		if max <= v {
			max = v
			ret = num
		}
	}
	return ret
}

func majorityElement2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	m := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == m {
			count++
		} else {
			count--
		}

		if count == 0 {
			m = nums[i]
			count = 1
		}
	}

	return m

}
