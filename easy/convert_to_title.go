package easy

/*
168. Excel Sheet Column Title https://leetcode.com/problems/excel-sheet-column-title/description/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=transfer2china#

Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...
Example 1:

Input: 1
Output: "A"
Example 2:

Input: 28
Output: "AB"
Example 3:

Input: 701
Output: "ZY"
*/

func convertToTitle(n int) string {
	res := make([]rune, 0)
	for n > 0 {
		if n%26 == 0 {
			res = append(res, 'Z')
			n -= 26
		} else {
			res = append(res, rune(n%26-1+'A'))
			n -= n % 26
		}
		n /= 26
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
