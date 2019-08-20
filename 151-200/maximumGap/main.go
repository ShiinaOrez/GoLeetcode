package main

import (
    "fmt"
    "sort"
)

func maximumGap(nums []int) int {
    length := len(nums)
    if length < 2 {
        return 0
    }
    sort.Ints(nums)
    max := nums[1]-nums[0]
    for i:=0; i<length-1; i++ {
        if nums[i+1] - nums[i] > max {
            max = nums[i+1]-nums[i]
        }
    }
    return max
}

func main() {
    fmt.Println(maximumGap([]int{3, 6, 9, 1}))
}

