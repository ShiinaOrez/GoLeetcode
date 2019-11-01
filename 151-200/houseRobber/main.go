package main

import "fmt"

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[length]
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
