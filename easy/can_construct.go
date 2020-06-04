package easy

/*
383. Ransom Note https://leetcode.com/problems/ransom-note/description/

Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

You may assume that both strings contain only lowercase letters.
*/

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	set := make(map[rune]int, len(magazine))
	for _, char := range magazine {
		set[char] += 1
	}
	for _, char := range ransomNote {
		if set[char] <= 0 {
			return false
		}
		set[char] -= 1
	}
	return true
}
