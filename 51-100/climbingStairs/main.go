package main

import "fmt"

func climbStairs(n int) int {
    dp := []int{}
    dp = append(dp, 1)
    dp = append(dp, 1)
    for i:=2; i<=n; i++ {
        dp = append(dp, dp[i-1] + dp[i-2])
    }
    return dp[n]
}

func main() {
    fmt.Println(climbStairs(70))
}
