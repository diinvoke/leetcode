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

func TestDivide(t *testing.T) {
	type testData struct {
		dividend int
		divisor  int
	}
	cases := []testData{
		{10, 3},
		{7, -3},
		{-2147483648, 1},
		{-2147483648, -1},
	}
	except := []int{3, -2, -2147483648, 2147483647}

	for index, data := range cases {
		result := divide(data.dividend, data.divisor)
		if except[index] != result {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], result)
			return
		}
	}
}

func TestNextPermutation(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{1, 1, 5},
		{1, 2, 7, 4, 3, 1},
	}

	except := [][]int{
		{1, 3, 2},
		{1, 2, 3},
		{1, 5, 1},
		{1, 3, 1, 2, 4, 7},
	}

	for index, nums := range cases {
		nextPermutation(nums)
		if !tools.EqualIntSlice(except[index], nums) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], nums)
			return
		}
	}
}

func TestPermute(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
	}
	except := [][][]int{
		{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		},
	}

	for index, nums := range cases {
		result := permute(nums)
		if !tools.EqualDoubleSlice(except[index], result, true) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], result)
			return
		}
	}
}

func TestPermuteUnique(t *testing.T) {
	cases := [][]int{
		{1, 1, 2},
	}
	except := [][][]int{
		{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
		},
	}

	for index, nums := range cases {
		result := permuteUnique(nums)
		if !tools.EqualDoubleSlice(except[index], result, true) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], result)
			return
		}
	}
}

func TestSearchRotatedSorted(t *testing.T) {
	type data struct {
		nums   []int
		target int
	}
	cases := []data{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
	}

	except := []int{4, -1}

	for index, testData := range cases {
		result := searchRotatedSorted(testData.nums, testData.target)
		if except[index] != result {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], result)
			return
		}
	}
}

func TestSearchRange(t *testing.T) {
	type data struct {
		nums   []int
		target int
	}
	cases := []data{
		{[]int{5, 7, 7, 8, 8, 10}, 8},
		{[]int{5, 7, 7, 8, 8, 10}, 6},
		{[]int{5, 7, 7, 8, 8, 10}, 8},
		{[]int{1, 2, 2, 3, 4, 4, 4}, 4},
		{[]int{1}, 1},
		{[]int{2, 2}, 2},
	}

	except := [][]int{
		{3, 4},
		{-1, -1},
		{3, 4},
		{4, 6},
		{0, 0},
		{0, 1},
	}
	for index, testData := range cases {
		result := searchRange(testData.nums, testData.target)
		if !tools.EqualIntSlice(except[index], result) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], result)
			return
		}
	}

}

func TestCombinationSum(t *testing.T) {
	type testData struct {
		candidates []int
		target     int
	}
	cases := []testData{
		{[]int{2, 3, 6, 7}, 7},
	}

	for _, data := range cases {
		result := combinationSum(data.candidates, data.target)
		t.Logf("----->%+v", result)
	}
}

func TestCombinationSum2(t *testing.T) {
	type testData struct {
		candidates []int
		target     int
	}
	cases := []testData{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
		//{[]int{2, 5, 2, 1, 2}, 5},
	}

	for _, data := range cases {
		result := combinationSum2(data.candidates, data.target)
		t.Logf("----->%+v", result)
	}
}

func TestAddTwo(t *testing.T) {
	type testData struct {
		num1 []int
		num2 []int
	}
	cases := []testData{
		{[]int{8, 6, 3, 1}, []int{0, 2, 1, 9}},
		{[]int{8, 8, 4, 0, 1}, []int{0, 0, 6, 5, 4}},
	}

	for _, data := range cases {
		result := addTwo(data.num1, data.num2)
		t.Logf("%+v", result)
	}
}

func TestMultiply(t *testing.T) {
	type testData struct {
		num1 string
		num2 string
	}
	cases := []testData{
		{"123", "456"},
		{"2", "3"},
		{"9133", "0"},
		{"2321435435", "464565765765876"},
	}

	for index, data := range cases {
		result := multiply(data.num1, data.num2)
		result2 := multiply2(data.num1, data.num2)
		if result2 != result {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, result, result2)
			return
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	type exceptData struct {
		array  []int
		length int
	}
	cases := [][]int{
		{1, 1, 1, 2, 2, 3},
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
	}
	except := []*exceptData{
		{[]int{1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 2, 3, 3}, 7},
	}

	for i, array := range cases {
		length := RemoveDuplicates(array)
		if length != except[i].length || !tools.EqualIntSlice(except[i].array, array[0:length]) {
			t.Errorf("index:%d, len:%d, want:%+v, but got:%+v", i, length, except[i].array, array[0:length])
			return
		}
	}
}

func TestCanCompleteCircuit(t *testing.T) {
	cases := [][][]int{
		{{1, 2, 3, 4, 5}, {3, 4, 5, 1, 2}},
		{{2, 3, 4}, {3, 4, 3}},
		{{5, 1, 2, 3, 4}, {4, 4, 1, 5, 1}},
	}
	except := []int{3, -1, 4}

	for i, testData := range cases {
		val := CanCompleteCircuit(testData[0], testData[1])
		if except[i] != val {
			t.Errorf("index:%d, want %d, but got %d", i, except[i], val)
			return
		}
	}
}

func TestMajorityElement(t *testing.T) {
	cases := [][]int{
		{3, 2, 3},
		{1, 1, 1, 3, 3, 2, 2, 2},
	}
	except := [][]int{
		{3},
		{1, 2},
	}

	for i, nums := range cases {
		ret := MajorityElement(nums)
		if !tools.HasSameElement(ret, except[i]) {
			t.Errorf("index:%d want:%+v but got: %+v", i, except[i], ret)
			return
		}
	}
}

func TestHIndex(t *testing.T) {
	cases := [][]int{
		{3, 0, 6, 1, 5},
	}
	except := []int{3}

	for i, citations := range cases {
		ret := HIndex2(citations)
		if except[i] != ret {
			t.Errorf("index:%d want:%d but got:%d", i, except[i], ret)
			return
		}
	}
}

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	type tt struct {
		nums []int
		k, t int
	}

	cases := []tt{
		{[]int{1, 2, 3, 1}, 3, 0},
		{[]int{1, 0, 1, 1}, 1, 2},
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3},
	}

	except := []bool{
		true, true, false,
	}

	for i, ttt := range cases {
		ret := ContainsNearbyAlmostDuplicate(ttt.nums, ttt.k, ttt.t)
		if except[i] != ret {
			t.Errorf("index:%d, want:%t but got:%t", i, except[i], ret)
			return
		}
	}
}

func TestCanJump(t *testing.T) {
	cases := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
	}
	except := []bool{
		true, false,
	}

	for i, nums := range cases {
		ret := CanJump(nums)
		if except[i] != ret {
			t.Errorf("index:%d, want:%t but got:%t", i, except[i], ret)
			return
		}
	}
}
