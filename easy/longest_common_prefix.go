package easy

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
Input: ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Note:
All given inputs are in lowercase letters a-z.
*/

func LongestCommonPrefix(strs []string) string {
	strLen := len(strs)
	if strLen == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < strLen; i += 1 {
		s := strs[i]
		for !(len(s) >= len(prefix) && s[0:len(prefix)] == prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}

	}

	return prefix
}
