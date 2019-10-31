package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	sq := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq+1; i++ {
		if num%i == 0 {
			return true
		}
	}
	return false
}

func divisorGame(N int) bool {
	if N == 2 {
		return true
	}
	if isPrime(N) {
		return !divisorGame(N - 1)
	} else {
		for i := 2; i <= N/2; i++ {
			if N%i == 0 {
				return !divisorGame(N - (N / i))
			}
		}
	}
	return false
}

func main() {
	fmt.Println(divisorGame(1001))
}
