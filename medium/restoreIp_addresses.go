package medium

import (
	"strconv"
)

/*
93. Restore IP Addresses https://leetcode.com/problems/restore-ip-addresses/
Given a string containing only digits, restore it by returning all possible valid IP address combinations.

A valid IP address consists of exactly four integers (each integer is between 0 and 255) separated by single points.

Example:

Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
*/

func restoreIpAddresses(s string) []string {
	res := make([]string, 0, 1)
	recursiveIp(s, 4, "", &res)
	return res
}

func recursiveIp(s string, k int, curRes string, res *[]string) {
	if k == 0 {
		if len(s) == 0 {
			*res = append(*res, curRes)
		}
		return
	}
	for i := 1; i <= 3; i++ { // 每段最多三个字符
		if len(s) < i || !validIp(s[:i]) {
			continue
		}

		if k == 1 { // 最后一个组合
			recursiveIp(s[i:], k-1, curRes+s[:i], res)
			continue
		}
		recursiveIp(s[i:], k-1, curRes+s[:i]+".", res)
	}
}

func validIp(s string) bool {
	if len(s) == 0 || len(s) > 3 || (len(s) > 1 && s[0] == '0') {
		return false
	}
	num, e := strconv.ParseInt(s, 10, 64)
	return e == nil && (num <= 255 && num >= 0)
}
