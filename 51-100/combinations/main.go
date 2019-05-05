package main

import "fmt"

func combine(n int, k int) [][]int {
    res := [][]int{}
    dfs(1, n, k, []int{}, &res)
    return res
}

func dfs(now, n, k int, line []int, res *[][]int) {
    if k == 0 {
        *res = append(*res, line)
        return
    }
    for i:=now; i<=n; i++ {
        newLine := make([]int, len(line)+1)
        copy(newLine, line)
        newLine[len(line)] = i
        dfs(i+1, n, k-1, newLine, res)
    }
    return
}

func main() {
    fmt.Println(combine(4, 2))
}
