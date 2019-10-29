package main

import "fmt"

func romanToInt(s string) int {
	dict := map[byte]int{
		byte('I'): 1,
		byte('V'): 5,
		byte('X'): 10,
		byte('L'): 50,
		byte('C'): 100,
		byte('D'): 500,
		byte('M'): 1000,
	}
	res := 0
	for i, r := range s {
		if i > 0 {
			if dict[byte(r)] > dict[s[i-1]] {
				res -= 2 * dict[s[i-1]]
			}
		}
		res += dict[byte(r)]
	}
	return res
}

func main() {
	fmt.Println(romanToInt("III"))
}
