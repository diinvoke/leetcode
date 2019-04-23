package tools

import "sort"

func EqualIntSlice(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for index, val := range s1 {
		if s2[index] != val {
			return false
		}
	}

	return true
}

func EqualDoubleSlice(s1, s2 [][]int, needSort bool) bool {
	var ss1 []int
	var ss2 []int

	for _, val := range s1 {
		ss1 = append(ss1, val...)
	}

	for _, val := range s2 {
		ss2 = append(ss2, val...)
	}

	if needSort {
		sort.Ints(ss1)
		sort.Ints(ss2)
	}

	return EqualIntSlice(ss1, ss2)
}
