package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	blockPoint := [][2]int{
		[2]int{1, 1},
		[2]int{1, 4},
		[2]int{1, 7},
		[2]int{4, 1},
		[2]int{4, 4},
		[2]int{4, 7},
		[2]int{7, 1},
		[2]int{7, 4},
		[2]int{7, 7},
	}
	for _, point := range blockPoint {
		if ok := checkBlock(board, point); !ok {
			return false
		}
	}

	numMap := make(map[byte]bool)
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			now := board[i][j]
			if _, ok := numMap[now]; !ok {
				numMap[now] = true
			} else if now != '.' {
				return false
			}
		}
		numMap = make(map[byte]bool)
	}

	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			now := board[j][i]
			if _, ok := numMap[now]; !ok {
				numMap[now] = true
			} else if now != '.' {
				return false
			}
		}
		numMap = make(map[byte]bool)
	}

	return true
}

func checkBlock(board [][]byte, point [2]int) bool {
	x := point[0]
	y := point[1]
	numMap := make(map[byte]bool)
	now := board[x][y]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}

	now = board[x-1][y-1]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}

	now = board[x-1][y]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}

	now = board[x][y-1]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}

	now = board[x+1][y-1]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}

	now = board[x-1][y+1]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}

	now = board[x][y+1]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}

	now = board[x+1][y]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}

	now = board[x+1][y+1]
	if _, ok := numMap[now]; !ok {
		numMap[now] = true
	} else if now != '.' {
		return false
	}
	return true
}

func main() {
	fmt.Println()
}
