package main

import (
    "fmt"
    "sort"
)

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)   
    res := [][]int{}
    generate(&res, candidates, []int{}, target, 0, 0)
    return res
}

func generate(res *[][]int, candidates, newRow []int, target, start, sum int) {
    if sum == target {
        (*res) = append((*res), newRow)
        return
    } else if sum > target {
        return
    }
    length := len(candidates)
    for i:=start; i<length; i++ {
        item := candidates[i]
        row := make([]int, len(newRow)+1)
        copy(row, newRow)
        row[len(row)-1] = item
        generate(res, candidates, row, target, i+1, sum + item)
        
        next := i
        for candidates[next] == candidates[i] {
            next += 1
            if next == length {
                break
            }
        }
        i = next-1
    }
    return
}

func main() {
    fmt.Println(combinationSum2([]int{2, 3, 1, 7, 1}, 7))
}
