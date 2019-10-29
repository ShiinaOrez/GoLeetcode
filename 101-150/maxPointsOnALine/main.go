package main

import (
	"fmt"
	"strconv"
)

func maxPoints(points [][]int) int {
	count := len(points)
	if count <= 2 {
		return count
	}
	maxPoints := 0
	for i := 0; i < count; i++ {
		points[i] = []int{points[i][0] + 1, points[i][1] + 1}
	}
	for i := 0; i < count; i++ {
		m := make(map[string]int)
		same := 1
		for j := 0; j < count; j++ {
			if i == j {
				continue
			}
			g := gcd(points[j][0]-points[i][0], points[j][1]-points[i][1])
			if g == 0 {
				same += 1
				continue
			}
			x, y := (points[j][0]-points[i][0])/g, (points[j][1]-points[i][1])/g
			m[strconv.Itoa(x)+"-"+strconv.Itoa(y)] += 1
		}
		max := 0
		for k := range m {
			if m[k] > max {
				max = m[k]
			}
		}
		if max+same > maxPoints {
			maxPoints = max + same
		}
	}
	return maxPoints
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	fmt.Println(maxPoints([][]int{
		[]int{0, 0},
		[]int{1, 1},
		[]int{0, 0},
	}))
}
