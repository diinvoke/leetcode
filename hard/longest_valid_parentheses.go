package hard

/*
https://leetcode.com/problems/longest-valid-parentheses/
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"

Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
*/

func longestValidParentheses(s string) int {
	stack := new(simpleStack)
	stack.push(-1) // 窗口开始位置
	var max int

	for index, character := range s {
		if character == '(' {
			stack.push(index)
			continue
		}

		stack.pop()
		// 栈空表示窗口需要左移动
		if stack.empty() {
			stack.push(index)
			continue
		}
		windowL := stack.peek()
		if index-windowL > max {
			max = index - windowL
		}
	}

	return max
}

type simpleStack struct {
	array []int
}

func (s *simpleStack) push(v int) {
	s.array = append(s.array, v)
}

func (s *simpleStack) pop() int {
	if len(s.array) == 0 {
		return 0
	}
	k := s.array[len(s.array)-1]
	if len(s.array)-1 > 0 {
		s.array = s.array[:len(s.array)-1]
	} else {
		s.array = []int{}
	}

	return k
}

func (s *simpleStack) empty() bool {
	return len(s.array) == 0
}

func (s *simpleStack) clean() {
	s.array = []int{}
}

func (s *simpleStack) peek() int {
	if len(s.array) == 0 {
		return 0
	}

	return s.array[len(s.array)-1]
}

func (s *simpleStack) count() int {
	return len(s.array)
}
