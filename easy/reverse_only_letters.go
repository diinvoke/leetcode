package easy

/*
https://leetcode.com/problems/reverse-only-letters/
Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.

Example 1:

Input: "ab-cd"
Output: "dc-ba"

Example 2:

Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"

Example 3:

Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"
*/

func reverseOnlyLetters(S string) string {
	ss := make([]uint8, 0, len(S))
	ss = append(ss, S...)
	left, right := 0, len(S)-1
	for left < right {
		lv := ss[left]
		rv := ss[right]

		if !isLetter(lv) {
			left += 1
			continue
		}
		if !isLetter(rv) {
			right -= 1
			continue
		}

		ss[left], ss[right] = ss[right], ss[left]
		left += 1
		right -= 1
	}

	return string(ss)
}

func isLetter(c uint8) bool {
	return (c >= 65 && c <= 90) || (c >= 97 && c <= 122)
}
