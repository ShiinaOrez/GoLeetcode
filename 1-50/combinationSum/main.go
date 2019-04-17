package main

import (
    "fmt"
    "sort"
)

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    // response
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
//        fmt.Println("Gen:", row)
        generate(res, candidates, row, target, i, sum + item)
    }
    return
}

func main() {
    fmt.Println(combinationSum([]int{2, 3, 7}, 18))
}

/*
func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    combMap := make(map[int][][]int)
    combMap = generate(candidates, combMap, target)
    return combMap[target]
}

func generate(candidates []int, combMap map[int][][]int, target int) map[int][][]int {
    if target < candidates[0] {
        combMap[0] = [][]int{}
        return combMap
    }
    
    for j:=len(candidates)-1; j>=0; j-- {
        now := now
        for i:=1; i*now <= target; i++ {
            sub := target - i*now
            
            var value [][]int
            if _, ok := combMap[sub]; !ok {
                combMap = generate(candidates, combMap, sub)
            }
            value, _ = combMap[sub]
            
            newRow := []int{}
            for k:=0; k<i; k++ {
                newRow = append(newRow, now)
            }
            
            for i, row := range value {
                value[i] = append(row, newRow...)
            }
            if len(value) == 0 && sub == 0 {
                value = [][]int{newRow}
            }
            
            if nowRes, ok := combMap[target]; !ok {
                combMap[target] = value
            } else {
                combMap[target] = append(nowRes, value...)
            }
        }
    }
    return combMap
}
*/