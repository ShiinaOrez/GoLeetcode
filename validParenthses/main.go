package main

import "fmt"

func isValid(s string) bool {
    a := 0
    a_l := []int{}
    b := 0
    b_l := []int{}
    c := 0
    c_l := []int{}
    for i, r := range(s) {
        ss := string(r)
        if ss == "(" {
            a += 1
            a_l = append(a_l, i)
        } else if ss == "{" {
            b += 1
            b_l = append(b_l, i)
        } else if ss == "[" {
            c += 1
            c_l = append(c_l, i)
        } else if ss == ")" {
            a -= 1
            if a < 0 {
                return false
            }
            if len(b_l) > 0 {
                if b_l[len(b_l)-1] > a_l[len(a_l)-1] {
                    return false
                }
            }
            if len(c_l) > 0 {
                if c_l[len(c_l)-1] > a_l[len(a_l)-1] {
                    return false
                }
            }
            a_l = a_l[:len(a_l)-1]
        } else if ss == "}" {
            b -= 1
            if b < 0 {
                return false
            }
            if len(a_l) > 0 {
                if b_l[len(b_l)-1] < a_l[len(a_l)-1] {
                    return false
                }
            }
            if len(c_l) > 0 {
                if c_l[len(c_l)-1] > b_l[len(b_l)-1] {
                    return false
                }
            }
            b_l = b_l[:len(b_l)-1]
        } else if ss == "]" {
            c -= 1
            if c < 0 {
                return false
            }
            if len(b_l) > 0 {
                if b_l[len(b_l)-1] > c_l[len(c_l)-1] {
                    return false
                }
            }
            if len(a_l) > 0 {
                if c_l[len(c_l)-1] < a_l[len(a_l)-1] {
                    return false
                }
            }
            c_l = c_l[:len(c_l)-1]
        }
    }
    if a != 0 || b != 0 || c != 0 {
        return false
    }
    return true
}

func main() {
    fmt.Println(isValid("{(})"))
}
