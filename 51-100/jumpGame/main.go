package main

import "fmt"

func canJump(nums []int) bool {
    n := len(nums)
    if n == 0 {
        return false
    }
    
    for j:=n-2; j>=0; j-- {
        if nums[j] + j >= n-1  {
            n = j + 1
            fmt.Println(n)
        }
    }
    return !(n > 1)
}

func main() {
    fmt.Println(canJump([]int{3, 2, 1, 0, 1}))
}
