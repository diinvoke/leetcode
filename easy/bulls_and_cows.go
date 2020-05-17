package easy

import (
	"fmt"
	"strconv"
)

/*
299. Bulls and Cows https://leetcode.com/problems/bulls-and-cows/

You are playing the following Bulls and Cows game with your friend: You write down a number and ask your friend to guess what the number is. Each time your friend makes a guess, you provide a hint that indicates how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and how many digits match the secret number but locate in the wrong position (called "cows"). Your friend will use successive guesses and hints to eventually derive the secret number.

Write a function to return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows.

Please note that both secret number and friend's guess may contain duplicate digits.

Example 1:
Input: secret = "1807", guess = "7810"

Output: "1A3B"

Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.

Example 2:

Input: secret = "1123", guess = "0111"

Output: "1A1B"

Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.

Note: You may assume that the secret number and your friend's guess only contain digits, and their lengths are always equal.
*/

func GetHint(secret string, guess string) string {
	has := make(map[uint8]int, len(secret))
	for i := 0; i < len(secret); i++ {
		has[secret[i]] = has[secret[i]] + 1
	}

	aCnt, bCnt := 0, 0
	for i := 0; i < len(guess) && i < len(secret); i++ {
		if guess[i] == secret[i] {
			has[secret[i]] -= 1
			aCnt += 1
			continue
		}
	}

	for i := 0; i < len(guess) && i < len(secret); i++ {
		if _, ok := has[guess[i]]; !ok || guess[i] == secret[i] {
			continue
		}

		if has[guess[i]] > 0 {
			bCnt += 1
			has[guess[i]] -= 1
		}
	}

	return fmt.Sprintf("%dA%dB", aCnt, bCnt)
}

func GetHint2(secret string, guess string) string {
	has := make(map[uint8]int, len(secret))
	aCnt, bCnt := 0, 0
	for i := 0; i < len(guess) && i < len(secret); i++ {
		if guess[i] == secret[i] {
			aCnt += 1
			continue
		}

		if has[secret[i]] < 0 {
			bCnt += 1
		}
		has[secret[i]] += 1
		if has[guess[i]] > 0 {
			bCnt += 1
		}
		has[guess[i]] -= 1
	}
	return strconv.Itoa(aCnt) + "A" + strconv.Itoa(bCnt) + "B"
}

func getHint(secret string, guess string) string {
	secretdigits := make([]int, 10)
	guessdigits := make([]int, 10)
	bulls := 0
	for i := 0; i < len(secret); i++ {
		s := digit(secret[i])
		g := digit(guess[i])
		if s == g {
			bulls++
			continue
		}

		secretdigits[s]++
		guessdigits[g]++
	}

	if bulls == len(secret) {
		return buildResponse(bulls, 0)
	}

	cows := 0
	for i := 0; i < 10; i++ {
		cows += min(secretdigits[i], guessdigits[i])
	}
	return buildResponse(bulls, cows)
}

func buildResponse(bulls, cows int) string {
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

func digit(b byte) int {
	return int(rune(b) - '0')
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
