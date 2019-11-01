package main

import "fmt"

func judge(n int) bool {
	for n%10 == 0 {
		n /= 10
	}
	return n == 1
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	for !judge(n) && n != 1 {
		if _, has := m[n]; has {
			return false
		}
		m[n] = true
		tot := 0
		for n != 0 {
			tot += (n % 10) * (n % 10)
			n /= 10
		}
		n = tot
	}
	return true
}

func main() {
	fmt.Println(isHappy(1024))
}
