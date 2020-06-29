package easy

/*
67. Add Binary https://leetcode.com/problems/add-binary/

Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

Each string consists only of '0' or '1' characters.
1 <= a.length, b.length <= 10^4
Each string is either "0" or doesn't contain any leading zero.
*/

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	res := make([]byte, max(lenA, lenB)+1)
	j := len(res) - 1
	carry := 0
	for ai, bi := lenA-1, lenB-1; ai >= 0 || bi >= 0; ai, bi = ai-1, bi-1 {
		avi, bvi := 0, 0
		if ai >= 0 {
			if a[ai] == '1' {
				avi = 1
			}
		}
		if bi >= 0 && b[bi] == '1' {
			bvi = 1
		}

		v := (avi + bvi + carry) % 2
		carry = (avi + bvi + carry) / 2
		if v == 1 {
			res[j] = '1'
		} else {
			res[j] = '0'
		}
		j--
	}

	if carry == 1 {
		res[0] = '1'
		return string(res)
	}
	return string(res[1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
