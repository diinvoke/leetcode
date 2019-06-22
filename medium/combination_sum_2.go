package medium

import "sort"

/*
https://leetcode.com/problems/combination-sum-ii/
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

Example 2:
Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]
*/

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var cur []int
	sort.Ints(candidates)
	findCombinationSum2(candidates, 0, target, &cur, &result)

	return result
}

func findCombinationSum2(candidates []int, start, target int, cur *[]int, result *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		res := make([]int, 0, len(*cur))
		res = append(res, *cur...)
		*result = append(*result, res)
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		*cur = append(*cur, candidates[i])
		findCombinationSum2(candidates, i+1, target-candidates[i], cur, result)
		r := *cur
		r = r[0 : len(r)-1]
		*cur = r
	}
}
