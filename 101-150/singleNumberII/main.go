package main

import "fmt"

func singleNumber(nums []int) int {
    a, b := 0, 0
    for _, num := range nums {
        b = (b^num)&^a 
        a = (a^num)&^b
    }
    return b
}

func main() {
    fmt.Println(singleNumber([]int{2, 2, 1, 2}))
}
