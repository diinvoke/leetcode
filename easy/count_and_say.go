package easy

import (
	"bytes"
	"fmt"
)

/*
https://leetcode.com/problems/count-and-say/
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

Example 1:
Input: 1
Output: "1"

Example 2:
Input: 4
Output: "1211"
*/

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}

	previous := countAndSay(n - 1)
	count := 0
	previousLen := len(previous)
	sb := bytes.Buffer{}
	for i := 0; i < previousLen; i++ {
		count += 1
		if i+1 < previousLen && previous[i] == previous[i+1] {
			continue
		}
		sb.WriteString(fmt.Sprintf("%d", count))
		sb.WriteByte(previous[i])
		count = 0
	}

	return sb.String()
}

func countAndSayLoop(n int) string {
	if n == 0 {
		return ""
	}

	res := "1"
	for n-1 > 0 {
		count := 1
		resLen := len(res)
		sb := bytes.Buffer{}
		for i := 0; i < resLen; i++ {
			if i+1 < resLen && res[i] == res[i+1] {
				count += 1
				continue
			}

			sb.WriteString(fmt.Sprintf("%d", count))
			sb.WriteByte(res[i])
			count = 1
		}

		res = sb.String()
		n -= 1
	}

	return res
}
