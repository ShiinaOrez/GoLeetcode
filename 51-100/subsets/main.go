package main

import "fmt"

func subsets(nums []int) [][]int {
    res := [][]int{}
    dfs(0, len(nums)-1, []int{}, nums, &res)
    return res
}

func dfs(now, end int, line, nums []int, res *[][]int) {
    for i:=now; i<=end; i++ {
        newLine := make([]int, len(line)+1)
        copy(newLine, line)
        newLine[len(line)] = nums[i]
        dfs(i+1, end, newLine, nums, res)
    }
    *res = append(*res, line)
    return
}

func main() {
    fmt.Println(subsets([]int{1, 2, 3}))
}
