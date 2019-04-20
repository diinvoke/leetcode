package medium

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:
Input: "cbbd"
Output: "bb"
*/

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var start, maxLen int

	for index := range s {
		start, maxLen = search(s, index, index, start, maxLen)
		start, maxLen = search(s, index, index+1, start, maxLen)
	}

	end := start + maxLen
	result := s[start:end]
	return string(result)
}

func search(s string, left, right, start, maxLen int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}

	if right-left-1 > maxLen {
		return left + 1, right - left - 1
	}

	return start, maxLen
}
