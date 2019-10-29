package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 1; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for j := 1; j < i-1; j++ {
			row[j] = res[i-2][j-1] + res[i-2][j]
		}
		res = append(res, row)
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
