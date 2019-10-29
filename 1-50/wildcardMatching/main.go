package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)
	// dp means: dp[i][j]: p[:i] is match s[:j]
	dp := make([][]bool, lenP+1)
	for i, _ := range dp {
		dp[i] = make([]bool, lenS+1)
	}

	dp[0][0] = true
	for j := 1; j < lenS; j++ {
		dp[0][j] = false
	}

	for i := 1; i <= lenP; i++ {
		if p[i-1] == '*' {
			flag := false
			for j := 0; j <= lenS; j++ {
				if dp[i-1][j] {
					flag = true
				}
				if flag {
					dp[i][j] = true
				}
			}
		} else if p[i-1] == '?' {
			for j := 0; j < lenS; j++ {
				if dp[i-1][j] {
					dp[i][j+1] = true
				}
			}
		} else {
			for j := 0; j < lenS; j++ {
				if dp[i-1][j] && s[j] == p[i-1] {
					dp[i][j+1] = true
				}
			}
		}
	}
	return dp[lenP][lenS]
}

func main() {
	fmt.Println(isMatch("aaabbc", "a*b*c"))
}
