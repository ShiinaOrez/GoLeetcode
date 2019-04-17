package main

import "fmt"

func longestValidParentheses(s string) int {
    max := func (x,  y int) int {
        if x >= y {
            return x
        }
        return y
    }
    length := len(s)
    return max(calc(s, 0, 1, length, '('), calc(s, length-1, -1, -1, ')'))
}

func calc(s string, i, flag, end int, cTem byte) int {
    var max, sum, currLen, validLen int
    for ; i != end; i += flag {
        if s[i] == cTem {
            sum += 1
        } else {
            sum -= 1
        }
        currLen += 1
        if sum < 0 {
            if max < validLen {
                max = validLen
            }
            sum = 0
            currLen = 0
            validLen = 0
        } else if sum == 0 {
            validLen = currLen
        }
    }
    if max > validLen {
        return max
    }
    return validLen
}

func main() {
    fmt.Println(longestValidParentheses("))))())()()(()"))
}
