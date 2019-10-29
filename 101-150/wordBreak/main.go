package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	length := len(s)
	dp := make([]bool, length+1)
	dp[0] = true
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			sub := s[j : i+1]
			if _, ok := dict[sub]; ok && dp[j] {
				dp[i+1] = true
				break
			}
		}
	}
	return dp[length]
}

func main() {
	fmt.Println(wordBreak(
		"leetcode",
		[]string{
			"leet",
			"code",
		}))
}
