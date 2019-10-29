package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var re uint32 = 1
	var ma uint32 = 2147483648
	var newNum uint32 = 0

	for i := 0; i < 32; i++ {
		if num%2 == 1 {
			newNum += ma / re
		}
		num /= 2
		re *= 2
	}
	return newNum
}

func main() {
	var v uint32 = 784731124
	fmt.Println(reverseBits(v))
}
