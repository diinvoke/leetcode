package easy

/*
119. Pascal's Triangle II https://leetcode.com/problems/pascals-triangle-ii/

Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.

Example:

Input: 3
Output: [1,3,3,1]
Follow up:

Could you optimize your algorithm to use only O(k) extra space?
*/

func PascalTriangleII(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	ret[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			ret[j] = ret[j] + ret[j-1]
		}
	}
	return ret
}
