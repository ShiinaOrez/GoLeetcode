package main

import "fmt"

func maxSubArray(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    m := ^int(^uint(0) >> 1)
    if nums[0] > m {
        m = nums[0]
    }
    for i:=1; i<n; i++ {
        if nums[i-1] > 0 {
            nums[i] = nums[i] + nums[i-1]
        }
        if nums[i] > m {
            m = nums[i]
        }
    }
    return m
}

func main() {
    fmt.Println(maxSubArray([]int{1, 2}))
}
