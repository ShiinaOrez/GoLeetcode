package main

import "fmt"

func uniquePaths(m int, n int) int {
    res := make([][]int, n)
    for i:=0; i<n; i++ {
        res[i] = make([]int, m)
    }
    for i:=0; i<n; i++ {
        res[i][0] = 1
    }
    for i:=0; i<m; i++ {
        res[0][i] = 1
    }
    for i:=1; i<n; i++ {
        for j:=1; j<m; j++ {
            res[i][j] = res[i-1][j] + res[i][j-1]
        }
    }
    return res[n-1][m-1]
}

func main() {
    fmt.Println(uniquePaths(7, 3))
}
