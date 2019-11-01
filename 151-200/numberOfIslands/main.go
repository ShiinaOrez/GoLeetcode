package main

import "fmt"

var width, height int

func removeIsland(m [][]byte, x, y int) {
	m[x][y] = 'x'
	if x-1 >= 0 && m[x-1][y] == '1' {
		removeIsland(m, x-1, y)
	}
	if x+1 < width && m[x+1][y] == '1' {
		removeIsland(m, x+1, y)
	}
	if y-1 >= 0 && m[x][y-1] == '1' {
		removeIsland(m, x, y-1)
	}
	if y+1 < height && m[x][y+1] == '1' {
		removeIsland(m, x, y+1)
	}
	return
}

func numIslands(grid [][]byte) int {
	width = len(grid)
	tot := 0
	if width != 0 {
		height = len(grid[0])
	} else {
		return tot
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if grid[i][j] == '1' {
				tot += 1
				removeIsland(grid, i, j)
			}
		}
	}
	return tot
}

func main() {
	fmt.Println()
}
