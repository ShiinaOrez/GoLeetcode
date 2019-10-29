package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	hasMap := make(map[string]int)
	ans := []string{}
	if len(s) < 10 {
		return ans
	} else {
		start := 0
		end := 10
		length := len(s)
		for end <= length {
			subStr := s[start:end]
			if _, ok := hasMap[subStr]; !ok {
				hasMap[subStr] = 1
			} else {
				hasMap[subStr] += 1
			}
			start += 1
			end += 1
		}
	}
	for k, v := range hasMap {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
