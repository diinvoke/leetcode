package hard

/*
316. Remove Duplicate Letters https://leetcode.com/problems/remove-duplicate-letters/

Given a string which contains only lowercase letters, remove duplicate letters so that every letter appears once and only once. You must make sure your result is the smallest in lexicographical order among all possible results.

Example 1:

Input: "bcabc"
Output: "abc"
Example 2:

Input: "cbacdcbc"
Output: "acdb"
Note: This question is the same as 1081: https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
*/

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for _, sr := range s {
		count[sr-'a']++
	}

	stack := new(simpleStack)
	used := [26]bool{}
	for _, sr := range s {
		count[sr-'a']--
		if used[sr-'a'] {
			continue
		}

		for !stack.empty() && stack.peek() > int(sr) && count[stack.peek()-'a'] > 0 {
			used[stack.peek()-'a'] = false
			stack.pop()
		}

		stack.push(int(sr))
		used[sr-'a'] = true
	}

	resb := make([]rune, stack.count())
	idx := stack.count() - 1
	for !stack.empty() {
		resb[idx] = rune(stack.pop())
		idx -= 1
	}
	return string(resb)
}
