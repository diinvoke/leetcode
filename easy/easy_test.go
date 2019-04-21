package easy

import (
	"testing"

	"github.com/diinvoke/leetcode/tools"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 13
	except := []int{0, 2}

	result := TwoSum(nums, target)

	equal := tools.EqualIntSlice(except, result)
	if !equal {
		t.Errorf("want:%+v, but got:%+v", except, result)
	}
}

func TestReverse(t *testing.T) {
	num := []int{123, -123, 120, 1000, -342, 901000, 1534236469}
	except := []int{321, -321, 21, 1, -243, 109, 0}

	for index, val := range num {
		if except[index] != Reverse(val) {
			t.Errorf("index:%d, want got:%d, but got:%d", index, except[index], Reverse(val))
			return
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	num := []int{121, 12, 21, -121, 1221, -1221, 10}
	except := []bool{true, false, false, false, true, false, false}

	for index, v := range num {
		if except[index] != IsPalindrome(v) {
			t.Errorf("index:%d, want:%t, but got %t", index, except[index], IsPalindrome(v))
			return
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	cases := [][]string{{"flower", "flow", "flight"}, {"dog", "racecar", "car"}}
	except := []string{"fl", ""}

	for index, v := range cases {
		if except[index] != LongestCommonPrefix(v) {
			t.Errorf("index:%d, want:%s, but got:%s", index, except[index], LongestCommonPrefix(v))
			return
		}
	}
}
