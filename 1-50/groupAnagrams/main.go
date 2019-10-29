package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	if length == 0 {
		return [][]string{}
	}
	stringMap := make(map[string][]string)

	for _, str := range strs {
		mid := []int{}
		for _, r := range str {
			mid = append(mid, int(byte(r)))
		}
		sort.Ints(mid)
		newBytes := []byte{}
		for _, i := range mid {
			newBytes = append(newBytes, byte(i))
		}
		newStr := string(newBytes)
		if v, ok := stringMap[newStr]; !ok {
			stringMap[newStr] = []string{str}
		} else {
			stringMap[newStr] = append(v, str)
		}
	}
	res := [][]string{}
	for _, v := range stringMap {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "ate", "bat"}))
}
