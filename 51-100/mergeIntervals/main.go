package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	for i := 0; i < n; i++ {
		v := intervals[i]
		start := v[0]
		end := v[1]
		j := i
		for intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
			i++
			if i == n {
				break
			}
		}
		if j != i {
			i -= 1
		}
		res = append(res, []int{start, end})
	}
	return res
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
