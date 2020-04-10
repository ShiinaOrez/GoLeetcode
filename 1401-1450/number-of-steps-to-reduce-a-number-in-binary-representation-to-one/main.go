package main

import "fmt"

func numSteps(s string) int {
    ans := 0
    for s != "1" {
        if s[len(s)-1] == '1' {
            for i:=len(s)-1; i>=0; i-- {
                if s[i] == '1' {
                    s = s[:i] + "0" + s[i+1:]
                    if i == 0 {
                        s = "1" + s
                    }
                } else {
                    s = s[:i] + "1" + s[i+1:]
                    break
                }
            }
        } else {
            s = s[:len(s)-1]
        }
        ans += 1
    }
    return ans
}

func main() {
    fmt.Println(numSteps("1101"))
}
