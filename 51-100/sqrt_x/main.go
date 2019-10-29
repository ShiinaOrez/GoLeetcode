package main

import "fmt"

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	start := 0
	end := x
	mid := 0
	for !(mid*mid <= x && (mid+1)*(mid+1) > x) {
		mid = (start + end) / 2
		if mid*mid > x {
			end = mid
		} else if mid*mid == x {
			return mid
		} else {
			start = mid
		}
	}
	return mid
}

func main() {
	fmt.Println(mySqrt(100))
}
