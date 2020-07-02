package medium

/*
59. Spiral Matrix II https://leetcode.com/problems/spiral-matrix-ii/

Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/

func generateMatrix(n int) [][]int {
	row, col := n, n
	leftC, rightC := 0, col-1
	aboveR, blowR := 0, row-1
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	total := n * n
	idx := 0
	for idx != total {
		// 左到右
		curLeftC := leftC
		for curLeftC <= rightC {
			matrix[aboveR][curLeftC] = idx + 1
			idx += 1
			curLeftC += 1
		}
		leftC += 1
		rightC -= 1
		if idx == total {
			break
		}

		// 上 -> 下
		aboveR += 1
		curAboveR := aboveR
		for curAboveR <= blowR {
			matrix[curAboveR][curLeftC-1] = idx + 1
			idx += 1
			curAboveR += 1
		}
		blowR -= 1
		if idx == row*col {
			break
		}

		// 右到左
		curRightC := rightC
		for curRightC >= leftC-1 {
			matrix[blowR+1][curRightC] = idx + 1
			idx += 1
			curRightC -= 1
		}
		if idx == row*col {
			break
		}

		// 下到上
		curBlowR := blowR
		for curBlowR >= aboveR {
			matrix[curBlowR][leftC-1] = idx + 1
			idx += 1
			curBlowR -= 1
		}
		if idx == row*col {
			break
		}
	}

	return matrix
}
