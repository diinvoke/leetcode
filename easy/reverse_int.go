package easy

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
Input: 123
Output: 321

Example 2:
Input: -123
Output: -321

Example 3:
Input: 120
Output: 21

Note:
Assume we are dealing with an environment which could
only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func Reverse(x int) int {
	num := x
	if num < 0 {
		num *= -1
	}

	var result int
	for num > 0 {
		remainderNum := num % 10
		if remainderNum > 0 || result > 0 {
			result = result*10 + remainderNum
		}

		num = num / 10
	}

	if x < 0 {
		result *= -1
	}

	if result > 1<<31-1 || result < -1<<31 {
		return 0
	}

	return result
}
