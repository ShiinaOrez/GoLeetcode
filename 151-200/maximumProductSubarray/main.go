package main

import "fmt"

func maxProduct(nums []int) int {
	length := len(nums)
	a, b := 1, 1
	maxNum := ^int(^uint(0) >> 1)
	for i := 0; i < length; i++ {
		tmp := b
		b = max(nums[i], max(a*nums[i], b*nums[i]))
		a = min(nums[i], min(a*nums[i], tmp*nums[i]))
		maxNum = max(maxNum, b)
	}
	return maxNum
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxProduct([]int{-3, 0, 1, -2}))
}
