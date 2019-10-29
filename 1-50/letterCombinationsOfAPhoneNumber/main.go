package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	m := map[byte]string{
		byte('2'): "abc",
		byte('3'): "def",
		byte('4'): "ghi",
		byte('5'): "jkl",
		byte('6'): "mno",
		byte('7'): "pqrs",
		byte('8'): "tuv",
		byte('9'): "wxyz",
	}
	var res []string
	if len(digits) > 1 {
		for _, r := range m[digits[0]] {
			c := byte(r)
			for _, s := range letterCombinations(digits[1:]) {
				res = append(res, string(c)+s)
			}
		}
	} else {
		for _, r := range m[digits[0]] {
			res = append(res, string(r))
		}
	}
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
