package easy

/*
387. First Unique Character in a String https://leetcode.com/problems/first-unique-character-in-a-string/description/

Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
*/

func firstUniqChar(s string) int {
	set := make(map[rune]int, len(s))
	for _, char := range s {
		set[char] += 1
	}
	for idx, char := range s {
		if set[char] == 1 {
			return idx
		}
	}
	return -1
}
