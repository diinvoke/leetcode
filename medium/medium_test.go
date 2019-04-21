package medium

import "testing"

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
