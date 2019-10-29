package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		var left int
		var right int
		right = 10
		left = 1
		xx := x
		for xx/10 > 0 {
			left *= 10
			xx /= 10
		}
		for left >= right {
			num1 := x / left
			num2 := x % right
			if num1 != num2 {
				return false
			} else {
				x %= left
				x /= right
				left /= 100
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(7654567))
}
