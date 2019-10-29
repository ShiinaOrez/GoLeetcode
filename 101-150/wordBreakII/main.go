package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	length := len(s)
	dp := make([]bool, length+1)
	dp[0] = true
	dict := make(map[string]bool)
	lines := make(map[int][][]string)
	lines[0] = [][]string{[]string{""}}

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
	if !dp[length] {
		return []string{}
	}

	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			sub := s[j : i+1]
			if _, ok := dict[sub]; ok && dp[j] {
				dp[i+1] = true
				now, k := lines[i+1]
				if !k {
					now = [][]string{}
				}
				v, _ := lines[j]
				for _, line := range v {
					newLine := make([]string, len(line)+1)
					copy(newLine, line)
					newLine[len(line)] = sub
					now = append(now, newLine)
				}
				lines[i+1] = now
			}
		}
	}
	res := []string{}
	v, _ := lines[length]
	for _, line := range v {
		res = append(res, strings.Join(line[1:], " "))
	}
	return res
}

func main() {
	fmt.Println(wordBreak("aaaaaaa", []string{"a", "aa", "aaaa"}))
}
