package main

import (
    "fmt"
)

var m map[int][]int
var count int = 0
var iT []int

func dfs(now, index, deepth int) {
    if deepth > count {
        count = deepth
    }
    if v, ok := m[m[now][index]]; !ok {
        return
    } else {
        l := len(v)
        for i:=0; i<l; i++ {
            dfs(m[now][index], i, deepth+iT[m[now][index]])
        }
    }
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    if n == 1 || n == 0 {
        return 0
    }

    m = make(map[int][]int)
    for i:=0; i<n; i++ {
        if v, ok := m[manager[i]]; !ok {
            m[manager[i]] = []int{i}
        } else {
            m[manager[i]] = append(v, i)
        }
    }
    count = 0
    iT = informTime
    
    dfs(-1, 0, 0)
    return count
}

func main() {
    fmt.Println()
}
