package main

import "fmt"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[0]
	}
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid
		} else {
			right = mid
		}
		if left == right-1 {
			return nums[right]
		}
	}
	return nums[right]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
