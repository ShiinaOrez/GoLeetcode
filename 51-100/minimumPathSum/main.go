package main

import "fmt"

func minPathSum(grid [][]int) int {
    n := len(grid)
    if n == 0 {
        return 0
    }
    m := len(grid[0])
    res := make([][]int, n)
    for i:=0; i<n; i++ {
        res[i] = make([]int, m)
    }
    
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if i == 0 && j == 0 {
                res[i][j] = grid[i][j]
            } else if i == 0 {
                res[i][j] = res[i][j-1] + grid[i][j]
            } else if j == 0 {
                res[i][j] = res[i-1][j] + grid[i][j]
            } else {
                res[i][j] = func (x, y int) int {
                    if x < y {
                        return x
                    }
                    return y
                }(res[i-1][j], res[i][j-1]) + grid[i][j]
            }
        }
    }
    return res[n-1][m-1]
}

func main() {
    fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
