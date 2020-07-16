package medium

/*
77. Combinations https://leetcode.com/problems/combinations/

Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {
	combine := make([]int, 0)
	combines := make([][]int, 0)
	combineRecursive(n, 1, k, &combine, &combines)
	return combines
}

func combineRecursive(n, start, k int, combine *[]int, combines *[][]int) {
	if len(*combine) == k {
		t := make([]int, k)
		copy(t, *combine)
		*combines = append(*combines, t)
		return
	}
	for i := start; i <= n; i++ {
		*combine = append(*combine, i)
		combineRecursive(n, i+1, k, combine, combines)
		*combine = (*combine)[:len(*combine)-1]
	}
}
