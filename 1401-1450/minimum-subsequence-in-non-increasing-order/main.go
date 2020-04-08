package main

import (
    "fmt"
    "sort"
)

func minSubsequence(nums []int) []int {
    count, ansCount := 0, 0
    for _, i := range nums {
        count += i
    }
    sort.Ints(nums)
    ans := make([]int, 0)
    l := len(nums)
    for i:=l-1; i>=0; i-- {
        ans = append(ans, nums[i])
        ansCount += nums[i]
        count -= nums[i]
        if ansCount > count {
            return ans
        }
    }
    return ans
}

func main() {
    fmt.Println(minSubsequence([]int{2, 3, 4, 1, 5}))
}
