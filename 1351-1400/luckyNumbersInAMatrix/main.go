package main

import (
    "fmt"
)

type Pos struct {
    y int
    v int
}

func luckyNumbers (matrix [][]int) []int {
    l := len(matrix)
    if l == 0 {
        return []int{}
    }
    if l == 1 {
        if len(matrix[0]) == 0 {
            return []int{}
        } else {
            m := matrix[0][0]
            for i:=1; i<len(matrix[0]); i++ {
                if matrix[0][i] < m {
                    m = matrix[0][i]
                }
            }
            return []int{m}
        }
    }
    
    h := len(matrix[0])
    if h == 1 {
        m := matrix[0][0]
        for i:=0; i<l; i++ {
            if matrix[i][0] > m {
                m = matrix[i][0]
            }
        }
        return []int{m}
    }
    poses := make([]Pos, l)
    ans := []int{}
    for i:=0; i<l; i++ {
        m := matrix[i][0]
        poses[i].y = 0
        for j:=1; j<h; j++ {
            if matrix[i][j] < m {
                m = matrix[i][j]
                poses[i].y = j
            }
        }
        poses[i].v = m
    }
    for i:=0; i<l; i++ {
        y := poses[i].y
        for x:=0; x<l; x++ {
            if matrix[x][y] > poses[i].v {
                break
            } else if x == l-1 {
                ans = append(ans, poses[i].v)
            }
        }
    }
    return ans
}

func main() {
    fmt.Println(luckyNumbers([][]int{[]int{3, 7, 8}, []int{9, 11, 13}, []int{15, 16, 17}}))
}
