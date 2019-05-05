package main

import "fmt"

func exist(board [][]byte, word string) bool {
    n := len(board)
    if n == 0 {
        return false
    }
    m := len(board[0])
    
    mask := make([][]bool, n)
    for i:=0; i<n; i++ {
        mask[i] = make([]bool, m)
    }
    
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if board[i][j] == word[0] {
                mask[i][j] = true
                res := search(board, word[1:], mask, i, j, n, m)
                if res {
                    return true
                } else {
                    mask[i][j] = false
                }
            }
        }
    }
    return false
}

func search(b [][]byte, word string, mask [][]bool, x, y, n, m int) bool {
    if word == "" {
        return true
    }
    if x-1 >= 0 {
        if b[x-1][y] == word[0] && !mask[x-1][y] {
            mask[x-1][y] = true
            res := search(b, word[1:], mask, x-1, y, n, m)
            if res {
                return res
            } else {
                mask[x-1][y] = false
            }
        }
    }
    if x+1 < n {
        if b[x+1][y] == word[0] && !mask[x+1][y] {
            mask[x+1][y] = true
            res := search(b, word[1:], mask, x+1, y, n, m)
            if res {
                return res
            } else {
                mask[x+1][y] = false
            }
        }
    }
    if y+1 < m {
        if b[x][y+1] == word[0] && !mask[x][y+1] {
            mask[x][y+1] = true
            res := search(b, word[1:], mask, x, y+1, n, m)
            if res {
                return res
            } else {
                mask[x][y+1] = false
            }
        }
    }
    if y-1 >= 0 {
        if b[x][y-1] == word[0] && !mask[x][y-1] {
            mask[x][y-1] = true
            res := search(b, word[1:], mask, x, y-1, n, m)
            if res {
                return res
            } else {
                mask[x][y-1] = false
            }
        }
    }
    return false
}

func main() {
    fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "AECCED"))
}
