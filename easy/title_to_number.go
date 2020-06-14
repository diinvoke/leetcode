package easy

import "math"

/*
171. Excel Sheet Column Number https://leetcode.com/problems/excel-sheet-column-number/description/

Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
Example 1:

Input: "A"
Output: 1
Example 2:

Input: "AB"
Output: 28
Example 3:

Input: "ZY"
Output: 701
*/

func titleToNumber(s string) int {
	power := float64(0)
	res := 0
	sr := []rune(s)
	for idx := len(s) - 1; idx >= 0; idx-- {
		res += int(math.Pow(26, power)) * int((sr[idx]-'A')+1)
		power += 1
	}
	return res
}
