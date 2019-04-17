package main

import (
    "sort"
    "fmt"
)

func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    length := len(nums)
    mask := make(map[int]int)
    sort.Ints(nums)
    for i:=0; i<length; i++ {
        if v, ok := mask[nums[i]]; !ok {
            mask[nums[i]] = 1
        } else {
            mask[nums[i]] = v + 1
        }
    }
    generate(&res, []int{}, nums, mask)
    return res
}

func generate(res *[][]int, now, nums []int, mask map[int]int) {
    fmt.Println("Now:", now)
    if len(now) == len(nums) {
        *res = append(*res, now)
        return
    }
    length := len(nums)
    before :=  -234
    for ind:=0; ind<length; ind++ {
        if v, _ := mask[nums[ind]]; v > 0 {
            if before == nums[ind] {
                continue
            }
            next := make([]int, len(now)+1)
            copy(next, now)
            next[len(now)] = nums[ind]
            mask[nums[ind]] = v-1
            generate(res, next, nums, mask)
            mask[nums[ind]] = v
            before = nums[ind]
        }
    }
    return
}

func main() {
    fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
}
