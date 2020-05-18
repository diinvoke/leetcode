package easy

/*
118. Pascal's Triangle

Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func GeneratePascalTriangle(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][len(ret[i])-1] = 1
		if len(ret[i]) <= 2 {
			continue
		}

		start := 0
		for idx := 1; idx < len(ret[i])-1 && start < len(ret[i-1]); idx++ {
			ret[i][idx] = ret[i-1][start] + ret[i-1][start+1]
			start += 1
		}
	}

	return ret
}
