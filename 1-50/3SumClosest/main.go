package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	min := 0
	num := 0
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		j := i + 1
		k := length - 1
		for j < k {
			d := target - (nums[i] + nums[j] + nums[k])
			if min > int(math.Abs(float64(d))) || min == 0 {
				min = int(math.Abs(float64(d)))
				num = target - d
			}
			if d == 0 {
				return target
			} else if d < 0 {
				k -= 1
			} else {
				j += 1
			}
		}
	}
	return num
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
