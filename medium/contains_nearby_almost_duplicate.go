package medium

import (
	"sort"
)

/*
220. Contains Duplicate III https://leetcode.com/problems/contains-duplicate-iii/description/

Given an array of integers, find out whether there are two distinct indices i and j in the array such that the
absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3, t = 0
Output: true

Example 2:

Input: nums = [1,0,1,1], k = 1, t = 2
Output: true

Example 3:

Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false
*/

func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	s := new(set)
	for i, num := range nums {
		f := s.floor(num)
		c := s.cell(num)

		if (f != nil && num-*f <= t) || (c != nil && *c-num <= t) {
			return true
		}

		s.add(num)
		if i >= k {
			s.remove(nums[i-k])
		}
	}
	return false
}

type set struct {
	nums []int
}

func (s *set) add(num int) {
	s.nums = append(s.nums, num)
	sort.Ints(s.nums)
}

func (s *set) remove(num int) {
	for i, n := range s.nums {
		if n == num {
			if i >= len(s.nums)-1 {
				s.nums = s.nums[:i]
				return
			}

			s.nums = append(s.nums[:i], s.nums[i+1:]...)
		}
	}
}

func (s *set) floor(x int) *int {
	if len(s.nums) == 0 || s.nums[0] > x {
		return nil
	}

	for i := len(s.nums) - 1; i >= 0; i-- {
		n := s.nums[i]
		if n <= x {
			return &n
		}
	}

	return nil
}

func (s *set) cell(x int) *int {
	if len(s.nums) == 0 || s.nums[len(s.nums)-1] < x {
		return nil
	}

	for _, n := range s.nums {
		if n >= x {
			return &n
		}
	}

	return nil
}
