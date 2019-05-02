package main

import "fmt"

func setZeroes(matrix [][]int)  {
    memx := -1
    memy := -1
    flag := false
    n := len(matrix)
    if n == 0 {
        return
    }
    m := len(matrix[0])
    
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if !flag && matrix[i][j] == 0 {
                memx = i
                memy = j
                flag = true
            }
            if matrix[i][j] == 0 {
                matrix[memx][j] = 0
                matrix[i][memy] = 0
            }
        }
    }
    if !flag {
        return
    }
    for i:=0; i<n; i++ {
        if i == memx {
            continue
        }
        if matrix[i][memy] == 0 {
            for j:=0; j<m; j++ {
                matrix[i][j] = 0
            }
        }
    }
    for j:=0; j<m; j++ {
        if j == memy {
            continue
        }
        if matrix[memx][j] == 0 {
            for i:=0; i<n; i++ {
                matrix[i][j] = 0
            }
        }
    }
    for i:=0; i<n; i++ {
        matrix[i][memy] = 0
    }
    for j:=0; j<m; j++ {
        matrix[memx][j] = 0
    }
    return
}

func main() {
    matrix := [][]int{{1,1,1}, {1,0,1}, {1,1,1}}
    setZeroes(matrix)
    fmt.Println(matrix)
}
