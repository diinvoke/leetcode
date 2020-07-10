package medium

import "fmt"

/*
74. Search a 2D Matrix https://leetcode.com/problems/search-a-2d-matrix/
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
Example 1:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
Example 2:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
*/

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	rowMax := make([]int, m)
	for i := 0; i < m; i++ {
		rowMax[i] = matrix[i][n-1]
	}
	rowIdx := binarySearch(rowMax, target, 0, m-1, true)
	fmt.Println(rowIdx)
	if rowIdx == -1 {
		return false
	}

	cloArr := make([]int, n)
	for i := 0; i < n; i++ {
		cloArr[i] = matrix[rowIdx][i]
	}

	idx := binarySearch(cloArr, target, 0, n-1, false)
	return idx != -1
}

func binarySearch(array []int, target, left, right int, nearby bool) int {
	fmt.Println(left, right)
	if left > right {
		if nearby {
			return left
		}
		return -1
	}

	mid := (right + left) / 2
	if array[mid] == target {
		return mid
	}
	if left == right {
		if nearby {
			if array[mid] < target {
				if len(array)-1 == left {
					return left
				}
				return left + 1
			}
			if array[mid] > target {
				if right == 0 {
					return 0
				}
				if array[mid-1] > target {
					return right - 1
				}
				return right
			}
		}
		return -1
	}

	if array[mid] < target {
		return binarySearch(array, target, mid+1, right, nearby)
	}
	return binarySearch(array, target, left, mid-1, nearby)

}

func searchMatrix2(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	i := 0
	j := len(matrix[0]) - 1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			i += 1
			continue
		}
		if matrix[i][j] > target {
			j -= 1
		}
	}
	return false

}
