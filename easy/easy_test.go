package easy

import (
	"testing"

	"github.com/diinvoke/leetcode/tools"
	. "github.com/smartystreets/goconvey/convey"
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

func TestRemoveDuplicates(t *testing.T) {
	cases := [][]int{
		{1, 1, 1, 2, 2},
		{1, 1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}

	except := []int{2, 2, 5}

	for index, nums := range cases {
		result := RemoveDuplicates(nums)
		if result != except[index] {
			t.Errorf("index:%d, want:%d, but got:%d", index, except[index], result)
			return
		}
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates([]int{1, 1, 1, 2, 2})
	}
}

func BenchmarkRemoveDuplicates2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicates1([]int{1, 1, 1, 2, 2})
	}
}

func TestRemoveElement(t *testing.T) {
	type testData struct {
		nums []int
		val  int
	}

	cases := []*testData{
		{[]int{3, 2, 2, 3}, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
	}

	except := []int{2, 5}

	for index, data := range cases {
		result := removeElement(data.nums, data.val)
		if result != except[index] {
			t.Errorf("index:%d, want:%d, but got:%d", index, except[index], result)
			return
		}
	}
}

func TestStrStr(t *testing.T) {
	type testData struct {
		haystack string
		needle   string
	}

	cases := []testData{
		{"hello", "ll"},
		{"aaaaa", "bba"},
		{"aaaaa", "a"},
		{"hello", ""},
		{"hello", "hello"},
		{"hello", "hello0"},
		{"mississippi", "issipi"},
	}

	except := []int{2, -1, 0, 0, 0, -1, -1}

	for index, data := range cases {
		result := strStr(data.haystack, data.needle)
		if except[index] != result {
			t.Errorf("index:%d, want:%d, but got:%d", index, except[index], result)
			return
		}
	}
}

func TestHasCycle(t *testing.T) {
	cases := []*ListNode{
		genCycleList([]int{3, 2, 0, -4, 5}, 1),
		genCycleList([]int{3, 2, 0, -4}, 2),
		genCycleList([]int{1, 2}, 0),
		genCycleList([]int{1, 2}, -1),
	}

	except := []bool{true, true, true, false}
	for index, data := range cases {
		result := hasCycle(data)
		if except[index] != result {
			t.Errorf("index:%d, want:%t, but got:%t", index, except[index], result)
			return
		}
	}
}

func TestReverseOnlyLetters(t *testing.T) {
	cases := []string{
		"ab-cd",
		"a-bC-dEf-ghIj",
		"Test1ng-Leet=code-Q!",
	}

	except := []string{
		"dc-ba",
		"j-Ih-gfE-dCba",
		"Qedo1ct-eeLg=ntse-T!",
	}

	for index, S := range cases {
		result := reverseOnlyLetters(S)
		if except[index] != result {
			t.Errorf("index:%d, want:%s, but got:%s", index, except[index], result)
			return
		}
	}
}

func TestSearchInsert(t *testing.T) {
	type data struct {
		nums   []int
		target int
	}
	cases := []data{
		{[]int{1, 3, 5, 6}, 5},
		{[]int{1, 3, 5, 6}, 2},
		{[]int{1, 3, 5, 6}, 7},
		{[]int{1, 3, 5, 6}, 0},
	}

	except := []int{2, 1, 4, 0}

	for index, testData := range cases {
		result := searchInsert(testData.nums, testData.target)
		if except[index] != result {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], result)
			return
		}
	}
}

func TestCountAndSay(t *testing.T) {
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	except := []string{
		"1",
		"11",
		"21",
		"1211",
		"111221",
		"312211",
		"13112221",
		"1113213211",
		"31131211131221",
		"13211311123113112211",
	}

	for index, n := range cases {
		result := countAndSayLoop(n)
		if except[index] != result {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], result)
			return
		}
	}
}

func TestRotateArray(t *testing.T) {
	type testData struct {
		k     int
		array []int
	}

	cases := []*testData{
		{3, []int{1, 2, 3, 4, 5, 6, 7}},
		{2, []int{-1, -100, 3, 99}},
	}

	except := [][]int{
		{5, 6, 7, 1, 2, 3, 4},
		{3, 99, -1, -100},
	}

	for i, testCase := range cases {
		RotateArray2(testCase.array, testCase.k)
		if !tools.EqualIntSlice(testCase.array, except[i]) {
			t.Errorf("index:%d, want:%+v, but got:%+v", i, except[i], testCase.array)
			return
		}
	}
}

func TestGetHint(t *testing.T) {
	cases := [][]string{
		{"1807", "7810"},
		{"1123", "0111"},
		{"11", "11"},
		{"1122", "2211"},
		{"1122", "1222"},
	}
	except := []string{
		"1A3B", "1A1B", "2A0B", "0A4B", "3A0B",
	}

	for i, testData := range cases {
		if except[i] != GetHint2(testData[0], testData[1]) {
			t.Errorf("index:%d, want got:%s, but got:%s", i, except[i], GetHint2(testData[0], testData[1]))
			//return
		}
	}
}

func TestGeneratePascalTriangle(t *testing.T) {
	cases := []int{7}
	for _, testData := range cases {
		val := GeneratePascalTriangle(testData)
		t.Logf("%+v", val)
	}
}

func TestContainsNearbyDuplicate(t *testing.T) {
	type testData struct {
		nums []int
		k    int
	}
	cases := []testData{
		{[]int{1, 2, 3, 1}, 3},
		{[]int{1, 0, 1, 1}, 1},
		{[]int{1, 2, 3, 1, 2, 3}, 2},
	}
	except := []bool{
		true, true, false,
	}
	for i, data := range cases {
		ret := ContainsNearbyDuplicate(data.nums, data.k)
		if ret != except[i] {
			t.Errorf("index:%d, want:%t but got:%t", i, except[i], ret)
			return
		}
	}
}

func TestLengthOfLastWord(t *testing.T) {
	cases := []string{
		"Hello World",
	}
	except := []int{
		5,
	}
	for i, s := range cases {
		ret := lengthOfLastWord(s)
		if except[i] != ret {
			t.Errorf("index:%d, want:%d but got:%d", i, except[i], ret)
			return
		}
	}
}

func TestFirstUniqChar(t *testing.T) {
	Convey("test", t, func() {
		cases := []string{
			"leetcode", "loveleetcode",
		}
		except := []int{0, 2}
		for i, s := range cases {
			ret := firstUniqChar(s)
			So(ret, ShouldEqual, except[i])
		}
	})
}
