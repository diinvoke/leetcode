package easy

/*
https://leetcode.com/problems/implement-strstr/

Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen > haystackLen {
		return -1
	}

	if needleLen == 0 {
		return 0
	}

	c := needle[0]
	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i] != c {
			continue
		}

		if i+needleLen <= haystackLen && haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}
