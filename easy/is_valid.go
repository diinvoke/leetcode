package easy

/*
https://leetcode.com/problems/valid-parentheses/

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1.Open brackets must be closed by the same type of brackets.
2.Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:
Input: "()"
Output: true

Example 2:
Input: "()[]{}"
Output: true

Example 3:
Input: "(]"
Output: false

Example 4:
Input: "([)]"
Output: false

Example 5:
Input: "{[]}"
Output: true
*/

func IsValid(s string) bool {
	return isValid(s)
}

func isValid(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return true
	}
	if sLen%2 != 0 {
		return false
	}

	stack := make([]rune, 0, sLen/2)
	for _, symbol := range s {
		if symbol == '(' || symbol == '{' || symbol == '[' {
			stack = append(stack, symbol)
			continue
		}
		if len(stack) == 0 {
			return false
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if getParenther(symbol) != top {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func getParenther(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	}

	return -1
}
