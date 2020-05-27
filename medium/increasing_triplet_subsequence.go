package medium

import "math"

/*
334. Increasing Triplet Subsequence https://leetcode.com/problems/increasing-triplet-subsequence/description/

Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.

Formally the function should:

Return true if there exists i, j, k
such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.

Example 1:

Input: [1,2,3,4,5]
Output: true
Example 2:

Input: [5,4,3,2,1]
Output: false
*/

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	firstMin, secondMin := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= firstMin {
			firstMin = num
		}
		if num > firstMin && num <= secondMin {
			secondMin = num
		}
		if num > secondMin {
			return true
		}
	}

	return false
}
