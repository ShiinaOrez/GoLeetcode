package main

import (
	"fmt"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if s == "" {
		return []string{}
	}
	res := []string{}
	rec(s, []string{}, &res)
	return res
}

func rec(s string, line []string, res *[]string) {
	if len(line) > 4 {
		return
	}
	if len(s) == 0 {
		if len(line) == 4 {
			*res = append(*res, strings.Join(line, "."))
			return
		}
		return
	}
	if s[0] == '0' {
		rec(s[1:], append(line, "0"), res)
		return
	}

	for i := 1; i <= 3; i++ {
		if i <= len(s) {
			v := atoi(s[:i])
			if v >= 0 && v <= 255 {
				rec(s[i:], append(line, s[:i]), res)
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
	fmt.Println(restoreIpAddresses("25525511135"))
}
