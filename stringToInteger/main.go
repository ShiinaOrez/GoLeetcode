package main

import "fmt"

func myAtoi(str string) int {
    var ans, intMax int64
    ans = 0
    intMax = 1<<31-1
    zero := int64('0')
    nine := int64('9')
    flag := false
    haveFlag := false
    haveNum := false
    for i, _ := range(str) {
        c := str[i]
        if c == byte('+') {
            if haveNum {
                break
            }
            if haveFlag {
                return 0
            }
            haveFlag = true
            continue
        } else if c == byte('-') {
            if haveNum {
                break
            }
            if haveFlag {
                return 0
            }
            flag = true
            haveFlag = true
            continue
        } else if c == byte(' ') {
            if haveNum || haveFlag {
                break
            }
            continue
        } else if c == byte('.') {
            break
        } else if int64(c) >= zero && int64(c) <= nine {
            haveNum = true
            ans *= 10
            ans += int64(c) - zero
            if ans > intMax && !flag {
                return int(intMax)
            } else if ans > intMax + 1 && flag{
                return -1 * int(intMax + 1)
            }
        } else {
            break
        }
    }
    if flag {
        return -1 * int(ans)
    }
    return int(ans)
}

func main() {
    fmt.Println(myAtoi("+0 123"))
}
