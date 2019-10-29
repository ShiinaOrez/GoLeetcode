package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	length := len(prices)
	prices = append(prices, 0)
	for i := 0; i < length; i++ {
		if prices[i] < prices[i+1] {
			profit += (prices[i+1] - prices[i])
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
