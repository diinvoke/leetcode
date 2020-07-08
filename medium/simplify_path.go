package medium

import (
	"fmt"
	"strings"
)

/*
71. Simplify Path https://leetcode.com/problems/simplify-path/
Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level.

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) must not end with a trailing /. Also, the canonical path must be the shortest string representing the absolute path.



Example 1:

Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
Example 2:

Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
Example 3:

Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
Example 4:

Input: "/a/./b/../../c/"
Output: "/c"
Example 5:

Input: "/a/../../b/../c//.//"
Output: "/c"
Example 6:

Input: "/a//b////c/d//././/.."
Output: "/a/b/c"
*/

func simplifyPath(path string) string {
	stack := new(simpleStack)

	for _, str := range strings.Split(path, "/") {
		if str == "." || str == " " || str == "" {
			continue
		}
		if str == ".." {
			stack.pop()
			continue
		}
		stack.push(str)
	}

	return fmt.Sprintf("/%s", strings.Join(stack.origin(), "/"))
}

type simpleStack struct {
	array []string
}

func (s *simpleStack) push(v string) {
	s.array = append(s.array, v)
}

func (s *simpleStack) pop() string {
	if len(s.array) == 0 {
		return ""
	}
	k := s.array[len(s.array)-1]
	if len(s.array)-1 > 0 {
		s.array = s.array[:len(s.array)-1]
	} else {
		s.array = []string{}
	}

	return k
}

func (s *simpleStack) origin() []string {
	return s.array
}
