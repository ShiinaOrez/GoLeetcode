package main

import "fmt"

func totalNQueens(n int) int {
    status := []bool{}
    pos := []int{}
    tot := 0
    for i:=0; i<n; i++ {
        status = append(status, true)
        pos = append(pos, 0)
    }
    set(&tot,status, pos, n, n)
    return tot
}

func set(b *int, s []bool, pos []int, n, N int) {
    if n == 0 {
        *b += 1
        return
    } else {
        for i:=0; i<N; i++ {
            if s[i] && pos[i] == 0 {
                s[i] = false
                pos[i] = 3
                newPos := []int{}
                for j:=0; j<N; j++ {
                    newPos = append(newPos, 0)
                }
                for j:=0; j<N; j++ {
                    if pos[j] == 1 && j>=1 {
                        newPos[j-1] += 1
                    } else if pos[j] == 2 && j < N-1 {
                        newPos[j+1] += 2
                    } else if pos[j] == 3 {
                        if j>=1 {
                            newPos[j-1] += 1
                        }
                        if j<N-1 {
                            newPos[j+1] += 2
                        }
                    }
                }
                set(b, s, newPos, n-1, N)
                s[i] = true
                pos[i] = 0
            }
        }
    }
    return
}

func main() {
    fmt.Println(totalNQueens(8))
}
