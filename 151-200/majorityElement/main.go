package main

import "fmt"

func majorityElement(nums []int) int {
	length := len(nums)
	m := make(map[int]int)
	for _, i := range nums {
		m[i] += 1
		if m[i] > length/2 {
			return i
		}
	}
	return nums[0]
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
