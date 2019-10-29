package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	rest, run, start := 0, 0, 0
	for i, _ := range gas {
		run += (gas[i] - cost[i])
		rest += (gas[i] - cost[i])
		if run < 0 {
			start = i + 1
			run = 0
		}
	}
	if rest < 0 {
		return -1
	}
	return start
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
