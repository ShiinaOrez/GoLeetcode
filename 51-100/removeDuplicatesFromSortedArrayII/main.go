package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 0
	n := len(nums)
	if n == 0 {
		return 0
	}
	now := nums[0] - 1
	setPoint := 0
	for i := 0; i < n; i++ {
		if now != nums[i] {
			now = nums[i]
			count = 0
		}

		count += 1
		if count > 2 {
			continue
		} else {
			nums[setPoint] = nums[i]
			setPoint += 1
		}
	}
	return setPoint
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3, 3, 3}))
}
