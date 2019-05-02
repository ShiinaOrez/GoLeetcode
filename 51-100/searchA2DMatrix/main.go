package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
    n := len(matrix)
    if n == 0 {
        return false
    }
    m := len(matrix[0])
    if m == 0 {
        return false
    }
    
    start := 0
    end := n-1
    for start < end {
        if matrix[start][0] == target || matrix[end][0] == target {
            return true
        }
        mid := (start + end) / 2
        if target <= matrix[mid][0] {
            end = mid
        } else {
            start = mid
        }
        if start == end - 1 {
            if target >= matrix[end][0] {
                start = end
            }
            break
        }
    }
    if matrix[start][0] == target {
        return true
    }
    now := start
    start = 0
    end = m-1
    for start < end {
        if matrix[now][start] == target || matrix[now][end] == target {
            return true
        }
        mid := (start + end) / 2
        if target <= matrix[now][mid] {
            end = mid
        } else {
            start = mid
        }
        if start == end - 1 {
            if matrix[now][start] == target || matrix[now][end] == target {
                return true
            } else {
                return false
            }
        }
    }
    return matrix[now][start] == target
}

func main() {
    fmt.Println(searchMatrix([][]int{{1}, {3}, {5}}, 3))
}
