package hard

/*
编辑距离
https://leetcode.com/problems/edit-distance/

Share
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character

Example 1:
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/

func EditDistance(word1 string, word2 string) int {
	return minDistance(word1, word2)
}

func minDistance(word1 string, word2 string) int {
	distance := make([][]int, len(word1)+1)
	for i, _ := range distance {
		array := make([]int, len(word2)+1)
		distance[i] = array
	}
	for i := 0; i <= len(word1); i++ {
		distance[i][0] = i
	}
	// word1 长度为0，变化到str2 str2有多少个，就需要增加多少个
	for j := 0; j <= len(word2); j++ {
		distance[0][j] = j
	}

	// 算法下标从1开始
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				distance[i][j] = distance[i-1][j-1]
				continue
			}
			// min (add del swap)
			/*
				以变换到 word2[1...j] 为角度
				前面已经有 word1[1...i] -> word2[i....j-1] 本次添加 word2[i]
				前面已经有 word1[1...i-1] -> word2[i....j] 本次删除 word1[i]
				前面已经有 word1[1...i-1] -> word2[i....j-1] 本次替换
			*/
			distance[i][j] = minThree(distance[i][j-1]+1, distance[i-1][j]+1, distance[i-1][j-1]+1)
		}
	}

	return distance[len(word1)][len(word2)]
}

func minThree(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}

	return min
}
