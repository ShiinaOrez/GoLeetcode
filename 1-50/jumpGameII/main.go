package main

import "fmt"

func jump(nums []int) int {
    length := len(nums)
    arr := make([]int, length)
    for i:=0; i<length; i++ {
        arr[i] = jumpTo(nums[:i+1], arr)
    }
    return arr[length-1]
}

func jumpTo(nums, arr []int) int {
    res := -1
    length := len(nums)
    
    if length == 1 {
        return 0
    }
    
    for i:=length-2; i>=0; i-- {
        if nums[i] >= length-i-1 && nums[i] != 0 {
            if i == 0 {
                return 1
            }
            if arr[i] < 0 {
                continue
            }
            if arr[i]+1 < res || res == -1 {
                res = arr[i]+1
            }
        }
    }
    return res
}

func main() {
    fmt.Println(jump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}
