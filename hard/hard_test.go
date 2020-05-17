package hard

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	type data struct {
		nums1 []int
		nums2 []int
	}

	testDatas := []data{{[]int{1, 3}, []int{2}}, {[]int{1, 2}, []int{3, 4}}, {[]int{1}, []int{1}}}
	except := []float64{2.0, 2.5, 1.0}

	for index, testData := range testDatas {
		if except[index] != FindMedianSortedArrays(testData.nums1, testData.nums2) {
			t.Errorf("solution 1 index:%d, want:%f, but got:%f", index, except[index], FindMedianSortedArrays(testData.nums1, testData.nums2))
			return
		}
	}

	for index, testData := range testDatas {
		if except[index] != FindMedianSortedArrays2(testData.nums1, testData.nums2) {
			t.Errorf("solution 2 index:%d, want:%f, but got:%f", index, except[index], FindMedianSortedArrays(testData.nums1, testData.nums2))
			return
		}
	}

	for index, testData := range testDatas {
		if except[index] != FindMedianSortedArrays3(testData.nums1, testData.nums2) {
			t.Errorf("solution 3 index:%d, want:%f, but got:%f", index, except[index], FindMedianSortedArrays(testData.nums1, testData.nums2))
			return
		}
	}
}

func TestMergeKLists(t *testing.T) {
	cases := [][]*ListNode{
		{
			genListNode([]int{1, 4, 5}),
			genListNode([]int{1, 3, 4}),
			genListNode([]int{2, 6}),
		},
	}

	except := []*ListNode{
		genListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
	}

	for index, testData := range cases {
		result := MergeKLists(testData)
		if !equalListNode(except[index], result) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, stringListNode(except[index]), stringListNode(result))
			return
		}
	}
}

func TestReverseList(t *testing.T) {
	cases := []*ListNode{
		genListNode([]int{1, 2, 3, 4}),
		genListNode([]int{4, 3, 1, 2}),
		genListNode([]int{4, 5}),
		genListNode([]int{3, 6, 1, 7, 2}),
		genListNode([]int{1}),
		genListNode([]int{}),
	}

	except := []*ListNode{
		genListNode([]int{4, 3, 2, 1}),
		genListNode([]int{2, 1, 3, 4}),
		genListNode([]int{5, 4}),
		genListNode([]int{2, 7, 1, 6, 3}),
		genListNode([]int{1}),
		genListNode([]int{}),
	}

	for index, data := range cases {
		result := reverseList(data)
		if !equalListNode(except[index], result) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, stringListNode(except[index]), stringListNode(result))
			return
		}
	}
}

func TestReverseKGroup(t *testing.T) {
	type testData struct {
		head *ListNode
		k    int
	}

	cases := []testData{
		{genListNode([]int{1, 2, 3, 4, 5}), 2},
		{genListNode([]int{1, 2, 3, 4, 5}), 3},
		{genListNode([]int{1, 2, 3, 4, 5}), 1},
		{genListNode([]int{1, 2, 3, 4, 5}), 0},
		{genListNode([]int{1, 2, 3, 4, 5}), 4},
	}

	except := []*ListNode{
		genListNode([]int{2, 1, 4, 3, 5}),
		genListNode([]int{3, 2, 1, 4, 5}),
		genListNode([]int{1, 2, 3, 4, 5}),
		genListNode([]int{1, 2, 3, 4, 5}),
		genListNode([]int{4, 3, 2, 1, 5}),
	}

	for index, data := range cases {
		result := reverseKGroup(data.head, data.k)
		if !equalListNode(except[index], result) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, stringListNode(except[index]), stringListNode(result))
			return
		}
	}
}

func TestLongestValidParentheses(t *testing.T) {
	cases := []string{
		"()(()",
		"(()",
		")()())",
		")))))))",
		"()(())",
	}

	except := []int{
		2, 2, 4, 0, 6,
	}

	for index, s := range cases {
		result := longestValidParentheses(s)
		if result != except[index] {
			t.Errorf("index:%d, want:%d, but got:%d", index, except[index], result)
			return
		}
	}
}

func TestEditDistance(t *testing.T) {
	testData := [][]string{
		{"", "a"},
		{"horse", "ros"},
		{"kitten", "sitting"},
		{"intention", "execution"},
	}

	except := []int{1, 3, 3, 5}
	for index, data := range testData {
		minDis := EditDistance(data[0], data[1])
		if minDis != except[index] {
			t.Errorf("index:%d, want:%d, but got:%d", index, except[index], minDis)
			return
		}
	}
}

func TestFirstMissingPositive(t *testing.T) {
	cases := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
	}

	except := []int{3, 2, 1}
	for i, array := range cases {
		if FirstMissingPositive2(array) != except[i] {
			t.Errorf("index:%d want:%d, but got:%d", i, except[i], FirstMissingPositive(array))
			return
		}
	}
}
