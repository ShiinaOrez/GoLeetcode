package main

import (
    "fmt"
)

func numTimesAllBlue(light []int) int {
    length := len(light)
    if length == 0 {
        return 0
    }
    if length == 1 {
        return 1
    }
    m := -1
    ans := 0
    for i:=0; i<length; i++ {
        if light[i] > m {
            m = light[i]
        }
        if m == i+1 {
            ans += 1
        }
    }
    return ans
}

func main() {
    fmt.Println(numTimesAllBlue([]int{2, 1, 3, 5, 4}))
}
