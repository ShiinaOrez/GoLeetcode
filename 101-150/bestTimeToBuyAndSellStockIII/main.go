package main

import "fmt"

func maxProfit(prices []int) int {
	max := func(origin, part_1, part_2 int) int {
		if part_1+part_2 > origin {
			return part_1 + part_2
		}
		return origin
	}
	length := len(prices)
	profit := 0
	for i := 1; i < length; i++ {
		profit_1 := maxProfitInRange(prices[:i])
		profit_2 := maxProfitInRange(prices[i-1:])
		profit = max(profit, profit_1, profit_2)
	}
	return profit
}

func maxProfitInRange(prices []int) int {
	max := func(x, y int) int {
		if x > y {
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
	for i := 1; i < days; i++ {
		dp[i] = max(dp[i-1], prices[i]-lowest)
		if prices[i] < lowest {
			lowest = prices[i]
		}
	}
	return dp[days-1]
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
