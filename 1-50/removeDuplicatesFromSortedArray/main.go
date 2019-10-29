package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	newPoint := 0
	for point := 0; point < length; point++ {
		if nums[newPoint] != nums[point] {
			nums[newPoint] = nums[point]
			newPoint += 1
		} else {
			newPoint += 1
		}
		if point == length-1 {
			break
		}
		for nums[point] == nums[point+1] {
			point += 1
			if point == length-1 {
				break
			}
		}
	}
	return newPoint
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3}))
}
