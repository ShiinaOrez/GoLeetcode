package main

import "fmt"

func getRow(rowIndex int) []int {
	res := []int{1}
	for i := 2; i <= rowIndex+1; i++ {
		buf := make([]int, len(res))
		copy(buf, res)
		for j := 1; j < i-1; j++ {
			res[j] = buf[j-1] + buf[j]
		}
		res = append(res, 1)
	}
	return res
}

func main() {
	fmt.Println(getRow(6))
}
