package medium

/*
151. Reverse Words in a String https://leetcode.com/problems/reverse-words-in-a-string/description/

Given an input string, reverse the string word by word.



Example 1:

Input: "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Note:

A word is defined as a sequence of non-space characters.
Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
You need to reduce multiple spaces between two words to a single space in the reversed string.
*/

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	chars := []byte(s)
	reverseString(chars)
	storeIdx := 0
	for i := 0; i < len(s); i++ {
		if chars[i] == ' ' {
			continue
		}
		if storeIdx != 0 {
			chars[storeIdx] = ' '
			storeIdx += 1
		}

		j := i
		for j < len(s) && chars[j] != ' ' {
			chars[storeIdx] = chars[j]
			storeIdx += 1
			j += 1
		}
		reverseString(chars[storeIdx-(j-i) : storeIdx])
		i = j
	}
	return string(chars[:storeIdx])
}

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left += 1
		right -= 1
	}
}
