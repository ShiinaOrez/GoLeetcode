package main

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}

	p := make([][]int, n+1)
	for i, v := range intervals {
		if v[0] <= newInterval[0] {
			copy(p, intervals[:i+1])
			p[i+1] = newInterval
			if i+1 < n {
				copy(p[i+2:], intervals[i+1:])
			}
			break
		}
		if i == 0 && v[0] > newInterval[0] {
			p[0] = newInterval
			copy(p[1:], intervals)
			break
		}
	}
	return merge(p)
}

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
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}
