package medium

import "sort"

/*
56. Merge Intervals https://leetcode.com/problems/merge-intervals/
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/

func mergeIntervals(intervals [][]int) [][]int {
	left := make([]int, 0, len(intervals))
	rightSet := make(map[int]int, len(intervals))
	for _, interval := range intervals {
		left = append(left, interval[0])
		if rightSet[interval[0]] < interval[1] {
			rightSet[interval[0]] = interval[1]
		}
	}
	sort.Ints(left)
	res := make([][]int, 0)
	for i := 0; i < len(left); {
		interval := []int{left[i], rightSet[left[i]]}
		j := i + 1
		for ; j < len(left); j++ {
			if left[j] <= interval[1] {
				if interval[1] < rightSet[left[j]] {
					interval[1] = rightSet[left[j]]
				}
			} else {
				break
			}
		}
		res = append(res, interval)
		i = j
	}

	return res
}

func mergeIntervals2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	j := 0
	for i := 1; i < len(intervals); i++ {
		if overlap(intervals[j], intervals[i]) {
			intervals[j][1] = max(intervals[j][1], intervals[i][1])
		} else {
			j += 1
			intervals[j] = intervals[i]
		}
	}

	return intervals[:min(j+1, len(intervals))]
}

func overlap(interval1 []int, interval2 []int) bool {
	return interval1[1] >= interval2[0] && interval1[0] <= interval2[1]
}
