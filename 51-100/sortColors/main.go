package main

import "fmt"

func sortColors(nums []int)  {
    red := []int{}
    white := []int{}
    blue := []int{}
    n := len(nums)
    for i:=0; i<n; i++ {
        if nums[i] == 0 {
            red = append(red, 0)
        } else if nums[i] == 1 {
            white = append(white, 1)
        } else {
            blue = append(blue, 2)
        }
    }
    newnums := append(red, white...)
    newnums = append(newnums, blue...)
    copy(nums, newnums)
    return
}

func main() {
    nums := []int{0, 0, 2, 2, 1, 1}
    sortColors(nums)
    fmt.Println(nums)
}
