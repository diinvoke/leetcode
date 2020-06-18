package hard

import (
	"bytes"
	"strings"
)

/*
273. Integer to English Words https://leetcode.com/problems/integer-to-english-words/description/
Convert a non-negative integer to its english words representation. Given input is guaranteed to be less than 231 - 1.

Example 1:

Input: 123
Output: "One Hundred Twenty Three"
Example 2:

Input: 12345
Output: "Twelve Thousand Three Hundred Forty Five"
Example 3:

Input: 1234567
Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
Example 4:

Input: 1234567891
Output: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
*/

func numberToWords(num int) string {
	v1 := []string{"Thousand", "Million", "Billion"}
	resb := bytes.Buffer{}
	resb.WriteString(covert3Num(num % 1000))
	for i := 0; i < 3; i++ {
		num /= 1000
		tb := bytes.Buffer{}
		if num%1000 > 0 {
			tb.WriteString(covert3Num(num % 1000))
			tb.WriteString(" ")
			tb.WriteString(v1[i])
			tb.WriteString(" ")
			tb.Write(resb.Bytes())
			resb = tb
		}
	}

	r := strings.TrimRight(resb.String(), " ")
	if len(r) > 0 {
		return r
	}
	return "Zero"
}

func covert3Num(num int) string {
	v1 := []string{
		"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
	}
	v2 := []string{
		"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
	}

	res := bytes.Buffer{}
	a, b, c := num/100, num%100, num%10
	if a > 0 {
		res.WriteString(v1[a])
		res.WriteString(" Hundred")
	}
	if b == 0 {
		return res.String()
	}
	if a > 0 {
		res.WriteString(" ")
	}
	if b < 20 {
		res.WriteString(v1[b])
	} else {
		res.WriteString(v2[b/10])
		if c > 0 {
			res.WriteString(" ")
			res.WriteString(v1[c])
		}
	}
	return res.String()
}
