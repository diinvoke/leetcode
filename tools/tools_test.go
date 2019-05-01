package tools

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := [][]int{{1, 2, 3, 45, 6}, {4, 7, 1, 2, 9, 3, 4, 7}, {4, 4, 4, 4, 4, 1, 1, 6}}
	except := [][]int{{1, 2, 3, 6, 45}, {1, 2, 3, 4, 4, 7, 7, 9}, {1, 1, 4, 4, 4, 4, 4, 6}}

	for index, num := range nums {
		if !EqualIntSlice(except[index], QuickSort(num)) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, except[index], QuickSort(num))
			return
		}
	}
}

func TestEqualStrSlice(t *testing.T) {
	type testData struct {
		s1      []string
		s2      []string
		orderly bool
	}
	cases := []testData{
		{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, true},
		{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, []string{"cf", "ad", "ae", "af", "bd", "be", "bf", "cd", "ce"}, false},
	}

	for index, data := range cases {
		if !EqualStrSlice(data.s1, data.s2, data.orderly) {
			t.Errorf("index:%d, want:%+v, but got:%+v", index, data.s2, data.s1)
			return
		}
	}
}
