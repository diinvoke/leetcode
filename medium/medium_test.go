package medium

import (
	"testing"

	"github.com/diinvoke/leetcode/tools"
)

func TestEqualListNode(t *testing.T) {
	l1 := genListNode([]int{1, 3, 4})
	l2 := genListNode([]int{1, 3, 4})

	t.Logf("l1:%s", stringListNode(l1))
	t.Logf("l2:%s", stringListNode(l2))
	if !equalListNode(l1, l2) {
		t.Errorf("equal listNode failed, should true")
		return
	}

	l2 = nil
	if equalListNode(l1, l2) {
		t.Errorf("equal listNode failed, should false")
		return
	}

	l2 = genListNode([]int{1, 3, 4, 5})
	if equalListNode(l1, l2) {
		t.Errorf("equal listNode failed, should false")
		return
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := genListNode([]int{1, 3, 4})
	l2 := genListNode([]int{1, 3, 4})
	except := genListNode([]int{2, 6, 8})
	result := AddTwoNumbers(l1, l2)
	if !equalListNode(except, result) {
		t.Errorf("want:%s, but got:%s", stringListNode(except), stringListNode(result))
		return
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	ss := []string{"abcabcbb", "bbbbb", "pwwkew", "cdd", "pwwk", "abba", "abbba", "fsfjdls", "ckilbkd", "anviaj", "uqinntq",
		"wobgrovw"}
	except := []int{3, 1, 3, 2, 2, 2, 2, 5, 5, 5, 4, 6}

	for index, s := range ss {
		if LengthOfLongestSubstring(s) != except[index] {
			t.Errorf("want:%d, but got:%d", except[index], LengthOfLongestSubstring(s))
			return
		}

		//t.Logf("num:%d", LengthOfLongestSubstring(s))
	}
}

func TestLongestPalindrome(t *testing.T) {
	ss := []string{"abcdbbfcba"}
	except := []string{"bb"}

	for index, s := range ss {
		if except[index] != LongestPalindrome(s) {
			t.Errorf("index:%d, want:%s, but got:%s", index, except[index], LongestPalindrome(s))
			return
		}
	}
}

func TestZigZagConvert(t *testing.T) {
	type data struct {
		s   string
		num int
	}
	ss := []data{
		{"Apalindromeisaword,phrase,number,orothersequenceofunitsthatcanbereadthesamewayineitherdirection,withgeneralallowancesforadjustmentstopunctuationandworddividers.",
			2}}
	except := []string{
		"Aaidoeswr,haenme,rtesqecouishtabrateaeaietedrcinwtgnrlloacsoajsmnsoucutoadodiiesplnrmiaodprs,ubroohreunefnttacneedhsmwynihrieto,iheeaalwnefrdutettpntainnwrdvdr."}

	for index, s := range ss {
		result := ZigZagConvert(s.s, s.num)
		if except[index] != result {
			t.Errorf("index:%d, want:%s, but got:%s", index, except[index], result)
			return
		}
	}
}

func TestMyAtoi(t *testing.T) {
	ss := []string{"42", "   -42", "4193 with words", "words and 987", "-91283472332",
		"3.14159", "+-2", "   +2-", "+2- ", "   +0 123", "9223372036854775808", "-9223372036854775808", "0-1"}
	except := []int{42, -42, 4193, 0, -2147483648, 3, 0, 2, 2, 0, 2147483647, -2147483648, 0}

	for index, s := range ss {
		if except[index] != MyAtoi(s) {
			t.Errorf("index:%d, want:%d, but got:%d", index, except[index], MyAtoi(s))
			return
		}
	}
}

func TestMaxArea(t *testing.T) {
	cases := [][]int{{1, 8, 6, 2, 5, 4, 8, 3, 7}}
	except := []int{49}

	for index, testCase := range cases {
		if except[index] != MaxArea(testCase) {
			t.Errorf("index:%d, want:%d, but got:%d", except[index], MaxArea(testCase))
			return
		}
	}
}

func TestThreeSum(t *testing.T) {
	cases := [][]int{{-1, 0, 1, 2, -1, -4}}
	except := [][][]int{{{-1, 0, 1}, {-1, -1, 2}}}

	for index, nums := range cases {
		if !tools.EqualDoubleSlice(except[index], ThreeSum(nums), true) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], ThreeSum(nums))
			return
		}
	}
}

func TestThreeSumClosest(t *testing.T) {
	type data struct {
		nums   []int
		target int
	}
	cases := []data{
		{[]int{-1, 2, 1, -4}, 1},
		{[]int{1, 1, 1, 0}, 100},
		{[]int{0, 2, 1, -3}, 1},
	}
	except := []int{2, 3, 0}

	for index, testCase := range cases {
		if except[index] != ThreeSumClosest(testCase.nums, testCase.target) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], ThreeSumClosest(testCase.nums, testCase.target))
			return
		}
	}
}

func TestLetterCombinations(t *testing.T) {
	digits := []string{"23"}
	except := [][]string{{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}}

	for index, digit := range digits {
		if !tools.EqualStrSlice(LetterCombinations(digit), except[index], false) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], LetterCombinations(digit))
			return
		}

		t.Logf("%+v", LetterCombinations(digit))
	}
}

func TestFourSum(t *testing.T) {
	type testData struct {
		nums   []int
		target int
	}
	cases := []*testData{
		{[]int{1, 0, -1, 0, -2, 2}, 0},
		{[]int{4, -9, -2, -2, -7, 9, 9, 5, 10, -10, 4, 5, 2, -4, -2, 4, -9, 5}, -13},
	}
	except := [][][]int{
		{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}},
		{{-10, -9, -4, 10}, {-10, -9, 2, 4}, {-9, -9, -4, 9}, {-9, -7, -2, 5}, {-9, -4, -2, 2}, {-7, -2, -2, -2}},
	}

	for index, data := range cases {
		if !tools.EqualDoubleSlice(except[index], FourSum(data.nums, data.target), true) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], FourSum(data.nums, data.target))
			return
		}
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	type testData struct {
		list *ListNode
		n    int
	}

	cases := []*testData{
		{genListNode([]int{1, 2, 3, 4, 5}), 2},
		{genListNode([]int{1, 2, 3, 4, 5}), 5},
		{genListNode([]int{1}), 1},
	}
	except := []*ListNode{
		genListNode([]int{1, 2, 3, 5}),
		genListNode([]int{2, 3, 4, 5}),
		genListNode([]int{}),
	}

	for index, data := range cases {
		result := RemoveNthFromEnd(data.list, data.n)
		if !equalListNode(except[index], result) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, stringListNode(except[index]), stringListNode(result))
			return
		}
	}
}

func TestSwapPairs(t *testing.T) {
	cases := []*ListNode{
		genListNode([]int{1, 2, 3, 4}),
		genListNode([]int{1, 2}),
		genListNode([]int{1}),
		genListNode([]int{}),
	}

	except := []*ListNode{
		genListNode([]int{2, 1, 4, 3}),
		genListNode([]int{2, 1}),
		genListNode([]int{1}),
		genListNode([]int{}),
	}

	for index, head := range cases {
		result := SwapPairs(head)
		if !equalListNode(except[index], result) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, stringListNode(except[index]), stringListNode(result))
			return
		}
	}
}
