package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	newBytes := []byte{}
	for i, _ := range s {
		if match(s[i]) {
			newBytes = append(newBytes, s[i])
		}
	}
	s = strings.ToUpper(string(newBytes))
	if s == "" {
		return true
	}
	length := len(s)
	for i := 0; i < length/2+1; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func match(b byte) bool {
	if b <= '9' && b >= '0' {
		return true
	} else if b <= 'z' && b >= 'a' {
		return true
	} else if b <= 'Z' && b >= 'A' {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
