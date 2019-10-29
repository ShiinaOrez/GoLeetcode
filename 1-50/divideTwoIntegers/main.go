package main

import "fmt"

func divide(dividend int, divisor int) int {
	limit := 1<<31 - 1
	flag := false
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		flag = true
	}

	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if 0-dividend > limit {
			return limit
		}
		return 0 - dividend
	}

	var res int64
	res = 0
	if divisor < 0 {
		divisor = 0 - divisor
	}
	if dividend < 0 {
		dividend = 0 - dividend
	}
	for dividend > 0 {
		dividend -= divisor
		if dividend < 0 {
			break
		}
		res += 1
		if res > int64(limit) {
			return limit
		}
	}
	if flag {
		return int(0 - res)
	}
	return int(res)
}

func main() {
	fmt.Println(divide(-2147483648, -3))
}
