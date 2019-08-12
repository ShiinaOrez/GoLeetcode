package main

import (
	"fmt"
)

var res [][]string
var line []string 

func partition(s string) [][]string {
    if s == "" {
        return [][]string{}
    }
    res = [][]string{}
    line = []string{}
    divide(s, 0, len(s)-1)
    return res
}

func divide(s string, a, b int) {
    if a > b {
        res = append(res, line)
    }
    for i:=1; i<=b-a+1; i++ {
        if judge(s[a:a+i]) {
            line = append(line, s[a:a+i])
            divide(s, a+i, b)
            newLine := line
            line = make([]string, len(newLine)-1)
            copy(line, newLine)
        }
    }
}

func judge(s string) bool {
    length := len(s)
    for i:=0; i<length/2+1; i++ {
        if s[i] != s[length-1-i] {
            return false
        }
    }
    return true
}

func main() {
	fmt.Println(partition("cbbbcc"))
}