package medium

/*
https://leetcode.com/problems/combination-sum/
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.

Example 1:
Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]

Example 2:
Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var cur []int
	findCombinationSum(candidates, 0, target, &cur, &result)

	return result
}

func findCombinationSum(candidates []int, start, target int, cur *[]int, result *[][]int) {
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
		*cur = append(*cur, candidates[i])
		findCombinationSum(candidates, i, target-candidates[i], cur, result)
		r := *cur
		r = r[0 : len(r)-1]
		*cur = r
	}
}
