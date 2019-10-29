package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	_, pos := BoyerMoore(haystack, needle)
	return pos
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func preBmBc(temp string) []int {
	var res []int
	tempLength := len(temp)
	for i := 0; i < 256; i++ {
		res = append(res, tempLength)
	}
	for i := 0; i < tempLength-1; i++ {
		res[temp[i]] = tempLength - 1 - i
	}
	return res
}

func getSuffix(temp string) [99999]int {
	var suff [99999]int
	var f int
	tempLength := len(temp)
	suff[tempLength-1] = tempLength
	g := tempLength - 1
	for i := tempLength - 2; i >= 0; i-- {
		if i > g && suff[i+tempLength-1-f] < i-g {
			suff[i] = suff[i+tempLength-1-f]
		} else {
			if i < g {
				g = i
			}
			f = i
			for g >= 0 && temp[g] == temp[g+tempLength-1-f] {
				g--
			}
			suff[i] = f - g
		}
	}
	return suff
}

func preBmGs(temp string) [99999]int {
	var res [99999]int
	tempLength := len(temp)
	suffix := getSuffix(temp)
	for i := 0; i < tempLength; i++ {
		res[i] = tempLength
	}
	j := 0
	for i := tempLength - 1; i >= 0; i-- {
		if suffix[i] == i+1 {
			for ; j < tempLength-1-i; j++ {
				if res[j] == tempLength {
					res[j] = tempLength - 1 - i
				}
			}
		}
	}
	for i := 0; i <= tempLength-2; i++ {
		res[tempLength-1-suffix[i]] = tempLength - 1 - i
	}
	return res
}

func BoyerMoore(master, temp string) (result bool, pos int) {
	masterLength := len(master)
	tempLength := len(temp)
	bmBc := preBmBc(temp)
	bmGs := preBmGs(temp)
	for j := 0; j <= masterLength-tempLength; {
		var i int
		for i = tempLength - 1; i >= 0 && temp[i] == master[i+j]; i-- {
		}
		if i < 0 {
			return true, j
		} else {
			j += max(bmBc[master[i+j]]-tempLength+1+i, bmGs[i])
		}
	}
	return false, -1
}

func main() {
	fmt.Println(strStr("aaaaaaaa", "aaaaaab"))
}
