package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
func calculateMinimumHP(dungeon [][]int) int {
    m, n := len(dungeon), len(dungeon[0])
    dp := make([][]int, m+1)
    for i:=0; i<=m; i++ {
        dp[i] = make([]int, n+1)
        for j:=0; j<=n; j++ {
            dp[i][j] = INT_MAX
        }
    }
    dp[m][n-1] = 1
    dp[m-1][n] = 1
    for i:=m-1; i>=0; i-- {
        for j:=n-1; j>=0; j-- {
            dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
        }
    }
    return dp[0][0]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func main() {
    fmt.Println()
}
