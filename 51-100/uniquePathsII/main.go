package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n == 0 {
		return 0
	}
	m := len(obstacleGrid[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				res[i][j] = 0
			} else if i == 0 && j == 0 {
				res[i][j] = 1
			} else if i == 0 {
				res[i][j] = res[i][j-1]
			} else if j == 0 {
				res[i][j] = res[i-1][j]
			} else {
				res[i][j] = res[i-1][j] + res[i][j-1]
			}
		}
	}
	return res[n-1][m-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
