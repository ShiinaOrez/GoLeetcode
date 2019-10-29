package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make(map[int][]byte)
	line := 1
	up := true
	for _, r := range s {
		rows[line] = append(rows[line], byte(r))
		if up {
			line += 1
			if line > numRows {
				line -= 2
				up = false
			}
		} else {
			line -= 1
			if line < 1 {
				line += 2
				up = true
			}
		}
	}
	var res []byte
	for i := 1; i <= numRows; i++ {
		for _, i := range rows[i] {
			res = append(res, i)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(convert("LEETCODEISHIRING", 4))
}
