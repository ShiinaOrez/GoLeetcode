package main

import "fmt"

func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    maxLength := 0
    var ran [2]int
    for i, r := range(s) {
        c := byte(r)
        if len(s) == 2 {
            if s[0] == s[1] {
                return s
            }
        }
        if i-2 >= 0 {
            if s[i-2] == c {
                length := 3
                if length > maxLength {
                    maxLength = length
                    ran[0] = i-2
                    ran[1] = i
                }
                for d := 2; (i-1-d >= 0) && (i-1+d < len(s)); d++ {
                    if s[i-1-d] == s[i-1+d] {
                        length += 2
                        if length > maxLength {
                            maxLength = length
                            ran[0] = i-1-d
                            ran[1] = i-1+d
                        }
                    } else {
                        break
                    }
                }
            }
        }
        if i-1 >= 0 {
            if s[i-1] == c {
                length := 2
                if length > maxLength {
                    maxLength = length
                    ran[0] = i-1
                    ran[1] = i
                }
                for d := 1; (i-(d+1) >= 0) && (i+d < len(s)); d++ {
                    if s[i-(d+1)] == s[i+d] {
                        length += 2
                        if length > maxLength {
                            maxLength = length
                            ran[0] = i-d-1
                            ran[1] = i+d
                        }
                    } else {
                        break
                    }
                }
            }
        }
    }
    return s[ran[0]:ran[1]+1]
}

func main() {
    fmt.Println(longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}