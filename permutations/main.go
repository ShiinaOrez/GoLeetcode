package main

import "fmt"

func permute(nums []int) [][]int {
    res := [][]int{}
    mask := []bool{}
    length := len(nums)
    for i:=0; i<length; i++ {
        mask = append(mask, false)
    }
    generate(&res, []int{}, nums, mask)
    return res
}

func generate(res *[][]int, now, nums []int, mask []bool) {
    if len(now) == len(nums) {
        *res = append(*res, now)
        return
    }
    
    for ind, i := range mask {
        if !i {
            next := append(now, nums[ind])
            mask[ind] = true
            generate(res, next, nums, mask)
            mask[ind] = false
        }
    }
    return
}

func main() {
    fmt.Println(permute([]int{3, 2, 1}))
}
