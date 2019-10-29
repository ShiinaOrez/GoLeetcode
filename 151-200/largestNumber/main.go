package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type dataSlice []string

func (slice dataSlice) Len() int {
	return len(slice)
}

func (slice dataSlice) Less(i, j int) bool {
	return strings.Compare(slice[i]+slice[j], slice[j]+slice[i]) == 1
}

func (slice dataSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
	return
}

func largestNumber(nums []int) string {
	length := len(nums)
	stringSlice := make(dataSlice, length)
	for index, num := range nums {
		stringSlice[index] = strconv.Itoa(num)
	}

	sort.Sort(stringSlice)
	ans := ""
	for _, str := range stringSlice {
		ans += str
	}
	ans = strings.TrimLeftFunc(ans, func(r rune) bool {
		return r == '0'
	})
	if ans == "" {
		return "0"
	}
	return ans
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
}
