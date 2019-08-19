package main

import "fmt"

func findMin(nums []int) int {
    left := 0
    right := len(nums)-1
    if nums[left] < nums[right] {
        return nums[0]
    }
    for left < right {
        mid := (left+right)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else if nums[mid] < nums[right] {
            right = mid
        } else {
            right -= 1
        }
    }
    return nums[left]
}

func main() {
    fmt.Println(findMin([]int{3, 3, 1, 3}))
}
