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

func TestIsValid(t *testing.T) {
	cases := []string{"()", "()[]{}", "(]", "([)]", "{[]}", "", "{]", "){", "(("}
	except := []bool{true, true, false, false, true, true, false, false, false}

	for index, v := range cases {
		if except[index] != IsValid(v) {
			t.Errorf("index:%d, want:%t, but got:%t", index, except[index], IsValid(v))
			return
		}
	}
}

func TestMergeTwoLists(t *testing.T) {
	type testData struct {
		l1 *ListNode
		l2 *ListNode
	}

	cases := []*testData{
		{genListNode([]int{1, 2, 4}), genListNode([]int{1, 3, 4})},
	}

	except := []*ListNode{
		genListNode([]int{1, 1, 2, 3, 4, 4}),
	}

	for index, testData := range cases {
		result := MergeTwoLists(testData.l1, testData.l2)
		if !equalListNode(except[index], result) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, stringListNode(except[index]), stringListNode(result))
			return
		}
	}
}

func TestMergeArray(t *testing.T) {
	type testData struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}

	cases := []*testData{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		{[]int{2, 5, 6, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		{[]int{0, 0, 0}, 0, []int{2, 5, 6}, 3},
	}

	except := [][]int{
		{1, 2, 2, 3, 5, 6},
		{2, 2, 5, 5, 6, 6},
		{2, 5, 6},
	}

	for index, testData := range cases {
		MergeArray(testData.nums1, testData.m, testData.nums2, testData.n)
		if !tools.EqualIntSlice(except[index], testData.nums1) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], testData.nums1)
			return
		}
	}
}
