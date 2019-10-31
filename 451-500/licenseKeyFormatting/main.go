package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	s := strings.ToUpper(S)
	s = strings.Replace(s, "-", "", -1)
	slice := []string{}
	for len(s) >= K {
		slice = append([]string{s[len(s)-K:]}, slice...)
		s = s[:len(s)-K]
	}
	if len(s) > 0 {
		slice = append([]string{s}, slice...)
	}
	return strings.Join(slice, "-")
}

func main() {
	fmt.Println(licenseKeyFormatting("378ry-38r7-3r2yr23", 3))
}
