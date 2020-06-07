package easy

/*
242. Valid Anagram https://leetcode.com/problems/valid-anagram/description/
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := [26]int{}
	for _, c := range s {
		m[c-'a'] += 1
	}
	for _, c := range t {
		m[c-'a'] -= 1
		if m[c-'a'] < 0 {
			return false
		}
	}

	return true
}
