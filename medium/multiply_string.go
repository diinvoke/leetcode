package medium

import (
	"bytes"
	"fmt"
)

/*
https://leetcode.com/problems/multiply-strings/
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Example 1:
Input: num1 = "2", num2 = "3"
Output: "6"

Example 2:
Input: num1 = "123", num2 = "456"
Output: "56088"
Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contain only digits 0-9.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

func multiply2(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	num1Len, num2Len := len(num1), len(num2)
	result := make([]int, num2Len+num1Len)
	for i := num1Len - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := num2Len - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			low, high := i+j+1, i+j
			sum := n1*n2 + result[low]
			result[low] = sum % 10
			result[high] += sum / 10
		}
	}

	fmt.Println(result)

	var buf bytes.Buffer
	for _, n := range result {
		if buf.Len() > 0 || n != 0 {
			buf.WriteRune(rune(n + '0'))
		}
	}

	if buf.Len() == 0 {
		return "0"
	}

	return buf.String()
}

func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	num1Len, num2Len := len(num1), len(num2)

	var result [][]int
	for i := num1Len - 1; i >= 0; i-- {
		re := make([]int, num1Len-1-i)
		carry := 0
		n1 := int(num1[i] - '0')
		for j := num2Len - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			sum := n1*n2 + carry
			re = append(re, sum%10)
			carry = sum / 10
		}
		if carry > 0 {
			re = append(re, carry)
		}
		if !zeroList(re) {
			result = append(result, re)
		}
	}

	size := len(result)
	if size == 0 {
		return "0"
	}
	interval := 1
	for interval < size {
		for i := 0; i < size-interval; i += interval * 2 {
			result[i] = addTwo(result[i], result[i+interval])
		}
		interval *= 2
	}

	resultInt := result[0]
	var buf bytes.Buffer
	for i := len(resultInt) - 1; i >= 0; i-- {
		n := resultInt[i]
		buf.WriteString(fmt.Sprintf("%d", n))
	}

	re := buf.String()
	if re == "" {
		re = "0"
	}

	return re

}

func zeroList(list []int) bool {
	for _, n := range list {
		if n != 0 {
			return false
		}
	}

	return true
}

func addTwo(num1 []int, num2 []int) []int {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	num1Len, num2Len := len(num1), len(num2)
	result := make([]int, 0, num1Len)
	j := 0
	carry := 0
	for i := 0; i < num1Len; i++ {
		sum := carry + num1[i]
		if j < num2Len {
			sum += num2[j]
			j++
		}

		carry = sum / 10
		result = append(result, sum%10)
	}

	if carry > 0 {
		result = append(result, carry)
	}

	return result
}
