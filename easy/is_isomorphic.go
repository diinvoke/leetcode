package easy

/*
205. Isomorphic Strings https://leetcode.com/problems/isomorphic-strings/description/
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true
Note:
You may assume both s and t have the same length.
*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sModel := replaceModel(s)
	tModel := replaceModel(t)

	for i, m := range sModel {
		if tModel[i] != m {
			return false
		}
	}

	return true
}

func replaceModel(s string) []int {
	b := []byte(s)
	set := make(map[byte]int, len(b))
	for _, bb := range b {
		set[bb] += 1
	}

	model := make([]int, len(b), len(b))
	modelCnt := make(map[byte]int, len(b))
	c := 1
	for i, bb := range b {
		if cnt := set[bb]; cnt == 1 {
			model[i] = 0
			continue
		}
		if mc := modelCnt[bb]; mc > 0 {
			model[i] = mc
			continue
		}
		modelCnt[bb] = c
		model[i] = c
		c += 1
	}

	return model
}

func isIsomorphic2(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)
	s2t := make(map[byte]byte, len(s)/2)
	t2s := make(map[byte]byte, len(t)/2)
	for i, sr := range sb {
		tr := tb[i]
		if trr, ok := s2t[sr]; ok {
			if trr != tr {
				return false
			}
		} else {
			s2t[sr] = tr
		}

		if srr, ok := t2s[tr]; ok {
			if srr != sr {
				return false
			}
		} else {
			t2s[tr] = sr
		}
	}

	return true
}
