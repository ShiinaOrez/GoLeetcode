package main

import "fmt"

func maxProfit(prices []int) int {
    max := func(x, y int) int {
        if x>y {
            return x
        }
        return y
    }
    days := len(prices)
    if days == 0 {
        return 0
    }
    dp := make([]int, days)
    dp[0] = 0
    lowest := prices[0]
    for i:=1; i<days; i++ {
        dp[i] = max(dp[i-1], prices[i]-lowest)
        if prices[i] < lowest {
            lowest = prices[i]
        }
    }
    return dp[days-1]
}

func main() {
    fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
