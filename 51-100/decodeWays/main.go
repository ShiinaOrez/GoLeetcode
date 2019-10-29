package main

import (
	"fmt"
)

func numDecodings(s string) int {
	if s == "0" || s == "" {
		return 0
	}
	var count int
	rec(s, &count)
	return count
}

func rec(s string, c *int) {
	if len(s) <= 1 && s != "0" {
		*c += 1
		return
	} else if s == "0" {
		return
	}
	for i := 1; i <= 2; i++ {
		if i <= len(s) {
			v := atoi(s[:i])
			if v >= 1 && v <= 26 {
				rec(s[i:], c)
			} else {
				break
			}
		}
	}
	return
}

func atoi(s string) (count int) {
	m := len(s)
	re := 1
	for i := m - 1; i >= 0; i-- {
		count += re * (int(s[i]) - 48)
		re *= 10
	}
	return
}

func main() {
	fmt.Println(numDecodings("798326493284732"))
}
