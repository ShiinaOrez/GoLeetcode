package main

import "fmt"

func solve(board [][]byte)  {
    height := len(board)
    if height == 0 {
        return
    }
    width := len(board[0])
    mask := make([][]bool, height)
    for i:=0; i<height; i++ {
        mask[i] = make([]bool, width)
    }
    
    var fill func(x,y int)
    
    fill = func(x,y int) {
        mask[x][y] = true
        if x-1 >= 0 && board[x-1][y] == 'O' && !mask[x-1][y] {
            fill(x-1, y)
        }
        if x+1 < height && board[x+1][y] == 'O' && !mask[x+1][y] {
            fill(x+1, y)
        }
        if y-1 >= 0 && board[x][y-1] == 'O' && !mask[x][y-1] {
            fill(x, y-1)
        }
        if y+1 < width && board[x][y+1] == 'O' && !mask[x][y+1] {
            fill(x, y+1)
        }
    }
    
    for i:=0; i<height; i++ {
        if board[i][0] == 'O' && !mask[i][0] {
            fill(i, 0)
        }
        if board[i][width-1] == 'O' && !mask[i][width-1] {
            fill(i, width-1)
        }
    }
    for j:=0; j<width; j++ {
        if board[0][j] == 'O' && !mask[0][j] {
            fill(0, j)
        }
        if board[height-1][j] == 'O' && !mask[height-1][j] {
            fill(height-1, j)
        }
    }
    for i:=0; i<height; i++ {
        for j:=0; j<width; j++ {
            if !mask[i][j] {
                board[i][j] = 'X'
            }
        }
    }
    return
}

func main() {
	fmt.Println()
}