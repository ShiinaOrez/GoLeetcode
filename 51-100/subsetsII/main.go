package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	v := []int{}
	res = append(res, v)
	n := len(nums)
	right := 1
	left := 0
	length := 0
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			left = len(res) - length
		} else {
			left = 0
		}
		right = len(res)
		length = right - left
		for j := left; j < right; j++ {
			u := make([]int, len(res[j]))
			copy(u, res[j])
			u = append(u, nums[i])
			res = append(res, u)
		}
	}
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
