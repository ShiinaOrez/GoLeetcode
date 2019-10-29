package main

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length

	copy(nums, append(nums[length-k:], nums[:length-k]...))
	return
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(testSlice, 4)
	fmt.Println(testSlice)
}
