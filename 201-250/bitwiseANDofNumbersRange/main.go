package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	offset := 1
	for ; m != n; offset *= 2 {
		m /= 2
		n /= 2
	}
	return m * offset
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}
