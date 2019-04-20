package tools

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
