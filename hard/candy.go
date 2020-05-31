package hard

/*
135. Candy https://leetcode.com/problems/candy/description/

There are N children standing in a line. Each child is assigned a rating value.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
What is the minimum candies you must give?

Example 1:

Input: [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
Example 2:

Input: [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
             The third child gets 1 candy because it satisfies the above two conditions.
*/

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}

	for i := 0; i < len(candies)-1; i++ {
		if ratings[i] < ratings[i+1] {
			candies[i+1] = candies[i] + 1
		}
	}

	for i := len(candies) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			candies[i-1] = max(candies[i-1], candies[i]+1)
		}
	}

	total := 0
	for _, num := range candies {
		total += num
	}

	return total
}

func candy2(ratings []int) int {
	pre, cnt := 1, 0
	res := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			if cnt > 0 {
				res += cnt * (cnt + 1) / 2
				if cnt >= pre { // 第一个开始递减的前一个是否需要+1
					res += cnt - pre + 1
				}
				cnt = 0
				pre = 1 // 到最后一个递减的人了
			}

			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre += 1
			}

			res += pre
		} else {
			cnt += 1
		}
	}

	if cnt > 0 {
		res += cnt * (cnt + 1) / 2
		if cnt >= pre {
			res += cnt - pre + 1
		}
	}
	return res
}
