package medium

import "fmt"

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var left int
	var tl, tr int
	var window = make(map[int32]int, len(s)/2)

	for index, val := range s {
		lastIndex, ok := window[val]
		if ok && lastIndex >= left {
			left = lastIndex + 1
		}

		window[val] = index
		if index-left > tr-tl {
			//res = index - left + 1
			tl = left
			tr = index
		}

	}

	fmt.Println(tl, string(s[tl:tr+1]))
	//fmt.Println(res)

	return tr - tl + 1
}
