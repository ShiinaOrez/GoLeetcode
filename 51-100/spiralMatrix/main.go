package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}

	m := len(matrix[0])
	mask := make([][]bool, n+2)
	for i := 0; i < n+2; i++ {
		mask[i] = make([]bool, m+2)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			mask[i][j] = true
		}
	}
	i := 1
	j := 1
	res := []int{}
	for k := 1; k <= n*m; k++ {
		res = append(res, matrix[i-1][j-1])
		mask[i][j] = false
		if !mask[i-1][j] && !mask[i][j-1] && !mask[i+1][j] {
			j += 1
			continue
		}
		if !mask[i][j-1] && !mask[i-1][j] && !mask[i][j+1] {
			i += 1
			continue
		}
		if !mask[i][j+1] && !mask[i-1][j] && !mask[i+1][j] {
			j -= 1
			continue
		}
		if !mask[i][j-1] && !mask[i][j+1] && !mask[i+1][j] {
			i -= 1
			continue
		}
		if !mask[i-1][j] {
			if !mask[i][j+1] {
				i += 1
			} else {
				j += 1
			}
			continue
		}

		if !mask[i][j+1] {
			if !mask[i+1][j] {
				j -= 1
			} else {
				i += 1
			}
			continue
		}

		if !mask[i+1][j] {
			if !mask[j-1][i] {
				i -= 1
			} else {
				j -= 1
			}
			continue
		}

		if !mask[j-1][i] {
			if !mask[i-1][j] {
				j += 1
			} else {
				i -= 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
