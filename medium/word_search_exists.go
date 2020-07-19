package medium

/*
79. Word Search https://leetcode.com/problems/word-search/

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.


Constraints:

board and word consists only of lowercase and uppercase English letters.
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
*/

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if searchExists(board, i, j, 0, visited, word, m, n) {
				return true
			}
		}
	}
	return false
}

func searchExists(board [][]byte, i, j int, level int, visited [][]bool, word string, m, n int) bool {
	if level == len(word) {
		return true
	}
	if i < 0 || j < 0 || j >= n || i >= m || visited[i][j] || board[i][j] != word[level] {
		return false
	}

	visited[i][j] = true
	res := searchExists(board, i-1, j, level+1, visited, word, m, n) ||
		searchExists(board, i+1, j, level+1, visited, word, m, n) ||
		searchExists(board, i, j-1, level+1, visited, word, m, n) ||
		searchExists(board, i, j+1, level+1, visited, word, m, n)
	visited[i][j] = false
	return res
}
