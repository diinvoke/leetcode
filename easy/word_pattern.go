package easy

import "strings"

/*
290. Word Pattern https://leetcode.com/problems/word-pattern/description/

Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

Input: pattern = "abba", str = "dog cat cat dog"
Output: true
Example 2:

Input:pattern = "abba", str = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
Example 4:

Input: pattern = "abba", str = "dog dog dog dog"
Output: false
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.
*/

func wordPattern(pattern string, str string) bool {
	ss := strings.Split(str, " ")
	pb := []rune(pattern)
	if len(ss) != len(pb) {
		return false
	}
	p2s := make(map[rune]string, len(pattern)/2)
	s2p := make(map[string]rune, len(ss)/2)

	for i, sss := range ss {
		pr := pb[i]
		if prr, ok := s2p[sss]; ok {
			if pr != prr {
				return false
			}
		} else {
			s2p[sss] = pr
		}

		if srr, ok := p2s[pr]; ok {
			if srr != sss {
				return false
			}
		} else {
			p2s[pr] = sss
		}
	}

	return true
}
