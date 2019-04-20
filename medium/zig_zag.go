package medium

import (
	"bytes"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/

func ZigZagConvert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= 1 {
		return s
	}

	sLen := len(s)

	gap := 2 * (numRows - 1)
	var buf bytes.Buffer
	for line := 0; line < numRows; line += 1 {
		firstIndex := line
		secondIndex := line + 2*(numRows-line-1)
		thirdIndex := line + gap

		firstGap := secondIndex - firstIndex
		secondGap := thirdIndex - secondIndex

		if firstGap == 0 && secondGap != 0 {
			firstGap = secondGap
		}

		if secondGap == 0 && firstGap != 0 {
			secondGap = firstGap
		}

		if firstGap == 0 && secondGap == 0 {
			continue
		}

		if firstIndex >= sLen {
			continue
		}

		index := firstIndex
		buf.WriteByte(s[firstIndex])
		for cnt := 0; index < sLen-1; cnt += 1 {
			if cnt%2 == 0 {
				index += firstGap
			} else {
				index += secondGap
			}

			if index >= sLen {
				break
			}
			//
			buf.WriteByte(s[index])
		}
	}

	return buf.String()

}
