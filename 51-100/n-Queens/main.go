package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	board := [][]string{}
	now := []string{}
	status := []bool{}
	pos := []int{}
	for i := 0; i < n; i++ {
		status = append(status, true)
		pos = append(pos, 0)
	}
	set(&board, now, status, pos, n, n)
	return board
}

func set(b *[][]string, now []string, s []bool, pos []int, n, N int) {
	if n == 0 {
		res := make([]string, N)
		copy(res, now)
		(*b) = append(*b, res)
		return
	} else {
		newLine := []byte{}
		for i := 0; i < N; i++ {
			newLine = append(newLine, '.')
		}
		for i := 0; i < N; i++ {
			if s[i] && pos[i] == 0 {
				newLine[i] = 'Q'
				s[i] = false
				pos[i] = 3
				newPos := []int{}
				for j := 0; j < N; j++ {
					newPos = append(newPos, 0)
				}
				for j := 0; j < N; j++ {
					if pos[j] == 1 && j >= 1 {
						newPos[j-1] += 1
					} else if pos[j] == 2 && j < N-1 {
						newPos[j+1] += 2
					} else if pos[j] == 3 {
						if j >= 1 {
							newPos[j-1] += 1
						}
						if j < N-1 {
							newPos[j+1] += 2
						}
					}
				}
				//                fmt.Println("NewLine:", newLine)
				//                fmt.Println("Position:", newPos)
				set(b, append(now, string(newLine)), s, newPos, n-1, N)
				newLine[i] = '.'
				s[i] = true
				pos[i] = 0
			}
		}
	}
	return
}

func print(b []string) {
	for _, str := range b {
		fmt.Println(str)
	}
}

func main() {
	fmt.Println(len(solveNQueens(7)))
}
