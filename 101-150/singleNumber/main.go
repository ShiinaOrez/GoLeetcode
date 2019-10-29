package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 4, 4, 1}))
}
