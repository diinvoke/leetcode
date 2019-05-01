package medium

/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number/

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

Note:
Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

var phoneLetter = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var letterDicts [][]string
	for _, digit := range digits {
		letterDicts = append(letterDicts, phoneLetter[string(digit)])
	}

	var combinations []string
	for _, letterDict := range letterDicts {
		loopCombinations := combinations[:]
		combinations = nil

		for _, letter := range letterDict {
			if len(loopCombinations) == 0 {
				combinations = append(combinations, letter)
				continue
			}

			for _, rs := range loopCombinations {
				combinations = append(combinations, rs+letter)
			}
		}
	}

	return combinations
}
