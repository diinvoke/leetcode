package easy

/*
69. Sqrt(x) https://leetcode.com/problems/sqrtx/
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/

// xi+1=xi - (xi2 - n) / (2xi) = xi - xi / 2 + n / (2xi) = xi / 2 + n / 2xi = (xi + n/xi) / 2
func mySqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

func mySqrt2(x int) int {
	left, right := 0, x

	if left > right {
		mid := left + (right-left)/2
		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right - 1

}
