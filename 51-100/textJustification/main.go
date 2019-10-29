package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	now := []string{}
	length := 0
	for _, word := range words {
		n := len(word)
		if length == 0 {
			if n <= maxWidth {
				now = append(now, word)
				length += n
			}
			continue
		}
		if length+1+n <= maxWidth {
			now = append(now, word)
			length += (1 + n)
		} else {
			spaces := maxWidth - length
			aver := spaces
			mod := 0
			if len(now) > 1 {
				aver = spaces / (len(now) - 1)
				mod = spaces % (len(now) - 1)
			}
			newLine := ""
			for i, word := range now {
				newLine = newLine + word
				if i < len(now)-1 || len(now) == 1 {
					for i := 1; i <= aver; i++ {
						newLine = newLine + " "
					}
					if len(now) > 1 {
						newLine = newLine + " "
					}
					if mod > 0 {
						newLine = newLine + " "
						mod -= 1
					}
				}
			}
			res = append(res, newLine)
			now = []string{word}
			length = len(word)
		}
	}
	lastLine := strings.Join(now, " ")
	for i := maxWidth - len(lastLine); i > 0; i-- {
		lastLine = lastLine + " "
	}
	res = append(res, lastLine)
	return res
}

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}
