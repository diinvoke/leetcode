package medium

/*
54. Spiral Matrix https://leetcode.com/problems/spiral-matrix/

Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	row := len(matrix)
	col := len(matrix[0])
	leftC, rightC := 0, col-1
	aboveR, blowR := 0, row-1

	res := make([]int, row*col)
	idx := 0
	for idx < row*col {
		// 左到右
		curLeftC := leftC
		for curLeftC <= rightC {
			res[idx] = matrix[aboveR][curLeftC]
			idx += 1
			curLeftC += 1
		}
		leftC += 1
		rightC -= 1
		if idx == row*col {
			break
		}

		// 上 -> 下
		aboveR += 1
		curAboveR := aboveR
		for curAboveR <= blowR {
			res[idx] = matrix[curAboveR][curLeftC-1]
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
			res[idx] = matrix[blowR+1][curRightC]
			idx += 1
			curRightC -= 1
		}
		if idx == row*col {
			break
		}

		// 下到上
		curBlowR := blowR
		for curBlowR >= aboveR {
			res[idx] = matrix[curBlowR][leftC-1]
			idx += 1
			curBlowR -= 1
		}
		if idx == row*col {
			break
		}
	}

	return res
}
