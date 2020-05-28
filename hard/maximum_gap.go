package hard

import "math"

/*
164. Maximum Gap https://leetcode.com/problems/maximum-gap/description/

Given an unsorted array, find the maximum difference between the successive elements in its sorted form.

Return 0 if the array contains less than 2 elements.

Example 1:

Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.
Example 2:

Input: [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.
*/

func maximumGap(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	ma, mi := math.MinInt32, math.MaxInt32
	for _, num := range nums {
		ma = max(ma, num)
		mi = min(mi, num)
	}
	bucketCap := (ma-mi)/len(nums) + 1
	bucketCnt := (ma-mi)/bucketCap + 1

	maxArray := initArray(bucketCnt, math.MinInt32)
	minArray := initArray(bucketCnt, math.MaxInt32)

	for _, num := range nums {
		idx := (num - mi) / bucketCap
		maxArray[idx] = max(maxArray[idx], num)
		minArray[idx] = min(minArray[idx], num)
	}

	res, pre := 0, 0
	for i := 1; i < bucketCnt; i++ {
		if maxArray[i] == math.MinInt32 || minArray[i] == math.MaxInt32 {
			continue
		}
		res = max(res, minArray[i]-maxArray[pre])
		pre = i
	}
	return res
}

func initArray(cnt, defaultVal int) []int {
	array := make([]int, cnt)
	for i := 0; i < len(array); i++ {
		array[i] = defaultVal
	}
	return array
}
