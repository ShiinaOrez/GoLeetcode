package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)
	newPoint := 0
	for point := 0; point < length; point++ {
		if nums[point] == val {
			continue
		}
		if nums[newPoint] != nums[point] {
			nums[newPoint] = nums[point]
		}
		newPoint += 1
	}
	return newPoint
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 1, 3, 3}, 3))
}
