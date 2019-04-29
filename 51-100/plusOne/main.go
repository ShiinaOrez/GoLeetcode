package main

import "fmt"

func plusOne(digits []int) []int {
    n := len(digits)
    if n == 0 {
        return []int{}
    }
    next := 1
    for i:=n-1; i>=0; i-- {
        digits[i] += next
        next = digits[i] / 10
        digits[i] = digits[i] % 10
    }
    if next > 0 {
        digits = append([]int{next}, digits...)
    }
    return digits
}

func main() {
    fmt.Println(plusOne([]int{9,9,9,9}))
}
