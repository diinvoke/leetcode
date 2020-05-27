package hard

/*
128. Longest Consecutive Sequence https://leetcode.com/problems/longest-consecutive-sequence/

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:

Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
*/

func longestConsecutive(nums []int) int {
	res := 0
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	for _, num := range nums {
		_, ok := set[num]
		if !ok {
			continue
		}

		delete(set, num)
		pre := num - 1
		for has(pre, set) {
			delete(set, pre)
			pre -= 1
		}

		next := num + 1
		for has(next, set) {
			delete(set, next)
			next += 1
		}

		res = max(res, next-pre-1)
	}
	return res
}

func has(key int, set map[int]struct{}) bool {
	_, ok := set[key]
	return ok

}
