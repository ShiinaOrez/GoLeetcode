package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	rows := len(triangle)
	buf := make([]int, rows)
	newBuf := make([]int, rows)
	buf[0] = triangle[0][0]
	for i := 1; i < rows; i++ {
		newBuf[0] = buf[0] + triangle[i][0]
		for j := 1; j < i; j++ {
			newBuf[j] = min(buf[j], buf[j-1]) + triangle[i][j]
		}
		newBuf[i] = buf[i-1] + triangle[i][i]
		copy(buf, newBuf)
	}
	ans := buf[0]
	for i := 0; i < rows; i++ {
		if ans > buf[i] {
			ans = buf[i]
		}
	}
	return ans
}

func main() {
	fmt.Println()
}
