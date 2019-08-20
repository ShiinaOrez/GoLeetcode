package main

import "fmt"

func findPeakElement(nums []int) int {
    length := len(nums)
    left := 0
    right := length-1
    for left < right {
        mid := (left+right)/2
        if nums[mid] > nums[mid+1] {
            right = mid
        } else {
            left = mid+1
        }
    }
    return left
}

func main() {
    fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}
