package main

import "fmt"

func hammingWeight(num uint32) int {
	tot := 0
	for num > 0 {
		if num%2 == 1 {
			tot += 1
		}
		num /= 2
	}
	return tot
}

func main() {
	v := uint32(903879643)
	fmt.Println(hammingWeight(v))
}
