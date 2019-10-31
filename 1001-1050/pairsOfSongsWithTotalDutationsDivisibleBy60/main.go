package main

import "fmt"

func numPairsDivisibleBy60(time []int) int {
	tot := 0
	restMap := make(map[int]int)
	for i := 0; i < 60; i++ {
		restMap[i] = 0
	}
	for i, t := range time {
		time[i] = t % 60
		restMap[t%60] += 1
	}
	for i := 1; i < 60; i++ {
		if i == 30 {
			continue
		}
		tot += restMap[i] * restMap[60-i]
	}
	tot /= 2
	for i := 1; i < restMap[30]; i++ {
		tot += i
	}
	for i := 1; i < restMap[0]; i++ {
		tot += i
	}
	return tot
}

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{20, 40, 100, 150, 30}))
}
