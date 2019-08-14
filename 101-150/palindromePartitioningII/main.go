package main

import (
	"fmt"
	"time"
)

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func minCut(s string) int {
    length := len(s)
    dp := make([]int, length + 1)
    for i:=0; i<=length; i++ {
    	dp[i] = i-1
    }
    for i:=0; i<length; i++ {
    	dp[i+1] = min(dp[i+1], dp[i]+1)
    	if i == length-1 {
    		break
    	}
    	left, right := i, i+1
    	for s[left] == s[right] {
    		dp[right+1] = min(dp[right+1], dp[left]+1)
    		if left-1 == -1 || right+1 == length {
    			break
    		}
    		left -= 1; right += 1
    	}
    	left, right = i, i
    	for s[left] == s[right] {
    		dp[right+1] = min(dp[right+1], dp[left]+1)
    		if left-1 == -1 || right+1 == length {
    			break
    		}
    		left -= 1; right += 1
    	}
        if dp[length] == 0 {
        	return 0
        }
    }
    fmt.Println(dp)
    return dp[length]
}

func main() {
	t1 := time.Now()
	fmt.Println(minCut("ccaacabacb"))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}