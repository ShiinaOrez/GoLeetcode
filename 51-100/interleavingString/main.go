package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	point1 := 0
	point2 := 0
	len1 := len(s1)
	len2 := len(s2)
	len3 := len(s3)
	if len3 != len2+len1 {
		return false
	}
	if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}

	for i, r := range s3 {
		if s1[point1] == byte(r) && s2[point2] == byte(r) {
			res1 := isInterleave(s1[point1+1:], s2[point2:], s3[i+1:])
			if res1 {
				return true
			} else {
				res2 := isInterleave(s1[point1:], s2[point2+1:], s3[i+1:])
				return res2
			}
		} else {
			if s1[point1] == byte(r) {
				point1 += 1
				if point1 == len1 {
					return s3[i+1:] == s2[point2:]
				}
				continue
			} else if s2[point2] == byte(r) {
				point2 += 1
				if point2 == len2 {
					return s3[i+1:] == s1[point1:]
				}
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isInterleave("a", "", "a"))
}
