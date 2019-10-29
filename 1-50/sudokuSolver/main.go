package main

import "fmt"

func solveSudoku(board [][]byte) {
	_ = solve(&board, 0, 0)
}

func solve(b *[][]byte, x, y int) bool {
	if x == 8 && y == 8 {
		return true
	}
	//    fmt.Printf("%d - %d \n", x, y)
	//    Print(*b)
	for i := 0; i <= 8; i++ {
		if i < x {
			continue
		}
		for j := 0; j <= 8; j++ {
			if i == x && j < y {
				continue
			}
			if (*b)[i][j] == '.' {

				//                fmt.Printf("\nFill: [%d, %d] ", i, j)

				for k := '1'; k <= '9'; k++ {
					//                    fmt.Printf("%c: ", byte(k))
					(*b)[i][j] = byte(k)
					if checkRow(b, i) && checkLine(b, j) && checkBoard(b) {
						ok := solve(b, i, j)
						if !ok {
							//                            fmt.Printf("fail. ")
							(*b)[i][j] = '.'
						} else {
							return true
						}
					} else {
						//                        fmt.Printf("fail. ")
						(*b)[i][j] = '.'
					}
				}
				if (*b)[i][j] == '.' {
					return false
				}
			}
		}
	}
	return true
}

func checkRow(b *[][]byte, row int) bool {
	numMap := make(map[byte]bool)
	for i := 0; i <= 8; i++ {
		now := (*b)[row][i]
		if _, ok := numMap[now]; !ok {
			numMap[now] = true
		} else if now != '.' {
			return false
		}
	}
	return true
}

func checkLine(b *[][]byte, line int) bool {
	numMap := make(map[byte]bool)
	for i := 0; i <= 8; i++ {
		now := (*b)[i][line]
		if _, ok := numMap[now]; !ok {
			numMap[now] = true
		} else if now != '.' {
			return false
		}
	}
	return true
}

func checkBoard(b *[][]byte) bool {
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
		if ok := checkBlock(b, point); !ok {
			return false
		}
	}
	return true
}

func checkBlock(b *[][]byte, point [2]int) bool {
	board := (*b)
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

func Print(b [][]byte) {
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			fmt.Printf("%c ", b[i][j])
		}
		fmt.Println()
	}
}

func main() {
	res := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(res)
}
