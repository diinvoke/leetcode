package medium

/*
229. Majority Element II https://leetcode.com/problems/majority-element-ii/description/

Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Note: The algorithm should run in linear time and in O(1) space.

Example 1:

Input: [3,2,3]
Output: [3]

Example 2:

Input: [1,1,1,3,3,2,2,2]
Output: [1,2]
*/

func MajorityElement(nums []int) []int {
	aCnt, bCnt := 0, 0
	a, b := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == a {
			aCnt += 1
			continue
		}
		if nums[i] == b {
			bCnt += 1
			continue
		}

		if aCnt == 0 {
			a = nums[i]
			aCnt = 1
			continue
		}
		if bCnt == 0 {
			b = nums[i]
			bCnt = 1
			continue
		}

		aCnt -= 1
		bCnt -= 1
	}

	aCnt = 0
	bCnt = 0
	for _, num := range nums {
		if num == a {
			aCnt += 1
		}
		if num == b {
			bCnt += 1
		}
	}

	ret := make([]int, 0, 2)
	if aCnt > len(nums)/3 {
		ret = append(ret, a)
	}

	if a == b {
		return ret
	}

	if bCnt > len(nums)/3 {
		ret = append(ret, b)
	}
	return ret
}
