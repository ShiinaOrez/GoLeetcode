package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	if version1 == version2 {
		return 0
	}
	versionSlice1 := strings.Split(version1, ".")
	versionSlice2 := strings.Split(version2, ".")
	length := max(len(versionSlice1), len(versionSlice2))
	versionNum1 := make([]int, length)
	for i, _ := range versionSlice1 {
		num, _ := strconv.Atoi(versionSlice1[i])
		versionNum1[i] = num
	}
	versionNum2 := make([]int, length)
	for i, _ := range versionSlice2 {
		num, _ := strconv.Atoi(versionSlice2[i])
		versionNum2[i] = num
	}
	fmt.Println(versionNum1, versionNum2, length)
	for i := 0; i < length; i++ {
		sub1 := versionNum1[i]
		sub2 := versionNum2[i]
		fmt.Println(sub1, sub2)
		if sub1 > sub2 {
			return 1
		} else if sub1 < sub2 {
			return -1
		}
	}
	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(compareVersion("1", "1.1"))
}
