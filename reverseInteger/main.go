package main

import "fmt"

func reverse(x int) int {
    flag := false
    if x < 0 {
        flag = true
    }
    var res []int
    if flag {
        x = -1 * x
    }
    for x > 0 {
        res = append(res, x%10)
        x = x/10
    }
    ans := 0
    for _, i := range(res) {
        ans *= 10
        ans += i
    }
    if ans > (1<<31) - 1 {
        return 0;
    }
    if flag {
        ans = -1 * ans
    }
    return ans
}

func main() {
    fmt.Println(reverse(274863721))
}
