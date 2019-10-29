package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" || s == "." {
		return false
	}
	hasNum := false
	hasMark := false
	hasE := false
	hasDoc := false
	for i, r := range s {
		b := byte(r)
		if b == '+' || b == '-' {
			if hasMark {
				if hasE && s[i-1] == 'e' {
					if i == len(s)-1 {
						return false
					}
					continue
				}
				return false
			} else {
				if hasNum || hasDoc {
					if i > 0 && i < len(s)-1 {
						if s[i-1] == 'e' {
							continue
						}
					}
					return false
				}
				hasMark = true
			}
		} else if int(b)-48 <= 9 && int(b)-48 >= 0 {
			hasNum = true
		} else if b == 'e' {
			if !hasNum {
				return false
			}
			if i == len(s)-1 {
				return false
			}
			if hasE {
				return false
			} else {
				hasE = true
			}
		} else if b == '.' {
			if hasE {
				return false
			}
			if hasDoc {
				return false
			} else {
				hasDoc = true
			}
		} else {
			return false
		}
	}
	return hasNum
}

func main() {
	fmt.Println(isNumber("95a54e53"))
}
