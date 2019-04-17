package main

import "fmt"

func rotate(matrix [][]int)  {
    n := len(matrix)
    for i:=0; i<n/2; i++ {
        for j:=i; j<n-1-i; j++ {
            matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
            fmt.Println(i, j)
            Print(matrix)
        }
    }
    return
}

func Print(res [][]int) {
    n := len(res)
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++ {
            fmt.Printf("%d ", res[i][j])
        }
        fmt.Println()
    }
}

func main() {
    res := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
    rotate(res)
    Print(res)
}
