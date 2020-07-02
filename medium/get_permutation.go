package medium

import (
	"bytes"
	"strconv"
)

/*
60. Permutation Sequence
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:

Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.
Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"
*/

func getPermutation(n int, k int) string {
	ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

	buf := bytes.Buffer{}
	total := n
	k -= 1
	for total > 0 {
		idx := k / f[total-1]
		buf.WriteString(strconv.FormatInt(int64(ns[idx]), 10))
		ns = erase(ns, idx)

		k %= f[total-1]
		total -= 1
	}

	return buf.String()
}

func erase(ns []int, idx int) []int {
	if idx >= len(ns)-1 {
		return ns[:len(ns)-1]
	}
	return append(ns[:idx], ns[idx+1:]...)
}
