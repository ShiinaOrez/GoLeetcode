package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	length := len(nums)
	var i int
	var res []int
	var minElement, index int
	for i = length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			minElement = nums[i]
			for j := i + 1; j < length; j++ {
				if nums[j] > nums[i] {
					if minElement == nums[i] || minElement > nums[j] {
						minElement = nums[j]
						index = j
					}
				}
			}
			if minElement != nums[i] {
				nums[i], nums[index] = nums[index], nums[i]
			}
			lastSlice := nums[i+1:]
			sort.Ints(lastSlice)
			res = nums[:i+1]
			for _, r := range lastSlice {
				res = append(res, r)
			}
			break
		}
	}
	if i == -1 {
		sort.Ints(nums)
	}
	nums = res
}

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
