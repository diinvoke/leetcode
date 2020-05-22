package easy

/*
217. Contains Duplicate https://leetcode.com/problems/contains-duplicate/description/

Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:

Input: [1,2,3,1]
Output: true
Example 2:

Input: [1,2,3,4]
Output: false


Example 3:

Input: [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

func ContainsDuplicate(nums []int) bool {
	numTable := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := numTable[num]; ok {
			return true
		}
		numTable[num] = struct{}{}
	}

	return false
}

/*
219. Contains Duplicate II https://leetcode.com/problems/contains-duplicate-ii/description/
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the
absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

func ContainsNearbyDuplicate(nums []int, k int) bool {
	numTable := make(map[int][]int)
	for i, num := range nums {
		ij, ok := numTable[num]
		if !ok || len(ij) == 1 {
			numTable[num] = append(numTable[num], i)
			continue
		}

		if len(ij) == 2 {
			if ij[1]-ij[0] > i-ij[1] {
				ij[0], ij[1] = ij[1], i
			}
		}
	}

	for _, ij := range numTable {
		if len(ij) != 2 {
			continue
		}

		if ij[1]-ij[0] <= k {
			return true
		}
	}

	return false
}

func ContainsNearbyDuplicate2(nums []int, k int) bool {
	if len(nums) < 2 || k < 1 {
		return false
	}

	set := make(map[int]int)
	for i, num := range nums {
		if j, ok := set[num]; (i-j) <= k && ok {
			return true
		}
		set[num] = i
	}
	return false
}
